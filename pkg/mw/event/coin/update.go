package coin

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update event_coins "
	if h.CoinValue != nil {
		_sql += fmt.Sprintf("%vcoin_value = '%v', ", set, *h.CoinValue)
		set = ""
	}
	if h.CoinPreUSD != nil {
		_sql += fmt.Sprintf("%vcoin_pre_usd = '%v', ", set, *h.CoinPreUSD)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v and app_id = '%v' ", *h.ID, *h.AppID)
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateEventCoin(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateEventCoin(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetEventCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid eventcoin")
	}
	h.ID = &info.ID
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateEventCoin(_ctx, tx)
	})
}