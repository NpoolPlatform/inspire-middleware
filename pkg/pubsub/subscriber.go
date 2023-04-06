package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	entpubsubmsg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"

	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	msgpb "github.com/NpoolPlatform/message/npool/basetypes/v1"

	registrationmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

/// 1 Every received message should be acked and responded
/// 2 If 1 cannot fullfiled due to crashed, when it's received again, just responded and acked
/// 3 Never re-apply a message, if it resends due to 2, just responded and acked
/// 4 All message should be one-on-one message

var processingMsg sync.Map

func resp(mid string, rid uuid.UUID, err error) error {
	_resp := pubsub.Resp{}
	if err != nil {
		_resp.Code = -1
		_resp.Msg = err.Error()
	}
	return pubsub.Publish(mid, nil, &rid, nil, &_resp)
}

func msgCommiter(ctx context.Context, tx *ent.Tx, mid string, uid uuid.UUID, rid *uuid.UUID, err error) error {
	state := msgpb.MsgState_StateSuccess
	if err != nil {
		state = msgpb.MsgState_StateFail
	}

	c := tx.PubsubMessage.
		Create().
		SetID(uid).
		SetMessageID(mid).
		SetState(state.String())
	if rid != nil {
		c.SetRespToID(*rid)
	}
	_, err = c.Save(ctx)
	return err
}

func prepare(mid, body string) (interface{}, error) {
	var req interface{}

	switch mid {
	case msgpb.MsgID_CreateRegistrationInvitationTry.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationConfirm.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationCancel.String():
		_req := registrationmgrpb.RegistrationReq{}
		if err := json.Unmarshal([]byte(body), &req); err != nil {
			logger.Sugar().Errorw(
				"handler",
				"MID", mid,
				"Body", body,
			)
			return nil, err
		}
		req = &_req
	default:
		return nil, nil
	}

	return req, nil
}

/// Query a request message
///  Return
///   bool   appliable == true, caller should go ahead to apply this message
///   error  error message
func statReq(ctx context.Context, mid string, uid uuid.UUID) (bool, error) {
	var msg *ent.PubsubMessage
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		msg, err = cli.
			PubsubMessage.
			Query().
			Where(
				entpubsubmsg.ID(uid),
			).
			Only(_ctx)
		return err
	})

	switch err {
	case nil:
	default:
		if !ent.IsNotFound(err) {
			return true, nil
		}
		logger.Sugar().Warnw(
			"stat",
			"MID", mid,
			"UID", uid,
			"Error", err,
		)
		return false, resp(mid, uid, err)
	}

	switch msg.State {
	case msgpb.MsgState_StateSuccess.String():
		return false, resp(mid, uid, nil)
	case msgpb.MsgState_StateFail.String():
		return false, resp(mid, uid, fmt.Errorf("unknown error"))
	}

	return false, nil
}

func statResp(ctx context.Context, mid string, rid uuid.UUID) (bool, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.
			PubsubMessage.
			Query().
			Where(
				entpubsubmsg.RespToID(rid),
			).
			Only(_ctx)
		return err
	})

	switch err {
	case nil:
	default:
		if !ent.IsNotFound(err) {
			return true, nil
		}
		logger.Sugar().Warnw(
			"stat",
			"MID", mid,
			"RID", rid,
			"Error", err,
		)
		return false, err
	}

	return false, nil
}

/// Query a message state in database
///  Return
///   bool    appliable == true, caller should go ahead to apply this message
///   error   error message
func statMsg(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) {
	switch mid {
	case msgpb.MsgID_CreateRegistrationInvitationTry.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationConfirm.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationCancel.String():
		return statReq(ctx, mid, uid)
	/// We do not need to process resp here, but we keep it as a template
	case msgpb.MsgID_CreateRegistrationInvitationTryResp.String():
		if rid == nil {
			logger.Sugar().Warnw(
				"stat",
				"MID", mid,
				"UID", uid,
				"State", "Resp Without rid",
			)
			return false, fmt.Errorf("resp without id")
		}
		return statResp(ctx, mid, *rid)
	default:
		return false, fmt.Errorf("invalid message")
	}
}

/// Stat if message in right status, and is appliable
///  Return
///   bool    appliable == true, the message needs to be applied
///   error   error happens
func stat(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) {
	return statMsg(ctx, mid, uid, rid)
}

/// Process will consume the message and return consuming state
///  Return
///   error   reason of error, if nil, means the message should be acked
func process(ctx context.Context, mid string, uid uuid.UUID, req interface{}) (err error) {
	switch mid {
	case msgpb.MsgID_CreateRegistrationInvitationTry.String():
	case msgpb.MsgID_CreateRegistrationInvitationConfirm.String():
		err = registration.CreateRegistrationV2(
			ctx,
			req.(*registrationmgrpb.RegistrationReq),
			func(ctx context.Context, tx *ent.Tx, err error) error {
				return msgCommiter(ctx, tx, mid, uid, nil, err)
			})
	case msgpb.MsgID_CreateRegistrationInvitationCancel.String():
	default:
		return nil
	}

	if err != nil {
		logger.Sugar().Warnw(
			"process",
			"MID", mid,
			"UID", uid,
			"Req", req,
			"Error", err,
		)
	}

	switch mid {
	case msgpb.MsgID_CreateRegistrationInvitationTry.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationConfirm.String():
		fallthrough //nolint
	case msgpb.MsgID_CreateRegistrationInvitationCancel.String():
		return resp(mid, uid, err)
	default:
	}

	return err
}

/// No matter what handler return, the message will be acked, unless handler halt
/// If handler halt, the service will be restart, all message will be requeue
func handler(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID, body string) error {
	req, err := prepare(mid, body)
	if err != nil {
		return err
	}
	if req == nil {
		return nil
	}

	processingMsg.Store(uid, true)
	defer processingMsg.Delete(uid)

	appliable, err := stat(ctx, mid, uid, rid)
	if err != nil {
		return err
	}
	if !appliable {
		return nil
	}

	return process(ctx, mid, uid, req)
}

func Subscribe(ctx context.Context) error {
	return pubsub.Subscribe(ctx, handler)
}

func Shutdown(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		var err error
		processingMsg.Range(func(key, value interface{}) bool {
			_, err = tx.
				PubsubMessage.
				UpdateOneID(key.(uuid.UUID)).
				SetState(msgpb.MsgState_StateFail.String()).
				Save(_ctx)
			return err == nil
		})
		return err
	})
}
