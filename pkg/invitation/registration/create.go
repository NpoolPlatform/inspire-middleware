package registration

import (
	"context"
	"fmt"

	entpubsubmessage "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	"github.com/google/uuid"

	registrationcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
	pubsubmessagecrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/pubsubmessage"

	msgstate "github.com/NpoolPlatform/message/npool/pubsub/v1"
	msg "github.com/NpoolPlatform/message/npool/pubsub/v1/signupinvitation"
)

func CreateRegistration(ctx context.Context, in *mgrpb.RegistrationReq) (*mgrpb.Registration, error) {
	return mgrcli.CreateRegistration(ctx, in)
}

func TxCreateRegistration(
	ctx context.Context,
	uniqueID uuid.UUID,
	messageID, sender string,
	body []byte,
	req *msg.Request,
) error {
	info, err := pubsubmessagecrud.Row(ctx, uniqueID)
	if err != nil && !ent.IsValidationError(err) {
		return err
	}
	if info != nil {
		switch info.State {
		case msgstate.MessageState_Success.String():
			return nil
		case msgstate.MessageState_Field.String():
			return fmt.Errorf("%s", info.ErrorMessage)
		case msgstate.MessageState_InProcess.String():
			return fmt.Errorf("message state is exceptions")
		default:
			return fmt.Errorf("message state is invalid")
		}
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c1, err := pubsubmessagecrud.CreateSet(
			cli.PubsubMessage.Create(),
			uniqueID,
			messageID,
			sender,
			body,
			msgstate.MessageState_InProcess.String(),
			nil,
			nil,
		)
		if err != nil {
			return err
		}
		_, err = c1.Save(ctx)
		return err
	})
	if err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		c, err := registrationcrud.CreateSet(tx.Registration.Create(), &mgrpb.RegistrationReq{
			AppID:     &req.AppID,
			InviterID: &req.InviterID,
			InviteeID: &req.InviteeID,
		})
		if err != nil {
			return err
		}
		_, err = c.Save(ctx)
		if err != nil {
			return err
		}

		info, err = tx.PubsubMessage.Query().Where(entpubsubmessage.ID(uniqueID)).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := pubsubmessagecrud.UpdateSet(info, msgstate.MessageState_Success.String())
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
}
