package calculate

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	AppID                  uuid.UUID
	UserID                 uuid.UUID
	GoodID                 uuid.UUID
	OrderID                uuid.UUID
	PaymentID              uuid.UUID
	CoinTypeID             uuid.UUID
	PaymentCoinTypeID      uuid.UUID
	PaymentCoinUSDCurrency decimal.Decimal
	Units                  decimal.Decimal
	PaymentAmount          decimal.Decimal
	GoodValue              decimal.Decimal
	SettleType             types.SettleType
	HasCommission          bool
	OrderCreatedAt         uint32
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

func WithAppID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = _id
		return nil
	}
}

func WithUserID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.UserID = _id
		return nil
	}
}

func WithGoodID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.GoodID = _id
		return nil
	}
}

func WithOrderID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.OrderID = _id
		return nil
	}
}

func WithPaymentID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.PaymentID = _id
		return nil
	}
}

func WithCoinTypeID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.CoinTypeID = _id
		return nil
	}
}

func WithPaymentCoinTypeID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.PaymentCoinTypeID = _id
		return nil
	}
}

func WithPaymentCoinUSDCurrency(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentCoinUSDCurrency = _amount
		return nil
	}
}

func WithUnits(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.Units = _amount
		return nil
	}
}

func WithPaymentAmount(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentAmount = _amount
		return nil
	}
}

func WithGoodValue(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.GoodValue = _amount
		return nil
	}
}

func WithSettleType(settleType types.SettleType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch settleType {
		case types.SettleType_GoodOrderPercent:
		case types.SettleType_TechniqueFeePercent:
		default:
			return fmt.Errorf("invalid settletype")
		}
		h.SettleType = settleType
		return nil
	}
}

func WithHasCommission(value bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.HasCommission = value
		return nil
	}
}

func WithOrderCreatedAt(value uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.OrderCreatedAt = value
		return nil
	}
}
