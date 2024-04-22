package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql   string
	appID string
}

func (h *updateHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "update app_configs "
	_sql += "set "
	comma = ", "
	_sql += fmt.Sprintf("updated_at = %v", now)

	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v", comma, *h.StartAt)
	}

	_sql += " where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.id != %v and di.end_at=0 and di.deleted_at=0", h.appID, *h.ID)
	_sql += " limit 1)"

	if h.StartAt != nil {
		_sql += " and not exists ("
		_sql += " select 1 from (select * from app_configs) as di "
		_sql += fmt.Sprintf("where di.app_id='%v' and di.deleted_at=0 and di.end_at!=0 and %v < di.end_at",
			h.appID, *h.StartAt)
		_sql += " limit 1)"
	}

	h.sql = _sql
}

func (h *updateHandler) updateAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail update appconfig: %v", err)
	}
	return nil
}

func (h *Handler) UpdateAppConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid appconfig")
	}
	h.ID = &info.ID
	handler.appID = info.AppID

	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateAppConfig(_ctx, tx)
	})
}
