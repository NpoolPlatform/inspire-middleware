package registration

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func UpdateRegistration(ctx context.Context, in *mgrpb.RegistrationReq) (*mgrpb.Registration, error) {
	return mgrcli.UpdateRegistration(ctx, in)
}
