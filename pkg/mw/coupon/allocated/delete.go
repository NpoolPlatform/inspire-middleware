package allocated

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
)

func (h *Handler) DeleteCoupon(ctx context.Context) (*npool.Coupon, error) {
	info, err := h.GetCoupon(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := allocatedcrud.UpdateSet(
			cli.CouponAllocated.UpdateOneID(*h.ID),
			&allocatedcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
