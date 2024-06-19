//nolint:dupl
package orderstatement

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good"
	goodcoinachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good/coin"
	orderstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order"
	achievementusercrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/user"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievementuser "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievementuser"
	entorderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"

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
	inviteeConsumeAmount    decimal.Decimal
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
			entorderpaymentstatement.DeletedAt(0),
		).
		SetDeletedAt(h.now).
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
				d := h.entGoodAchievement.TotalCommissionUsd.Sub(h.commissionAmountUSD)
				return &d
			}(),
			SelfCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodAchievement.SelfCommissionUsd.Sub(h.selfCommissionAmountUSD)
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

func (h *deleteHandler) updateAchievementUser(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AchievementUser.
		Query().
		Where(
			entachievementuser.AppID(*h.AppID),
			entachievementuser.UserID(*h.UserID),
			entachievementuser.DeletedAt(0),
		).Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		&achievementusercrud.Req{
			TotalCommission: func() *decimal.Decimal {
				d := info.TotalCommission.Sub(h.commissionAmountUSD)
				return &d
			}(),
			SelfCommission: func() *decimal.Decimal {
				d := info.SelfCommission.Sub(h.selfCommissionAmountUSD)
				return &d
			}(),
			DirectConsumeAmount: func() *decimal.Decimal {
				d := info.DirectConsumeAmount.Sub(h.selfPaymentAmountUSD)
				return &d
			}(),
			InviteeConsumeAmount: func() *decimal.Decimal {
				d := info.InviteeConsumeAmount.Sub(h.inviteeConsumeAmount)
				return &d
			}(),
		}).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) getStatement(ctx context.Context, tx *ent.Tx) (*ent.OrderStatement, error) {
	stm := tx.OrderStatement.Query().Where(entorderstatement.DeletedAt(0))
	if h.ID == nil && h.EntID == nil {
		return nil, wlog.Errorf("id and ent id is empty")
	}
	if h.ID != nil {
		stm.Where(entorderstatement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderstatement.EntID(*h.EntID))
	}
	rows, err := stm.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	if len(rows) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return rows[0], nil
}

func (h *Handler) DeleteStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.getStatement(ctx, tx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	h.AppID = &info.AppID
	h.UserID = &info.UserID
	h.AppGoodID = &info.AppGoodID
	h.GoodCoinTypeID = &info.GoodCoinTypeID

	handler := &deleteHandler{
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.requireAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}

	handler.paymentAmountUSD = info.GoodValueUsd
	handler.inviteeConsumeAmount = info.GoodValueUsd
	handler.units = info.Units
	handler.commissionAmountUSD = info.CommissionAmountUsd
	if info.UserID == info.OrderUserID {
		handler.selfPaymentAmountUSD = handler.paymentAmountUSD
		handler.selfUnits = handler.units
		handler.selfCommissionAmountUSD = handler.commissionAmountUSD
		handler.inviteeConsumeAmount = decimal.NewFromInt(0)
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
	if err := handler.updateGoodCoinAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.updateAchievementUser(ctx, tx)
}

func (h *Handler) DeleteStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteStatementWithTx(_ctx, tx)
	})
}
