package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
)

type updateHandler struct {
	*Handler
	sql        string
	appID      string
	goodID     string
	appGoodID  string
	settleType string
}

func (h *updateHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "update app_good_commission_configs "
	_sql += "set "
	comma = ", "
	_sql += fmt.Sprintf("updated_at = %v", now)
	if h.AmountOrPercent != nil {
		_sql += fmt.Sprintf("%vamount_or_percent = '%v'", comma, *h.AmountOrPercent)
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v", comma, *h.StartAt)
	}
	if h.ThresholdAmount != nil {
		_sql += fmt.Sprintf("%vthreshold_amount = '%v'", comma, *h.ThresholdAmount)
	}
	if h.Invites != nil {
		_sql += fmt.Sprintf("%vinvites = %v", comma, *h.Invites)
	}
	if h.Disabled != nil {
		_sql += fmt.Sprintf("%vdisabled = %v", comma, *h.Disabled)
	}
	if h.Level != nil {
		_sql += fmt.Sprintf("%vlevel = %v", comma, *h.Level)
	}
	_sql += " where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_good_commission_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.good_id = '%v' and di.app_good_id = '%v' and di.level = '%v' and di.id != %v and di.end_at=0 and di.deleted_at=0",
		h.appID, h.goodID, h.appGoodID, *h.Level, *h.ID)
	_sql += " limit 1)"

	if !*h.Disabled {
		_sql += " and exists ("
		_sql += " select 1 from app_configs "
		_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and deleted_at=0 and %v < max_level",
			h.appID, *h.Level)
		_sql += " limit 1)"
	}

	if h.StartAt != nil {
		_sql += " and not exists ("
		_sql += " select 1 from (select * from app_good_commission_configs) as di "
		_sql += fmt.Sprintf("where di.app_id='%v' and di.app_good_id='%v' and di.settle_type='%v' and di.level=%v and di.deleted_at=0 and di.end_at!=0 and %v < di.end_at",
			h.appID, h.appGoodID, h.settleType, *h.Level, *h.StartAt)
		_sql += " limit 1)"
	}

	h.sql = _sql
}

func (h *updateHandler) updateCommissionConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail update appgoodcommissionconfig: %v", err)
	}
	return nil
}

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	handler := &updateHandler{
		Handler: h,
	}
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid appgoodcommissionconfig")
	}
	h.ID = &info.ID

	if h.Level == nil {
		h.Level = &info.Level
	}
	if h.Disabled == nil {
		h.Disabled = &info.Disabled
	}

	handler.appID = info.AppID
	handler.goodID = info.GoodID
	handler.appGoodID = info.AppGoodID
	handler.settleType = info.SettleTypeStr

	handler.constructSQL()

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCommissionConfig(_ctx, tx)
	})
	if err != nil {
		return nil, err
	}

	return h.GetCommissionConfig(ctx)
}
