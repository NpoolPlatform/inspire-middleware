package allocated

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
)

func (h *Handler) UpdateCoupon(ctx context.Context) (*npool.Coupon, error) {
	if h.Used != nil && *h.Used && h.UsedByOrderID == nil {
		return nil, fmt.Errorf("invalid usedbyorderid")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := allocatedcrud.UpdateSet(
			cli.CouponAllocated.UpdateOneID(*h.ID),
			&allocatedcrud.Req{
				Used:          h.Used,
				UsedByOrderID: h.UsedByOrderID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoupon(ctx)
}
