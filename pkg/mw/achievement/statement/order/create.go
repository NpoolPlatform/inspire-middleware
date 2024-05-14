//nolint:dupl
package orderstatement

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good"
	goodcoinachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good/coin"
	orderpaymentstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order/payment"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	goodachievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/good"
	goodcoinachievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/good/coin"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*achievementQueryHandler
	sql                          string
	sqlCreateGoodAchievement     string
	sqlCreateGoodCoinAchievement string
	selfOrder                    bool
	selfUnits                    decimal.Decimal
	selfAmountUSD                decimal.Decimal
	selfCommissionAmountUSD      decimal.Decimal
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into order_statements "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "good_id"
	_sql += comma + "app_good_id"
	_sql += comma + "order_id"
	_sql += comma + "order_user_id"
	_sql += comma + "good_coin_type_id"
	_sql += comma + "units"
	_sql += comma + "good_value_usd"
	_sql += comma + "payment_amount_usd"
	_sql += comma + "commission_amount_usd"
	_sql += comma + "app_config_id"
	_sql += comma + "commission_config_id"
	_sql += comma + "commission_config_type"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	_sql += fmt.Sprintf("%v'%v' as order_user_id", comma, *h.OrderUserID)
	_sql += fmt.Sprintf("%v'%v' as good_coin_type_id", comma, *h.GoodCoinTypeID)
	_sql += fmt.Sprintf("%v'%v' as units", comma, *h.Units)
	_sql += fmt.Sprintf("%v'%v' as good_value_usd", comma, *h.GoodValueUSD)
	_sql += fmt.Sprintf("%v'%v' as payment_amount_usd", comma, *h.PaymentAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as commission_amount_usd", comma, *h.CommissionAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as app_config_id", comma, *h.AppConfigID)
	_sql += fmt.Sprintf("%v'%v' as commission_config_id", comma, *h.CommissionConfigID)
	_sql += fmt.Sprintf("%v'%v' as commission_config_type", comma, h.CommissionConfigType.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from order_statements "
	_sql += fmt.Sprintf(
		"where user_id = '%v' and order_id = '%v' ",
		*h.UserID,
		*h.OrderID,
	)
	_sql += "limit 1)"

	h.sql = _sql
}

func (h *createHandler) constructCreateGoodAchievementSQL(ctx context.Context) error {
	handler, err := goodachievement1.NewHandler(
		ctx,
		goodachievement1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		goodachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodachievement1.WithGoodID(func() *string { s := h.GoodID.String(); return &s }(), true),
		goodachievement1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
		goodachievement1.WithTotalAmountUSD(func() *string { s := h.PaymentAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfAmountUSD(func() *string { s := h.selfAmountUSD.String(); return &s }(), true),
		goodachievement1.WithTotalUnits(func() *string { s := h.Units.String(); return &s }(), true),
		goodachievement1.WithSelfUnits(func() *string { s := h.selfUnits.String(); return &s }(), true),
		goodachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateGoodAchievement = handler.ConstructCreateSQL()
	return nil
}

func (h *createHandler) constructCreateGoodCoinAchievementSQL(ctx context.Context) error {
	handler, err := goodcoinachievement1.NewHandler(
		ctx,
		goodcoinachievement1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		goodcoinachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodcoinachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodcoinachievement1.WithGoodCoinTypeID(func() *string { s := h.GoodCoinTypeID.String(); return &s }(), true),
		goodcoinachievement1.WithTotalAmountUSD(func() *string { s := h.PaymentAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithSelfAmountUSD(func() *string { s := h.selfAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithTotalUnits(func() *string { s := h.Units.String(); return &s }(), true),
		goodcoinachievement1.WithSelfUnits(func() *string { s := h.selfUnits.String(); return &s }(), true),
		goodcoinachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateGoodCoinAchievement = handler.ConstructCreateSQL()
	return nil
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n > 1 {
		return wlog.Errorf("fail create orderstatement: %v", err)
	}
	if n == 0 {
		return wlog.WrapError(cruder.ErrCreateNothing)
	}
	return nil
}

func (h *createHandler) createOrderStatement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *createHandler) createPaymentStatements(ctx context.Context, tx *ent.Tx) error {
	for _, req := range h.PaymentStatementReqs {
		if _, err := orderpaymentstatementcrud.CreateSet(
			tx.OrderPaymentStatement.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) updateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	_, err := goodachievementcrud.UpdateSet(
		tx.GoodAchievement.UpdateOneID(h.entGoodAchievement.ID),
		&goodachievementcrud.Req{
			TotalAmountUSD: func() *decimal.Decimal { d := h.entGoodAchievement.TotalAmountUsd.Add(*h.PaymentAmountUSD); return &d }(),
			SelfAmountUSD:  func() *decimal.Decimal { d := h.entGoodAchievement.SelfAmountUsd.Add(h.selfAmountUSD); return &d }(),
			TotalUnits:     func() *decimal.Decimal { d := h.entGoodAchievement.TotalUnits.Add(*h.Units); return &d }(),
			SelfUnits:      func() *decimal.Decimal { d := h.entGoodAchievement.SelfUnits.Add(h.selfUnits); return &d }(),
			TotalCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodAchievement.TotalCommissionUsd.Add(*h.CommissionAmountUSD)
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

func (h *createHandler) createOrUpdateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	if h.entGoodAchievement != nil {
		return h.updateGoodAchievement(ctx, tx)
	}
	err := h.execSQL(ctx, tx, h.sqlCreateGoodAchievement)
	if err == nil {
		return nil
	}
	if !wlog.Equal(err, cruder.ErrCreateNothing) {
		return wlog.WrapError(err)
	}
	if err := h.requireAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return h.updateGoodAchievement(ctx, tx)
}

func (h *createHandler) updateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	_, err := goodcoinachievementcrud.UpdateSet(
		tx.GoodCoinAchievement.UpdateOneID(h.entGoodCoinAchievement.ID),
		&goodcoinachievementcrud.Req{
			TotalAmountUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.TotalAmountUsd.Add(*h.PaymentAmountUSD)
				return &d
			}(),
			SelfAmountUSD: func() *decimal.Decimal { d := h.entGoodCoinAchievement.SelfAmountUsd.Add(h.selfAmountUSD); return &d }(),
			TotalUnits:    func() *decimal.Decimal { d := h.entGoodCoinAchievement.TotalUnits.Add(*h.Units); return &d }(),
			SelfUnits:     func() *decimal.Decimal { d := h.entGoodCoinAchievement.SelfUnits.Add(h.selfUnits); return &d }(),
			TotalCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.TotalCommissionUsd.Add(*h.CommissionAmountUSD)
				return &d
			}(),
			SelfCommissionUSD: func() *decimal.Decimal {
				d := h.entGoodCoinAchievement.SelfCommissionUsd.Add(h.selfCommissionAmountUSD)
				return &d
			}(),
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *createHandler) createOrUpdateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	if h.entGoodCoinAchievement != nil {
		return h.updateGoodCoinAchievement(ctx, tx)
	}
	err := h.execSQL(ctx, tx, h.sqlCreateGoodCoinAchievement)
	if err == nil {
		return nil
	}
	if !wlog.Equal(err, cruder.ErrCreateNothing) {
		return wlog.WrapError(err)
	}
	if err := h.requireAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return h.updateGoodCoinAchievement(ctx, tx)
}

func (h *Handler) CreateStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	for _, req := range h.PaymentStatementReqs {
		req.StatementID = h.EntID
	}
	handler := &createHandler{
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		selfOrder: *h.OrderUserID == *h.UserID,
		selfUnits: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.Units
			}
			return decimal.NewFromInt(0)
		}(),
		selfAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.PaymentAmountUSD
			}
			return decimal.NewFromInt(0)
		}(),
		selfCommissionAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.CommissionAmountUSD
			}
			return decimal.NewFromInt(0)
		}(),
	}

	if err := handler.getAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}

	handler.constructSQL()
	if err := handler.constructCreateGoodAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCreateGoodCoinAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.createOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentStatements(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrUpdateGoodAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.createOrUpdateGoodCoinAchievement(ctx, tx)
}

func (h *Handler) CreateStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateStatementWithTx(_ctx, tx)
	})
}
