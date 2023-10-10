package scope

import (
	"context"

	scopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	"github.com/google/uuid"
)

func (h *Handler) CreateScope(ctx context.Context) (*npool.Scope, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := scopecrud.CreateSet(
			cli.CouponScope.Create(),
			&scopecrud.Req{
				ID:          h.ID,
				AppID:       h.AppID,
				AppGoodID:   h.AppGoodID,
				CouponID:    h.CouponID,
				CouponScope: h.CouponScope,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetScope(ctx)
}
