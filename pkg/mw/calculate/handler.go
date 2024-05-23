package calculate

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	calculatemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	AppID            uuid.UUID
	UserID           uuid.UUID
	GoodID           uuid.UUID
	AppGoodID        uuid.UUID
	OrderID          uuid.UUID
	GoodCoinTypeID   uuid.UUID
	Units            decimal.Decimal
	PaymentAmountUSD decimal.Decimal
	GoodValueUSD     decimal.Decimal
	SettleType       types.SettleType
	SettleAmountType types.SettleAmountType
	HasCommission    bool
	OrderCreatedAt   uint32
	Payments         []calculatemwpb.Payment
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

func WithAppGoodID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppGoodID = _id
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

func WithGoodCoinTypeID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.GoodCoinTypeID = _id
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

func WithPaymentAmountUSD(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentAmountUSD = _amount
		return nil
	}
}

func WithGoodValueUSD(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.GoodValueUSD = _amount
		return nil
	}
}

func WithSettleType(settleType types.SettleType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch settleType {
		case types.SettleType_GoodOrderPayment:
		case types.SettleType_TechniqueServiceFee:
		default:
			return wlog.Errorf("invalid settletype")
		}
		h.SettleType = settleType
		return nil
	}
}

func WithSettleAmountType(settleAmount types.SettleAmountType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch settleAmount {
		case types.SettleAmountType_SettleByPercent:
		case types.SettleAmountType_SettleByAmount:
		default:
			return wlog.Errorf("invalid settleamount")
		}
		h.SettleAmountType = settleAmount
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

func WithPayments(payments []*calculatemwpb.Payment) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(payments) == 0 {
			return wlog.Errorf("invalid payments")
		}
		for _, payment := range payments {
			if _, err := uuid.Parse(payment.CoinTypeID); err != nil {
				return err
			}
			amount, err := decimal.NewFromString(payment.Amount)
			if err != nil {
				return err
			}
			if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
				return wlog.Errorf("invalid amount")
			}
			h.Payments = append(h.Payments, calculatemwpb.Payment{
				CoinTypeID: payment.CoinTypeID,
				Amount:     payment.Amount,
			})
		}
		return nil
	}
}
