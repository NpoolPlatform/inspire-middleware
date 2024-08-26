package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

type updateHandler struct {
	*Handler
	sql  string
	info *npool.TaskUser
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update task_users "
	if h.TaskState != nil {
		_sql += fmt.Sprintf("%vtask_state = '%v', ", set, h.TaskState.String())
		set = ""
	}
	if h.RewardState != nil {
		_sql += fmt.Sprintf("%vreward_state = '%v', ", set, h.RewardState.String())
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

func (h *updateHandler) updateTaskUser(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) validTaskState() error {
	if h.TaskState == nil {
		return nil
	}
	if h.info.TaskState == basetypes.TaskState_Done && *h.TaskState == basetypes.TaskState_InProgress {
		return wlog.Errorf("invalid taskstate")
	}
	return nil
}

func (h *updateHandler) validRewardState() error {
	if h.RewardState == nil {
		return nil
	}
	switch h.info.RewardState {
	case basetypes.RewardState_Issued:
		fallthrough //nolint
	case basetypes.RewardState_Revoked:
		if *h.RewardState == basetypes.RewardState_Issued {
			return wlog.Errorf("invalid rewardstate")
		}
	}
	return nil
}

func (h *Handler) UpdateTaskUser(ctx context.Context) error {
	info, err := h.GetTaskUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid taskuser")
	}
	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
		info:    info,
	}

	if err := handler.validTaskState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validRewardState(); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTaskUser(_ctx, tx)
	})
}
