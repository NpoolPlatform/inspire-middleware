package allocated

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"

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

func (h *Handler) UpdateCoupons(ctx context.Context) ([]*npool.Coupon, error) {
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if _, err := allocatedcrud.UpdateSet(
				tx.CouponAllocated.UpdateOneID(*req.ID),
				&allocatedcrud.Req{
					Used:          req.Used,
					UsedByOrderID: req.UsedByOrderID,
				},
			).Save(_ctx); err != nil {
				return err
			}
			ids = append(ids, *req.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &allocatedcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetCoupons(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
