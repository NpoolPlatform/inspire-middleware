package pubsub

import (
	"context"

	msgid "github.com/NpoolPlatform/message/npool/pubsub/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

func Subscrib(ctx context.Context) {
	err := pubsub.Subscrib(
		ctx,
		func(
			ctx context.Context,
			messageID, sender string,
			uniqueID uuid.UUID,
			body []byte,
			respondToID *uuid.UUID,
		) error {
			switch messageID {
			case msgid.MessageID_SignupInvitationReq.String():
				return signupInvitation(ctx, messageID, sender, uniqueID, body)
			default:
				return nil
			}
		},
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Subscrib",
			"Error", err,
		)
	}
}
