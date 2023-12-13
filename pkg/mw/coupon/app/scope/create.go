package scope

import (
	"context"
	"fmt"

	appgoodscopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppGoodScope(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodscopecrud.CreateSet(
		cli.AppGoodScope.Create(),
		&appgoodscopecrud.Req{
			EntID:       h.EntID,
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

func (h *Handler) CreateAppGoodScope(ctx context.Context) (*npool.Scope, error) {
	handler := &createHandler{
		Handler: h,
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
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createAppGoodScope(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetAppGoodScope(ctx)
}
