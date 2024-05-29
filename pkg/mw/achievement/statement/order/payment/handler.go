package orderpaymentstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	orderstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order"
	orderpaymentstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order/payment"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"

	"github.com/google/uuid"
)

type Handler struct {
	PaymentStatementConds *orderpaymentstatementcrud.Conds
	OrderStatementConds   *orderstatementcrud.Conds
	Offset                int32
	Limit                 int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		PaymentStatementConds: &orderpaymentstatementcrud.Conds{},
		OrderStatementConds:   &orderstatementcrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func (h *Handler) withPaymentStatementConds(conds *npool.Conds) error {
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PaymentStatementConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.PaymentCoinTypeID != nil {
		id, err := uuid.Parse(conds.GetPaymentCoinTypeID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PaymentStatementConds.PaymentCoinTypeID = &cruder.Cond{Op: conds.GetPaymentCoinTypeID().GetOp(), Val: id}
	}
	return nil
}

func (h *Handler) withOrderStatementConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withPaymentStatementConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withOrderStatementConds(conds)
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
