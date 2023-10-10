package scope

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	scopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"
)

func (h *Handler) DeleteScope(ctx context.Context) (*npool.Scope, error) {
	info, err := h.GetScope(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := scopecrud.UpdateSet(
			cli.CouponScope.UpdateOneID(*h.ID),
			&scopecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
