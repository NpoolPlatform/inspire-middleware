package user

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
	_sql := "insert into task_users "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "task_id"
	_sql += comma + "event_id"
	_sql += comma + "task_state"
	_sql += comma + "reward_state"
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
	_sql += fmt.Sprintf("%v'%v' as task_id", comma, *h.TaskID)
	_sql += fmt.Sprintf("%v'%v' as event_id", comma, *h.EventID)
	_sql += fmt.Sprintf("%v'%v' as task_state", comma, h.TaskState.String())
	_sql += fmt.Sprintf("%v'%v' as reward_state", comma, h.RewardState.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from (select * from task_configs) as di "
	_sql += fmt.Sprintf("where di.ent_id = '%v' and di.app_id = '%v' and di.event_id = '%v' and di.deleted_at=0", *h.TaskID, *h.AppID, *h.EventID)
	_sql += " limit 1)"
	_sql += " and exists ("
	_sql += "select 1 from (select * from events) as di "
	_sql += fmt.Sprintf("where di.ent_id = '%v' and di.app_id = '%v' and di.deleted_at=0", *h.EventID, *h.AppID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createTaskUser(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create taskuser: %v", err)
	}
	return nil
}

func (h *Handler) CreateTaskUser(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createTaskUser(_ctx, tx)
	})
}
