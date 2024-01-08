package coin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	couponcoincrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"
)

func (h *Handler) DeleteCouponCoin(ctx context.Context) (*npool.CouponCoin, error) {
	info, err := h.GetCouponCoin(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := couponcoincrud.UpdateSet(
			cli.CouponCoin.UpdateOneID(*h.ID),
			&couponcoincrud.Req{
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
