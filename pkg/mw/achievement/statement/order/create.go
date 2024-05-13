package orderstatement

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderpaymentstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order/payment"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	goodachievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/good"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	sql                      string
	sqlCreateGoodAchievement string
}

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
	_sql += "where not exist ("
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
		goodachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodachievement1.WithGoodID(func() *string { s := h.GoodID.String(); return &s }(), true),
		goodachievement1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
		goodachievement1.WithTotalAmountUSD(func() *string { s := h.PaymentAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfAmountUSD(func() *string {
			s := decimal.NewFromInt(0).String()
			if *h.UserID == *h.OrderUserID {
				s = h.PaymentAmountUSD.String()
			}
			return &s
		}(), true),
		goodachievement1.WithTotalUnits(func() *string { s := h.Units.String(); return &s }(), true),
		goodachievement1.WithSelfUnits(func() *string {
			s := decimal.NewFromInt(0).String()
			if *h.UserID == *h.OrderUserID {
				s = h.Units.String()
			}
			return &s
		}(), true),
		goodachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfCommissionUSD(func() *string {
			s := decimal.NewFromInt(0).String()
			if *h.UserID == *h.OrderUserID {
				s = h.CommissionAmountUSD.String()
			}
			return &s
		}(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateGoodAchievement = handler.ConstructCreateSQL()
	return nil
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create orderstatement: %v", err)
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

func (h *createHandler) createOrUpdateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	// TODO: create achievement if coin type id is not exist
	// TODO: update achievement
	return nil
}

func (h *createHandler) createOrUpdateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	return nil
}

func (h *Handler) CreateStatement(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	for _, req := range h.PaymentStatementReqs {
		req.StatementID = h.EntID
	}
	handler := &createHandler{
		Handler: h,
	}

	handler.constructSQL()
	if err := handler.constructCreateGoodAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOrderStatement(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createPaymentStatements(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createOrUpdateGoodAchievement(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createOrUpdateGoodCoinAchievement(_ctx, tx)
	})
}
