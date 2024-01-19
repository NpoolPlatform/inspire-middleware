package cashcontrol

import (
	"context"
	"fmt"

	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) getCoupon(ctx context.Context) error {
	couponID := h.CouponID.String()
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithEntID(&couponID, true),
	)
	if err != nil {
		return err
	}
	coupon, err := handler.GetCoupon(ctx)
	if err != nil {
		return err
	}
	if coupon == nil {
		return fmt.Errorf("invalid coupon")
	}
	if coupon.AppID != h.AppID.String() {
		return fmt.Errorf("invalid coupon")
	}
	if coupon.CouponType != inspiretypes.CouponType_FixAmount {
		return fmt.Errorf("invalid coupon type")
	}
	return nil
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
	if err := handler.getCoupon(ctx); err != nil {
		return nil, err
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
		return nil, fmt.Errorf("control type %v already exist", *h.ControlType)
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
