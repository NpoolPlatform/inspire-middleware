package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update coin_configs "
	if h.MaxValue != nil {
		_sql += fmt.Sprintf("%vmax_value = '%v', ", set, *h.MaxValue)
		set = ""
	}
	if h.Allocated != nil {
		_sql += fmt.Sprintf("%vallocated = '%v', ", set, *h.Allocated)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateCoinConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateCoinConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetCoinConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid coinconfig")
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCoinConfig(_ctx, tx)
	})
}
