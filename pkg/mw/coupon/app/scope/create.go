package scope

import (
	"context"
	"fmt"

	appgoodscopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

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
	return nil
}

func (h *createHandler) createAppGoodScope(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodscopecrud.CreateSet(
		tx.AppGoodScope.Create(),
		&appgoodscopecrud.Req{
			ID:          h.ID,
			AppID:       h.AppID,
			ScopeID:     h.ScopeID,
			AppGoodID:   h.AppGoodID,
			CouponID:    h.CouponID,
			CouponScope: h.CouponScope,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) getScope(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		scope, err := cli.CouponScope.Get(ctx, *h.ScopeID)
		if err != nil {
			return err
		}
		if scope == nil {
			return fmt.Errorf("scope not found %v", *h.ScopeID)
		}

		h.CouponID = &scope.CouponID
		if h.CouponScope == nil {
			couponScope := types.CouponScope(types.CouponScope_value[scope.CouponScope])
			h.CouponScope = &couponScope
		}
		return nil
	})
}

func (h *Handler) CreateAppGoodScope(ctx context.Context) (*npool.Scope, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.getScope(ctx); err != nil {
		return nil, err
	}

	h.Conds = &appgoodscopecrud.Conds{
		AppGoodID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		CouponScope: &cruder.Cond{Op: cruder.EQ, Val: *h.CouponScope},
	}
	exist, err := h.ExistAppGoodScopeConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coupon scope %v already exist", *h.CouponScope)
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.verifyCoupon(ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppGoodScope(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetAppGoodScope(ctx)
}
