package commission

import (
	"context"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	commissionmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/shopspring/decimal"
)

type Handler struct {
	SettleType    types.SettleType
	SettleAmount  types.SettleAmount
	Inviters      []*registrationmwpb.Registration
	Commissions   []*commissionmwpb.Commission
	PaymentAmount decimal.Decimal
	GoodValue     decimal.Decimal
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

func WithSettleType(settleType types.SettleType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SettleType = settleType
		return nil
	}
}

func WithSettleAmount(settleAmount types.SettleAmount) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SettleAmount = settleAmount
		return nil
	}
}

func WithInviters(inviters []*registrationmwpb.Registration) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Inviters = inviters
		return nil
	}
}

func WithCommissions(commissions []*commissionmwpb.Commission) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Commissions = commissions
		return nil
	}
}

func WithPaymentAmount(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentAmount = amount
		return nil
	}
}

func WithGoodValue(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.GoodValue = amount
		return nil
	}
}
