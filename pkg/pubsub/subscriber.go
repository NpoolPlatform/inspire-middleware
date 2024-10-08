package pubsub

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entpubsubmsg "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/pubsubmessage"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/pubsub/coupon/allocated"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/google/uuid"
)

var subscriber *pubsub.Subscriber
var publisher *pubsub.Publisher

// TODO: here we should call from DB transaction context
func finish(ctx context.Context, msg *pubsub.Msg, err error) error {
	state := basetypes.MsgState_StateSuccess
	if err != nil {
		state = basetypes.MsgState_StateFail
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.
			PubsubMessage.
			Create().
			SetEntID(msg.UID).
			SetMessageID(msg.MID).
			SetArguments(msg.Body).
			SetState(state.String())
		if msg.RID != nil {
			c.SetRespToID(*msg.RID)
		}
		if msg.UnID != nil {
			c.SetUndoID(*msg.UnID)
		}
		_, err = c.Save(ctx)
		return wlog.WrapError(err)
	})
}

func prepare(mid, body string) (req interface{}, err error) {
	switch mid {
	case basetypes.MsgID_UpdateCouponsUsedReq.String():
		req, err = allocated1.Prepare(body)
	default:
		return nil, nil
	}

	if err != nil {
		logger.Sugar().Errorw(
			"handler",
			"MID", mid,
			"Body", body,
		)
		return nil, err
	}

	return req, nil
}

// Query a request message
//  Return
//   bool   appliable == true, caller should go ahead to apply this message
//   error  error message
func statReq(ctx context.Context, mid string, uid uuid.UUID) (bool, error) {
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err = cli.
			PubsubMessage.
			Query().
			Where(
				entpubsubmsg.EntID(uid),
			).
			Only(_ctx)
		return wlog.WrapError(err)
	})

	switch err {
	case nil:
	default:
		if ent.IsNotFound(err) {
			return true, nil
		}
		logger.Sugar().Warnw(
			"stat",
			"MID", mid,
			"UID", uid,
			"Error", err,
		)
		return false, wlog.WrapError(err)
	}

	return false, nil
}

// Query a message state in database
//  Return
//   bool    appliable == true, caller should go ahead to apply this message
//   error   error message
func statMsg(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) { //nolint
	switch mid {
	case basetypes.MsgID_RewardEventReq.String():
		return statReq(ctx, mid, uid)
	case basetypes.MsgID_UpdateCouponsUsedReq.String():
		return statReq(ctx, mid, uid)
	default:
		return false, wlog.Errorf("invalid message")
	}
}

// Stat if message in right status, and is appliable
//  Return
//   bool    appliable == true, the message needs to be applied
//   error   error happens
func stat(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) {
	return statMsg(ctx, mid, uid, rid)
}

// Process will consume the message and return consuming state
//  Return
//   error   reason of error, if nil, means the message should be acked
func process(ctx context.Context, mid string, uid uuid.UUID, req interface{}) (err error) { //nolint
	switch mid {
	case basetypes.MsgID_UpdateCouponsUsedReq.String():
		err = allocated1.Apply(ctx, req, publisher)
	default:
		return nil
	}
	return wlog.WrapError(err)
}

// No matter what handler return, the message will be acked, unless handler halt
// If handler halt, the service will be restart, all message will be requeue
func handler(ctx context.Context, msg *pubsub.Msg) (err error) {
	var req interface{}
	var appliable bool

	defer func() {
		msg.Ack()
		if req != nil && appliable {
			_ = finish(ctx, msg, err)
		}
	}()

	req, err = prepare(msg.MID, msg.Body)
	if err != nil {
		return wlog.WrapError(err)
	}
	if req == nil {
		return nil
	}

	appliable, err = stat(ctx, msg.MID, msg.UID, msg.RID)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !appliable {
		return nil
	}

	err = process(ctx, msg.MID, msg.UID, req)
	return wlog.WrapError(err)
}

func Subscribe(ctx context.Context) (err error) {
	subscriber, err = pubsub.NewSubscriber()
	if err != nil {
		return wlog.WrapError(err)
	}

	publisher, err = pubsub.NewPublisher()
	if err != nil {
		return wlog.WrapError(err)
	}

	return subscriber.Subscribe(ctx, handler)
}

func Shutdown(ctx context.Context) error {
	if subscriber != nil {
		subscriber.Close()
	}
	if publisher != nil {
		publisher.Close()
	}

	return nil
}
