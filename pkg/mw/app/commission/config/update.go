package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sql        string
	appID      string
	settleType string
	level      uint32
	disabled   bool
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_commission_configs "
	if h.AmountOrPercent != nil {
		_sql += fmt.Sprintf("%vamount_or_percent = '%v', ", set, *h.AmountOrPercent)
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if h.ThresholdAmount != nil {
		_sql += fmt.Sprintf("%vthreshold_amount = '%v', ", set, *h.ThresholdAmount)
		set = ""
	}
	if h.Invites != nil {
		_sql += fmt.Sprintf("%vinvites = %v, ", set, *h.Invites)
		set = ""
	}
	if h.Disabled != nil {
		_sql += fmt.Sprintf("%vdisabled = %v, ", set, *h.Disabled)
		set = ""
	}
	if h.Level != nil {
		_sql += fmt.Sprintf("%vlevel = %v, ", set, *h.Level)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += " where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_commission_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.level = %v and di.id != %v and di.end_at=0 and di.deleted_at=0", h.appID, h.level, *h.ID)
	_sql += " limit 1)"

	if !h.disabled {
		_sql += " and exists ("
		_sql += " select 1 from app_configs "
		_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and deleted_at=0 and %v < max_level",
			h.appID, h.level)
		_sql += " limit 1)"
	}

	if h.StartAt != nil {
		_sql += " and not exists ("
		_sql += " select 1 from (select * from app_commission_configs) as di "
		_sql += fmt.Sprintf("where di.app_id='%v' and di.settle_type='%v' and di.level=%v and di.deleted_at=0 and di.end_at!=0 and %v < di.end_at",
			h.appID, h.settleType, h.level, *h.StartAt)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateCommissionConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail update appcommissionconfig: %v", err)
	}
	return nil
}

func (h *Handler) UpdateCommissionConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid appcommissionconfig")
	}

	h.ID = &info.ID
	handler.level = info.Level
	handler.disabled = info.Disabled
	if h.Level != nil {
		handler.level = *h.Level
	}
	if h.Disabled != nil {
		handler.disabled = *h.Disabled
	}

	handler.appID = info.AppID
	handler.settleType = info.SettleTypeStr

	if err := handler.constructSQL(); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCommissionConfig(_ctx, tx)
	})
}
