package cashcontrol

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Handler struct {
	cashcontrolcrud.Req
	Reqs   []*cashcontrolcrud.Req
	Conds  *cashcontrolcrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithCouponID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid couponid")
			}
			return nil
		}
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistCoupon(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid couponid")
		}
		_id := uuid.MustParse(*id)
		h.CouponID = &_id
		return nil
	}
}

func WithControlType(controlType *types.ControlType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if controlType == nil {
			if must {
				return fmt.Errorf("invalid control type")
			}
			return nil
		}
		switch *controlType {
		case types.ControlType_KycApproved:
		case types.ControlType_CreditsThreshold:
		case types.ControlType_OrderThreshold:
		case types.ControlType_PaymentAmountThreshold:
		default:
			return fmt.Errorf("invalid control type")
		}
		h.ControlType = controlType
		return nil
	}
}

func WithValue(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid value")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.Value = &_amount
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &cashcontrolcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.CouponID != nil {
			id, err := uuid.Parse(conds.GetCouponID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CouponID = &cruder.Cond{Op: conds.GetCouponID().GetOp(), Val: id}
		}
		if conds.ControlType != nil {
			h.Conds.ControlType = &cruder.Cond{Op: conds.GetControlType().GetOp(), Val: types.ControlType(conds.GetControlType().GetValue())}
		}
		if conds.CouponIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCouponIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CouponIDs = &cruder.Cond{Op: conds.GetCouponIDs().GetOp(), Val: ids}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
