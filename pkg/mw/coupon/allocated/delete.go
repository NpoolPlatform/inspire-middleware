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

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
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
			return wlog.WrapError(err)
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
		return wlog.Errorf("invalid coupontype")
	}
	if allocated.Cmp(h.coupon.Circulation) > 0 {
		return wlog.Errorf("insufficient circulation")
	}

	if _, err := tx.
		Coupon.
		UpdateOne(h.coupon).
		SetAllocated(allocated).
		Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteCoupon(ctx context.Context) error {
	info, err := h.GetCoupon(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID

	handler := &deleteHandler{
		Handler: h,
	}
	id := uuid.MustParse(info.CouponID)
	h.CouponID = &id
	if err := handler.getCoupon(ctx); err != nil {
		return wlog.WrapError(err)
	}

	now := uint32(time.Now().Unix())
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCoupon(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if _, err := allocatedcrud.UpdateSet(
			tx.CouponAllocated.UpdateOneID(*h.ID),
			&allocatedcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
