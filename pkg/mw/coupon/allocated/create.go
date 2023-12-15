package allocated

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	coupon *ent.Coupon
}

func (h *createHandler) getCoupon(ctx context.Context) error {
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
		h.coupon = coupon
		return nil
	})
}

func (h *createHandler) cashable() bool {
	probability := h.coupon.CashableProbabilityPerMillion
	if probability.Cmp(decimal.NewFromInt(0)) <= 0 {
		return false
	}
	if probability.Cmp(decimal.NewFromInt(1)) >= 0 {
		return true
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Float64() //nolint
	return decimal.NewFromFloat(value).Cmp(h.coupon.CashableProbabilityPerMillion) <= 0
}

func (h *createHandler) createAllocatedCoupon(ctx context.Context, tx *ent.Tx) error {
	startAt := uint32(time.Now().Unix())
	couponScope := inspiretypes.CouponScope(inspiretypes.CouponScope_value[h.coupon.CouponScope])
	_cashable := h.cashable()
	if _, err := allocatedcrud.CreateSet(
		tx.CouponAllocated.Create(),
		&allocatedcrud.Req{
			EntID:        h.EntID,
			AppID:        h.AppID,
			CouponID:     h.CouponID,
			UserID:       h.UserID,
			StartAt:      &startAt,
			Denomination: &h.coupon.Denomination,
			CouponScope:  &couponScope,
			Cashable:     &_cashable,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateCoupon(ctx context.Context, tx *ent.Tx) error {
	allocated := h.coupon.Allocated
	switch h.coupon.CouponType {
	case inspiretypes.CouponType_FixAmount.String():
		allocated = allocated.Add(h.coupon.Denomination)
	case inspiretypes.CouponType_Discount.String():
		allocated = allocated.Add(decimal.NewFromInt(1))
	default:
		return fmt.Errorf("invalid coupontype")
	}
	if allocated.Cmp(h.coupon.Circulation) > 0 {
		return fmt.Errorf("insufficient circulation")
	}

	if _, err := tx.
		Coupon.
		UpdateOneID(h.coupon.ID).
		SetAllocated(allocated).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCoupon(ctx context.Context) (*npool.Coupon, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoupon(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAllocatedCoupon(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateCoupon(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoupon(ctx)
}
