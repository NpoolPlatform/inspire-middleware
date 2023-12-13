package scope

import (
	"context"
	"fmt"

	scopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createScope(ctx context.Context, cli *ent.Client) error {
	if _, err := scopecrud.CreateSet(
		cli.CouponScope.Create(),
		&scopecrud.Req{
			EntID:       h.EntID,
			GoodID:      h.GoodID,
			CouponID:    h.CouponID,
			CouponScope: h.CouponScope,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateScope(ctx context.Context) (*npool.Scope, error) {
	handler := &createHandler{
		Handler: h,
	}

	h.Conds = &scopecrud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		CouponScope: &cruder.Cond{Op: cruder.EQ, Val: *h.CouponScope},
	}
	exist, err := h.ExistScopeConds(ctx)
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
		if err := handler.createScope(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetScope(ctx)
}
