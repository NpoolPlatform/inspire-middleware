package registration

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func GetRegistrations(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Registration, uint32, error) {
	return mgrcli.GetRegistrations(ctx, conds, offset, limit)
}

func GetRegistration(ctx context.Context, id string) (*mgrpb.Registration, error) {
	return mgrcli.GetRegistration(ctx, id)
}

func GetRegistrationOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.Registration, error) {
	return mgrcli.GetRegistrationOnly(ctx, conds)
}
