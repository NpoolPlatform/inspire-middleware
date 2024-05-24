package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	sql     string
	appID   string
	eventID string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update task_configs "
	if h.EventID != nil {
		_sql += fmt.Sprintf("%vevent_id = '%v', ", set, *h.EventID)
		set = ""
	}
	if h.TaskType != nil {
		_sql += fmt.Sprintf("%vtask_type = '%v', ", set, *h.TaskType)
		set = ""
	}
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.TaskDesc != nil {
		_sql += fmt.Sprintf("%vtask_desc = '%v', ", set, *h.TaskDesc)
		set = ""
	}
	if h.StepGuide != nil {
		_sql += fmt.Sprintf("%vstep_guide = '%v', ", set, *h.StepGuide)
		set = ""
	}
	if h.RecommendMessage != nil {
		_sql += fmt.Sprintf("%vrecommend_message = '%v', ", set, *h.RecommendMessage)
		set = ""
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v`index` = '%v', ", set, *h.Index)
		set = ""
	}
	if h.LastTaskID != nil {
		_sql += fmt.Sprintf("%vlast_task_id = '%v', ", set, *h.LastTaskID)
		set = ""
	}
	if h.MaxRewardCount != nil {
		_sql += fmt.Sprintf("%vmax_reward_count = '%v', ", set, *h.MaxRewardCount)
		set = ""
	}
	if h.CooldownSecord != nil {
		_sql += fmt.Sprintf("%vcooldown_secord = '%v', ", set, *h.CooldownSecord)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from task_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.event_id = '%v' and di.task_type = '%v' and di.id != %v and deleted_at=0", h.appID, h.eventID, h.TaskType.String(), *h.ID)
	_sql += " limit 1)"

	if h.EventID != nil {
		_sql += "and not exists ("
		_sql += "select 1 from (select * from task_users) as di "
		_sql += fmt.Sprintf("where di.app_id = '%v' and di.task_id = '%v' and di.event_id = '%v' and deleted_at=0", h.appID, *h.EntID, h.eventID)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateTaskConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTaskConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetTaskConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid taskconfig")
	}

	if h.TaskType == nil {
		h.TaskType = &info.TaskType
	}

	if h.EntID == nil {
		id := uuid.MustParse(info.EntID)
		h.EntID = &id
	}
	handler.eventID = info.EventID
	handler.appID = info.AppID

	if err := handler.constructSQL(); err != nil {
		return err
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTaskConfig(_ctx, tx)
	})
}
