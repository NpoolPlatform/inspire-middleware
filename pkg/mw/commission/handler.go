package commission

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID               *uuid.UUID
	AppID            *uuid.UUID
	UserID           *uuid.UUID
	GoodID           *uuid.UUID
	AppGoodID        *uuid.UUID
	FromGoodID       *uuid.UUID
	ToGoodID         *uuid.UUID
	FromAppGoodID    *uuid.UUID
	ToAppGoodID      *uuid.UUID
	SettleType       *types.SettleType
	SettleMode       *types.SettleMode
	SettleAmountType *types.SettleAmountType
	SettleInterval   *types.SettleInterval
	AmountOrPercent  *decimal.Decimal
	StartAt          *uint32
	Threshold        *decimal.Decimal
	ScalePercent     *decimal.Decimal
	Conds            *commissioncrud.Conds
	Offset           int32
	Limit            int32
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
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

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithFromGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid fromgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FromGoodID = &_id
		return nil
	}
}

func WithToGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid togoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ToGoodID = &_id
		return nil
	}
}

func WithFromAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid fromappgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FromAppGoodID = &_id
		return nil
	}
}

func WithToAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid toappgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ToAppGoodID = &_id
		return nil
	}
}

func WithSettleType(settleType *types.SettleType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleType == nil {
			if must {
				return fmt.Errorf("invalid settletype")
			}
			return nil
		}
		switch *settleType {
		case types.SettleType_GoodOrderPayment:
		case types.SettleType_TechniqueServiceFee:
		default:
			return fmt.Errorf("invalid settletype")
		}
		h.SettleType = settleType
		return nil
	}
}

func WithSettleMode(settleMode *types.SettleMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleMode == nil {
			if must {
				return fmt.Errorf("invalid settlemode")
			}
			return nil
		}
		switch *settleMode {
		case types.SettleMode_SettleWithGoodValue:
		case types.SettleMode_SettleWithPaymentAmount:
		default:
			return fmt.Errorf("invalid settlemode")
		}
		h.SettleMode = settleMode
		return nil
	}
}

func WithSettleAmountType(settleAmount *types.SettleAmountType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleAmount == nil {
			if must {
				return fmt.Errorf("invalid settleamounttype")
			}
			return nil
		}
		switch *settleAmount {
		case types.SettleAmountType_SettleByPercent:
		case types.SettleAmountType_SettleByAmount:
		default:
			return fmt.Errorf("invalid settleamount")
		}
		h.SettleAmountType = settleAmount
		return nil
	}
}

func WithSettleInterval(settleInterval *types.SettleInterval, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleInterval == nil {
			if must {
				return fmt.Errorf("invalid settleinterval")
			}
			return nil
		}
		switch *settleInterval {
		case types.SettleInterval_SettleAggregate:
		case types.SettleInterval_SettleMonthly:
		case types.SettleInterval_SettleYearly:
		case types.SettleInterval_SettleEveryOrder:
		default:
			return fmt.Errorf("invalid settlemode")
		}
		h.SettleInterval = settleInterval
		return nil
	}
}

func WithAmountOrPercent(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid amountorpercent")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.AmountOrPercent = &_amount
		return nil
	}
}

func WithStartAt(at *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StartAt = at
		return nil
	}
}

func WithThreshold(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid threshold")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.Threshold = &_amount
		return nil
	}
}

func WithScalePercent(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid scalepercent")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid scalepercent")
		}
		h.ScalePercent = &_amount
		return nil
	}
}

//nolint:funlen
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &commissioncrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{
				Op:  conds.GetAppGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.SettleType != nil {
			h.Conds.SettleType = &cruder.Cond{
				Op:  conds.GetSettleType().GetOp(),
				Val: types.SettleType(conds.GetSettleType().GetValue()),
			}
		}
		if conds.EndAt != nil {
			h.Conds.EndAt = &cruder.Cond{
				Op:  conds.GetEndAt().GetOp(),
				Val: conds.GetEndAt().GetValue(),
			}
		}
		if conds.UserIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetUserIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.UserIDs = &cruder.Cond{
				Op:  conds.GetUserIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.GoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.GoodIDs = &cruder.Cond{
				Op:  conds.GetGoodIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.StartAt != nil {
			h.Conds.StartAt = &cruder.Cond{
				Op:  conds.GetStartAt().GetOp(),
				Val: conds.GetStartAt().GetValue(),
			}
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
