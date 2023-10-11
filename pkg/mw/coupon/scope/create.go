package scope

import (
	"context"
	"fmt"

	scopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) verifyCoupon(ctx context.Context, tx *ent.Tx) error {
	coupon, err := tx.Coupon.Get(ctx, *h.CouponID)
	if err != nil {
		return err
	}
	if coupon == nil {
		return fmt.Errorf("coupon not found %v", *h.CouponID)
	}
	if h.CouponScope.String() != coupon.CouponScope {
		return fmt.Errorf("coupon scope not matched %v -> %v", h.CouponScope, coupon.CouponScope)
	}
	if *h.AppID != coupon.AppID {
		return fmt.Errorf("appid not matched %v -> %v", *h.AppID, coupon.AppID)
	}
	return nil
}

func (h *createHandler) createScope(ctx context.Context, tx *ent.Tx) error {
	if _, err := scopecrud.CreateSet(
		tx.CouponScope.Create(),
		&scopecrud.Req{
			ID:          h.ID,
			AppID:       h.AppID,
			AppGoodID:   h.AppGoodID,
			CouponID:    h.CouponID,
			CouponScope: h.CouponScope,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateScope(ctx context.Context) (*npool.Scope, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.verifyCoupon(ctx, tx); err != nil {
			return err
		}
		if err := handler.createScope(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetScope(ctx)
}
