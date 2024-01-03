package cashcontrol

import (
	"context"
	"fmt"

	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCashControl(ctx context.Context, cli *ent.Client) error {
	if _, err := cashcontrolcrud.CreateSet(
		cli.CashControl.Create(),
		&cashcontrolcrud.Req{
			EntID:       h.EntID,
			AppID:       h.AppID,
			CouponID:    h.CouponID,
			ControlType: h.ControlType,
			Value:       h.Value,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCashControl(ctx context.Context) (*npool.CashControl, error) {
	handler := &createHandler{
		Handler: h,
	}

	h.Conds = &cashcontrolcrud.Conds{
		AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		ControlType: &cruder.Cond{Op: cruder.EQ, Val: *h.ControlType},
	}
	exist, err := h.ExistCashControlConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coupon type %v already exist", *h.ControlType)
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createCashControl(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCashControl(ctx)
}
