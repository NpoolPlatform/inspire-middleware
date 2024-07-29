package coin

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
	_sql := "insert into event_coins "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "event_id"
	_sql += comma + "coin_config_id"
	_sql += comma + "coin_value"
	_sql += comma + "coin_pre_usd"
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
	_sql += fmt.Sprintf("%v'%v' as event_id", comma, *h.EventID)
	_sql += fmt.Sprintf("%v'%v' as coin_config_id", comma, *h.CoinConfigID)
	_sql += fmt.Sprintf("%v'%v' as coin_value", comma, *h.CoinValue)
	_sql += fmt.Sprintf("%v'%v' as coin_pre_usd", comma, *h.CoinPreUSD)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from event_coins "
	_sql += fmt.Sprintf("where app_id='%v' and event_id='%v' and coin_config_id='%v' and deleted_at=0", *h.AppID, *h.EventID, *h.CoinConfigID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from events "
	_sql += fmt.Sprintf("where app_id='%v' and ent_id='%v' and deleted_at=0", *h.AppID, *h.EventID)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createEventCoin(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create eventcoin: %v", err)
	}
	return nil
}

func (h *Handler) CreateEventCoin(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createEventCoin(_ctx, tx)
	})
}
