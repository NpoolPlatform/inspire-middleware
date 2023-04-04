package pubsub

import (
	"context"
	"fmt"

	pubsubmessagecrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/pubsubmessage"
	entpubsubmessage "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	msg "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

func responseSuccessMsg(respondToID uuid.UUID) error {
	return pubsub.Publish(msg.MessageID_SignupInvitationResponse.String(), &respondToID, pubsub.MessageResp{
		Code: 0,
	})
}

func responseFieldMsg(respondToID uuid.UUID, message string, body []byte) error {
	return pubsub.Publish(msg.MessageID_SignupInvitationResponse.String(), &respondToID, pubsub.MessageResp{
		Code:    1,
		Message: message,
		Body:    body,
	})
}

func createMessage(
	ctx context.Context,
	messageID, sender string,
	uniqueID uuid.UUID,
	body []byte,
	respondToID *uuid.UUID,
) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c1, err := pubsubmessagecrud.CreateSet(
			cli.PubsubMessage.Create(),
			uniqueID,
			messageID,
			sender,
			body,
			msg.MessageState_InProcess.String(),
			respondToID,
			nil,
		)
		if err != nil {
			return err
		}
		_, err = c1.Save(ctx)
		return err
	})
}

func processMessage(
	ctx context.Context,
	messageID string,
	tx *ent.Tx,
	body []byte,
	isFailed bool,
	isFailedMessage string,
) (tx1 *ent.Tx, err error) {
	switch messageID {
	case msg.MessageID_SignupInvitationRequest.String():
		if isFailed {
			return nil, fmt.Errorf("%v", isFailedMessage)
		}
		return registration.TxCreateRegistration(ctx, tx, body)
	default:
		return tx, nil
	}
}

func Subscribe(ctx context.Context) {
	err := pubsub.Subscribe(
		ctx,
		func(
			ctx context.Context,
			messageID, sender string,
			uniqueID uuid.UUID,
			body []byte,
			respondToID *uuid.UUID,
		) error {
			var isFailed = false
			var failedMessage string
			var err error

			switch messageID {
			case msg.MessageID_SignupInvitationRequest.String():
				info, err := pubsubmessagecrud.Row(ctx, uniqueID)
				if err != nil && !ent.IsValidationError(err) {
					return responseFieldMsg(uniqueID, err.Error(), body)
				}
				if info != nil {
					switch info.State {
					case msg.MessageState_Success.String():
						return nil
					case msg.MessageState_Fail.String():
						isFailed = true
						failedMessage = info.ErrorMessage
					case msg.MessageState_InProcess.String():
						return fmt.Errorf("message state is exceptions")
					default:
						return fmt.Errorf("message state is invalid")
					}
				}
				err = createMessage(ctx, messageID, sender, uniqueID, body, respondToID)
				if err != nil {
					return responseFieldMsg(uniqueID, err.Error(), body)
				}
			default:
				return nil
			}

			err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
				tx, err = processMessage(ctx, messageID, tx, body, isFailed, failedMessage)
				if err != nil {
					return err
				}
				_, err = tx.
					PubsubMessage.
					Update().
					Where(
						entpubsubmessage.ID(uniqueID),
					).
					SetState(msg.MessageState_Success.String()).
					Save(ctx)
				return err
			})
			if err != nil {
				return responseFieldMsg(uniqueID, err.Error(), body)
			}
			return responseSuccessMsg(uniqueID)
		},
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Subscrib",
			"Error", err,
		)
	}
}

func SetMessageStatusOnExit() error {
	return db.WithClient(context.Background(), func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.
			PubsubMessage.
			Update().
			Where(
				entpubsubmessage.State(msg.MessageState_InProcess.String()),
			).
			SetState(msg.MessageState_Fail.String()).
			Save(_ctx)
		return err
	})
}
