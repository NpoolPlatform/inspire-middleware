package registration

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	registrationcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"

	pubsubhandler "github.com/NpoolPlatform/inspire-middleware/pkg/pubsub/handler"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func CreateRegistration(ctx context.Context, in *mgrpb.RegistrationReq) (*mgrpb.Registration, error) {
	return mgrcli.CreateRegistration(ctx, in)
}

func CreateRegistrationV2(
	ctx context.Context,
	in *mgrpb.RegistrationReq,
	postHandler pubsubhandler.PostHandler,
) error {
	var err error

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		defer func() {
			if postHandler != nil {
				_ = postHandler(ctx, tx, err)
			}
		}()

		var c *ent.RegistrationCreate
		c, err = registrationcrud.CreateSet(tx.Registration.Create(), in)
		if err != nil {
			return err
		}
		_, err = c.Save(ctx)
		return err
	})
}
