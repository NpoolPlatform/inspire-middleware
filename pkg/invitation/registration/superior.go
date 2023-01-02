//nolint:dupl
package registration

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entreg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/registration"

	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
)

func GetSuperiores(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Registration, uint32, error) {
	var infos []*mgrpb.Registration
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		sel := stm.
			Select(
				entreg.FieldID,
				entreg.FieldAppID,
				entreg.FieldInviterID,
				entreg.FieldInviteeID,
				entreg.FieldCreatedAt,
				entreg.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				// TODO: get parent recursively
			})

		_total, err := sel.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return sel.
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
