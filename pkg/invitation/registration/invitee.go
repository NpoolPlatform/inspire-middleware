package registration

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/invitation/registration"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func GetInvitees(
	ctx context.Context,
	conds *mgrpb.Conds,
	offset, limit int32,
) (
	infos []*mgrpb.Registration,
	total uint32,
	err error,
) {
	var entities []*ent.Registration

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		entities, err = stm.
			Offset(int(offset)).
			Limit(int(limit)).
			All(_ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}

	return converter.Ent2GrpcMany(entities), total, nil
}
