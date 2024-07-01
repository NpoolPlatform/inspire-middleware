package orderstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievementuser "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievementuser"
	entgoodachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodachievement"
	entgoodcoinachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodcoinachievement"
	entorderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*achievementQueryHandler
	selfOrder               bool
	selfCommissionAmountUSD decimal.Decimal
	statement               *ent.OrderStatement
	payment                 *ent.OrderPaymentStatement
}

func (h *updateHandler) updateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	goodCoinAchievement, err := tx.
		GoodCoinAchievement.
		Query().
		Where(
			entgoodcoinachievement.AppID(*h.AppID),
			entgoodcoinachievement.UserID(*h.UserID),
			entgoodcoinachievement.GoodCoinTypeID(*h.GoodCoinTypeID),
			entgoodcoinachievement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	totalCommissionUsd := goodCoinAchievement.TotalCommissionUsd.Add(*h.CommissionAmountUSD)
	selfCommissionUsd := goodCoinAchievement.SelfCommissionUsd.Add(h.selfCommissionAmountUSD)
	if _, err := tx.
		GoodCoinAchievement.
		UpdateOne(goodCoinAchievement).
		SetTotalCommissionUsd(totalCommissionUsd).
		SetSelfCommissionUsd(selfCommissionUsd).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	goodAchievement, err := tx.
		GoodAchievement.
		Query().
		Where(
			entgoodachievement.AppID(*h.AppID),
			entgoodachievement.UserID(*h.UserID),
			entgoodachievement.GoodID(*h.GoodID),
			entgoodachievement.AppGoodID(*h.AppGoodID),
			entgoodachievement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	totalCommissionUsd := goodAchievement.TotalCommissionUsd.Add(*h.CommissionAmountUSD)
	selfCommissionUsd := goodAchievement.SelfCommissionUsd.Add(h.selfCommissionAmountUSD)
	if _, err := tx.
		GoodAchievement.
		UpdateOne(goodAchievement).
		SetTotalCommissionUsd(totalCommissionUsd).
		SetSelfCommissionUsd(selfCommissionUsd).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAchievementUser(ctx context.Context, tx *ent.Tx) error {
	achievementUser, err := tx.
		AchievementUser.
		Query().
		Where(
			entachievementuser.AppID(*h.AppID),
			entachievementuser.UserID(*h.UserID),
			entachievementuser.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	totalCommission := achievementUser.TotalCommission.Add(*h.CommissionAmountUSD)
	selfCommission := achievementUser.SelfCommission.Add(h.selfCommissionAmountUSD)
	if _, err := tx.
		AchievementUser.
		UpdateOne(achievementUser).
		SetTotalCommission(totalCommission).
		SetSelfCommission(selfCommission).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updatePaymentStatement(ctx context.Context, tx *ent.Tx) error {
	if h.statement.CommissionConfigType != types.CommissionConfigType_LegacyCommissionConfig.String() {
		return nil
	}

	for _, req := range h.PaymentStatementReqs {
		if req.Amount.Cmp(h.payment.Amount) != 0 {
			return wlog.Errorf("mismatch amount")
		}

		if req.CommissionAmount.Cmp(h.payment.CommissionAmount) == 0 {
			if req.CommissionAmount.Cmp(decimal.NewFromInt(0)) == 0 && h.statement.CommissionConfigID == uuid.Nil {
				if _, err := tx.
					OrderStatement.
					UpdateOneID(h.statement.ID).
					SetCommissionConfigID(*h.CommissionConfigID).
					Save(ctx); err != nil {
					return wlog.WrapError(err)
				}
			}
			return nil
		}

		if h.payment.CommissionAmount.Cmp(decimal.NewFromInt(0)) != 0 {
			return wlog.Errorf("permission denied")
		}

		if _, err := tx.
			OrderPaymentStatement.
			UpdateOneID(h.payment.ID).
			SetCommissionAmount(*req.CommissionAmount).
			Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updateGoodAchievement(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updateGoodCoinAchievement(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updateAchievementUser(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *updateHandler) requireOrderPaymentStatement(ctx context.Context, tx *ent.Tx) error {
	payment, err := tx.
		OrderPaymentStatement.
		Query().
		Where(
			entorderpaymentstatement.StatementID(h.statement.EntID),
			entorderpaymentstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	h.payment = payment
	return nil
}

func (h *updateHandler) requireOrderStatement(ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		OrderStatement.
		Query().
		Where(
			entorderstatement.AppID(*h.AppID),
			entorderstatement.UserID(*h.UserID),
			entorderstatement.OrderID(*h.OrderID),
			entorderstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.statement = statement
	return nil
}

func (h *Handler) UpdateStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	if err := h.validateCommissionAmount(); err != nil {
		return wlog.WrapError(err)
	}

	handler := &updateHandler{
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		selfOrder: *h.OrderUserID == *h.UserID,
		selfCommissionAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.CommissionAmountUSD
			}
			return decimal.NewFromInt(0)
		}(),
	}

	if err := handler.requireOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.requireOrderPaymentStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePaymentStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateStatementWithTx(_ctx, tx)
	})
}
