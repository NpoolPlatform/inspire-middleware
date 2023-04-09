package pubsub

import (
	"context"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	entpubsubmsg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

var processingMsg sync.Map
var subscriber *pubsub.Subscriber
var publisher *pubsub.Publisher

func msgCommiter(ctx context.Context, tx *ent.Tx, mid string, uid uuid.UUID, rid *uuid.UUID, err error) error { //nolint
	state := basetypes.MsgState_StateSuccess
	if err != nil {
		state = basetypes.MsgState_StateFail
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

func prepare(mid, body string) (req interface{}, err error) {
	switch mid {
	case basetypes.MsgID_RewardEventReq.String():
		req, err = prepareRewardEvent(body)
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

/// Query a request message
///  Return
///   bool   appliable == true, caller should go ahead to apply this message
///   error  error message
func statReq(ctx context.Context, mid string, uid uuid.UUID) (bool, error) {
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err = cli.
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
		return false, err
	}

	return false, nil
}

/// Query a message state in database
///  Return
///   bool    appliable == true, caller should go ahead to apply this message
///   error   error message
func statMsg(ctx context.Context, mid string, uid uuid.UUID, rid *uuid.UUID) (bool, error) { //nolint
	switch mid {
	case basetypes.MsgID_RewardEventReq.String():
		return statReq(ctx, mid, uid)
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
	defer func() {
		if err != nil {
			logger.Sugar().Warnw(
				"process",
				"MID", mid,
				"UID", uid,
				"Req", req,
				"Error", err,
			)
		}
	}()

	switch mid {
	case basetypes.MsgID_RewardEventReq.String():
		err = handleRewardEvent(ctx, req)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

/// No matter what handler return, the message will be acked, unless handler halt
/// If handler halt, the service will be restart, all message will be requeue
func handler(ctx context.Context, msg *pubsub.Msg) error {
	req, err := prepare(msg.MID, msg.Body)
	if err != nil {
		return err
	}
	if req == nil {
		return nil
	}

	processingMsg.Store(msg.UID, msg)
	defer func() {
		msg.Ack()
		processingMsg.Delete(msg.UID)
	}()

	appliable, err := stat(ctx, msg.MID, msg.UID, msg.RID)
	if err != nil {
		return err
	}
	if !appliable {
		return nil
	}

	return process(ctx, msg.MID, msg.UID, req)
}

func Subscribe(ctx context.Context) (err error) {
	subscriber, err = pubsub.NewSubscriber()
	if err != nil {
		return err
	}

	publisher, err = pubsub.NewPublisher()
	if err != nil {
		return err
	}

	return subscriber.Subscribe(ctx, handler)
}

/// TODO: if this will be run after signal catched ?
func Shutdown(ctx context.Context) error {
	if err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		var err error
		var info *ent.PubsubMessage
		processingMsg.Range(func(key, msg interface{}) bool {
			info, err = tx.
				PubsubMessage.
				Query().
				Where(
					entpubsubmsg.ID(key.(uuid.UUID)),
				).
				Only(_ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return false
				}
			}
			if info == nil {
				msg.(*pubsub.Msg).Nack()
				return true
			}

			_, err = tx.
				PubsubMessage.
				UpdateOneID(key.(uuid.UUID)).
				SetState(basetypes.MsgState_StateFail.String()).
				Save(_ctx)
			msg.(*pubsub.Msg).Ack()
			return err == nil
		})
		return err
	}); err != nil {
		return err
	}

	if subscriber != nil {
		subscriber.Close()
	}
	if publisher != nil {
		publisher.Close()
	}

	return nil
}
