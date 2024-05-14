package orderstatement

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good"
	goodcoinachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good/coin"
	orderstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entorderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*achievementQueryHandler
	now                     uint32
	paymentAmountUSD        decimal.Decimal
	selfPaymentAmountUSD    decimal.Decimal
	units                   decimal.Decimal
	selfUnits               decimal.Decimal
	commissionAmountUSD     decimal.Decimal
	selfCommissionAmountUSD decimal.Decimal
}

func (h *deleteHandler) deleteOrderStatement(ctx context.Context, tx *ent.Tx) error {
	_, err := orderstatementcrud.UpdateSet(
		tx.OrderStatement.UpdateOneID(*h.ID),
		&orderstatementcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) deletePaymentStatements(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.
		OrderPaymentStatement.
		Update().
		Where(
			entorderpaymentstatement.StatementID(*h.EntID),
		).
		Save(ctx)
	return err
}

func (h *deleteHandler) updateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	_, err := goodachievementcrud.UpdateSet(
		tx.GoodAchievement.UpdateOneID(h.entGoodAchievement.ID),
		&goodachievementcrud.Req{
			TotalAmountUSD: func() *decimal.Decimal { d := h.entGoodAchievement.TotalAmountUsd.Sub(h.paymentAmountUSD); return &d }(),
			SelfAmountUSD: func() *decimal.Decimal {
				d := h.entGoodAchievement.SelfAmountUsd.Sub(h.selfPaymentAmountUSD)
				return &d
			}(),
			TotalUnits: func() *decimal.Decimal { d := h.entGoodAchievement.TotalUnits.Sub(h.units); return &d }(),
			SelfUnits:  func() *decimal.Decimal { d := h.entGoodAchievement.SelfUnits.Sub(h.selfUnits); return &d }(),
			TotalCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodAchievement.TotalCommissionUsd.Add(h.commissionAmountUSD)
				return &d
			}(),
			SelfCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodAchievement.SelfCommissionUsd.Add(h.selfCommissionAmountUSD)
				return &d
			}(),
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) updateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	_, err := goodcoinachievementcrud.UpdateSet(
		tx.GoodCoinAchievement.UpdateOneID(h.entGoodCoinAchievement.ID),
		&goodcoinachievementcrud.Req{
			TotalAmountUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.TotalAmountUsd.Sub(h.paymentAmountUSD)
				return &d
			}(),
			SelfAmountUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.SelfAmountUsd.Sub(h.selfPaymentAmountUSD)
				return &d
			}(),
			TotalUnits: func() *decimal.Decimal { d := h.entGoodCoinAchievement.TotalUnits.Sub(h.units); return &d }(),
			SelfUnits:  func() *decimal.Decimal { d := h.entGoodCoinAchievement.SelfUnits.Sub(h.selfUnits); return &d }(),
			TotalCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.TotalCommissionUsd.Sub(h.commissionAmountUSD)
				return &d
			}(),
			SelfCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.SelfCommissionUsd.Sub(h.selfCommissionAmountUSD)
				return &d
			}(),
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.GetStatement(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	h.EntID = func() *uuid.UUID { uid, _ := uuid.Parse(info.EntID); return &uid }()
	h.UserID = func() *uuid.UUID { uid, _ := uuid.Parse(info.UserID); return &uid }()
	h.AppGoodID = func() *uuid.UUID { uid, _ := uuid.Parse(info.AppGoodID); return &uid }()
	h.GoodCoinTypeID = func() *uuid.UUID { uid, _ := uuid.Parse(info.GoodCoinTypeID); return &uid }()

	handler := &deleteHandler{
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.requireAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}

	handler.paymentAmountUSD = func() decimal.Decimal { amount, _ := decimal.NewFromString(info.PaymentAmountUSD); return amount }()
	handler.units = func() decimal.Decimal { amount, _ := decimal.NewFromString(info.Units); return amount }()
	handler.commissionAmountUSD = func() decimal.Decimal { amount, _ := decimal.NewFromString(info.CommissionAmountUSD); return amount }()
	if info.UserID == info.OrderUserID {
		handler.selfPaymentAmountUSD = handler.paymentAmountUSD
		handler.selfUnits = handler.units
		handler.selfCommissionAmountUSD = handler.commissionAmountUSD
	}

	if err := handler.deleteOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deletePaymentStatements(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateGoodAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.updateGoodCoinAchievement(ctx, tx)
}

func (h *Handler) DeleteStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteStatementWithTx(_ctx, tx)
	})
}
