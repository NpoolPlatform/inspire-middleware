package reward

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into user_rewards "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "action_credits"
	_sql += comma + "coin_pre_usd"
	_sql += comma + "coupon_amount"
	_sql += comma + "coupon_cashable_amount"
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
	_sql += fmt.Sprintf("%v'%v' as action_credits", comma, *h.ActionCredits)
	_sql += fmt.Sprintf("%v'%v' as coin_pre_usd", comma, *h.CoinPreUSD)
	_sql += fmt.Sprintf("%v'%v' as coupon_amount", comma, *h.CouponAmount)
	_sql += fmt.Sprintf("%v'%v' as coupon_cashable_amount", comma, *h.CouponCashableAmount)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from user_rewards "
	_sql += fmt.Sprintf("where app_id='%v' and user_id='%v'", *h.AppID, *h.UserID)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create userreward: %v", err)
	}
	return nil
}

func (h *Handler) CreateUserReward(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createUserReward(_ctx, tx)
	})
}
