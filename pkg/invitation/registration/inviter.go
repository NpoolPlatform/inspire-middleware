package registration

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func GetInviters(
	ctx context.Context,
	conds *mgrpb.Conds,
	offset, limit int32,
) (
	[]*mgrpb.Registration,
	uint32,
	error,
) {
	return nil, 0, nil
}
