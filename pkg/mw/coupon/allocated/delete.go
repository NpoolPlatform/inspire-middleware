package allocated

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	coupon *ent.Coupon
}

//nolint:dupl
func (h *deleteHandler) getCoupon(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		coupon, err := cli.
			Coupon.
			Query().
			Where(
				entcoupon.EntID(*h.CouponID),
				entcoupon.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		now := time.Now().Unix()
		if now < int64(coupon.StartAt) || now > int64(coupon.EndAt) {
			return fmt.Errorf("coupon can not be issued in current time")
		}
		h.coupon = coupon
		return nil
	})
}

//nolint:dupl
func (h *deleteHandler) updateCoupon(ctx context.Context, tx *ent.Tx) error {
	allocated := h.coupon.Allocated
	switch h.coupon.CouponType {
	case inspiretypes.CouponType_FixAmount.String():
		allocated = allocated.Sub(h.coupon.Denomination)
	case inspiretypes.CouponType_Discount.String():
		allocated = allocated.Sub(decimal.NewFromInt(1))
	default:
		return fmt.Errorf("invalid coupontype")
	}
	if allocated.Cmp(h.coupon.Circulation) > 0 {
		return fmt.Errorf("insufficient circulation")
	}

	if _, err := tx.
		Coupon.
		UpdateOne(h.coupon).
		SetAllocated(allocated).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCoupon(ctx context.Context) (*npool.Coupon, error) {
	info, err := h.GetCoupon(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler: h,
	}
	id := uuid.MustParse(info.CouponID)
	h.CouponID = &id
	if err := handler.getCoupon(ctx); err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCoupon(ctx, tx); err != nil {
			return err
		}
		if _, err := allocatedcrud.UpdateSet(
			tx.CouponAllocated.UpdateOneID(*h.ID),
			&allocatedcrud.Req{
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
