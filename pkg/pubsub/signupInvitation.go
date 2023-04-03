package pubsub

import (
	"context"
	"encoding/json"

	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	entpubsubmessage "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"
	msg "github.com/NpoolPlatform/message/npool/pubsub/v1"
	msgsignupinvitation "github.com/NpoolPlatform/message/npool/pubsub/v1/signupinvitation"
	"github.com/google/uuid"
)

func responseSuccessMsg(respondToID uuid.UUID) error {
	return pubsub.Publish(msg.MessageID_SignupInvitationRes.String(), &respondToID, nil)
}

func responseFieldMsg(respondToID uuid.UUID, userID string) error {
	return pubsub.Publish(msg.MessageID_SignupInvitationRes.String(), &respondToID, msgsignupinvitation.Response{
		State:  msg.MessageState_Field,
		UserID: userID,
	})
}

func signupInvitation(ctx context.Context,
	messageID, sender string,
	uniqueID uuid.UUID,
	body []byte) error {
	var req msgsignupinvitation.Request

	err := json.Unmarshal(body, &req)
	if err != nil {
		return responseFieldMsg(uniqueID, req.InviteeID)
	}
	err = registration.TxCreateRegistration(ctx, uniqueID, sender, messageID, body, &req)
	if err != nil {
		return responseFieldMsg(uniqueID, req.InviteeID)
	}

	return responseSuccessMsg(uniqueID)
}

func AdjustmentMessageState() error {
	return db.WithClient(context.Background(), func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.
			PubsubMessage.
			Update().
			Where(
				entpubsubmessage.State(msg.MessageState_InProcess.String()),
			).
			SetState(msg.MessageState_Field.String()).
			Save(_ctx)
		return err
	})
}
