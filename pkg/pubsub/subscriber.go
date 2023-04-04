package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	pubsubmsgcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/pubsubmessage"

	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	msgpb "github.com/NpoolPlatform/message/npool/basetypes/v1"

	registrationmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

var processingMsg sync.Map

func response(msgID string, respID *uuid.UUID, body []byte, msg string, code int) error {
	return pubsub.Publish(msgID, respID, pubsub.MessageResp{
		Code:    code,
		Message: msg,
		Body:    body,
	})
}

func postHandler(ctx context.Context, tx *ent.Tx, msgID string, uid uuid.UUID, err error) error {
	state := msgpb.MessageState_Success
	if err != nil {
		state = msgpb.MessageState_Fail
	}

	_, err = tx.
		PubsubMessage.
		UpdateOneID(uid).
		SetState(state.String()).
		Save(ctx)
	if err != nil {
		return err
	}

	// Dispatch resp according to err and message type
	switch msgID { //nolint
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		return nil
	}

	return nil
}

func prepare(msgID string, body []byte) (interface{}, error) {
	var req interface{}

	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		_req := registrationmgrpb.RegistrationReq{}
		if err := json.Unmarshal(body, &req); err != nil {
			logger.Sugar().Errorw(
				"handler",
				"MsgID", msgID,
				"Body", string(body),
			)
			// For message with invalid body, nobody can process it, so just act directly
			return nil, err
		}
		req = &_req
	default:
		return nil, nil
	}

	return req, nil
}

var applyWhenFail = map[string]bool{}

func init() {
	applyWhenFail[msgpb.MessageID_CreateRegistrationInvitationReq.String()] = false
}

/// Stat if message in right status, and is appliable
///  Return
///   bool    appliable = true, the message needs to be applied
///   error   error happens
func stat(ctx context.Context, msgID string, uid uuid.UUID, respToID *uuid.UUID) (bool, error) {
	var msg *ent.PubsubMessage
	var err error

	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		msg, err = pubsubmsgcrud.Row(ctx, uid)
		if err != nil {
			if !ent.IsNotFound(err) {
				return false, err
			}
		}
	default:
		return false, nil
	}

	if msg != nil {
		switch msg.State {
		case msgpb.MessageState_Processing.String():
			return false, fmt.Errorf("processing")
		case msgpb.MessageState_Success.String():
			return false, nil
		case msgpb.MessageState_Fail.String():
			return applyWhenFail[msg.MessageID], nil
		}
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.
			PubsubMessage.
			Create().
			SetID(uid).
			SetMessageID(msgID).
			SetState(msgpb.MessageState_Processing.String())
		if respToID != nil {
			c.SetRespToID(*respToID)
		}
		_, err := c.Save(_ctx)
		return err
	})

	return err == nil, err
}

func process(ctx context.Context, msgID string, uid uuid.UUID, req interface{}) error {
	var err error

	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		err = registration.CreateRegistrationV2(
			ctx,
			req.(*registrationmgrpb.RegistrationReq),
			func(ctx context.Context, tx *ent.Tx, err error) error {
				return postHandler(ctx, tx, msgID, uid, err)
			})
	default:
		return nil
	}

	return err
}

func handler(ctx context.Context, msgID, sender string, uid uuid.UUID, body []byte, respToID *uuid.UUID) error {
	req, err := prepare(msgID, body)
	if err != nil {
		return response(msgID, &uid, body, err.Error(), -1)
	}
	if req == nil {
		return response(msgID, &uid, body, "", 0)
	}

	processingMsg.Store(uid, true)
	defer processingMsg.Delete(uid)

	appliable, err := stat(ctx, msgID, uid, respToID)
	if err != nil {
		return response(msgID, &uid, body, err.Error(), -1)
	}
	if !appliable {
		return response(msgID, &uid, body, "", -1)
	}

	err = process(ctx, msgID, uid, req)
	if err != nil {
		return response(msgID, &uid, body, "", -1)
	}
	return response(msgID, &uid, body, "", 0)
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
				SetState(msgpb.MessageState_Fail.String()).
				Save(_ctx)
			return err == nil
		})
		return err
	})
}
