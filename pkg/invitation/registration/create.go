package registration

import (
	"context"
	"encoding/json"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	registrationcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func CreateRegistration(ctx context.Context, in *mgrpb.RegistrationReq) (*mgrpb.Registration, error) {
	return mgrcli.CreateRegistration(ctx, in)
}

func TxCreateRegistration(
	ctx context.Context,
	tx *ent.Tx,
	body []byte,
) (*ent.Tx, error) {
	var req = mgrpb.RegistrationReq{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	c, err := registrationcrud.CreateSet(tx.Registration.Create(), &req)
	if err != nil {
		return nil, err
	}
	_, err = c.Save(ctx)

	return nil, err
}
