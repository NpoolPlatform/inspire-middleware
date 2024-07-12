package config

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	devicecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/task/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	enttaskconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/taskconfig"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TaskConfigSelect
	stmCount  *ent.TaskConfigSelect
	infos     []*npool.TaskConfig
	total     uint32
}

func (h *queryHandler) selectTaskConfig(stm *ent.TaskConfigQuery) {
	h.stmSelect = stm.Select(enttaskconfig.FieldID)
}

func (h *queryHandler) queryTaskConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TaskConfig.Query().Where(enttaskconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttaskconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttaskconfig.EntID(*h.EntID))
	}
	h.selectTaskConfig(stm)
	return nil
}

func (h *queryHandler) queryTaskConfigs(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.TaskConfig.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectTaskConfig(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(enttaskconfig.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttaskconfig.FieldEntID),
			t1.C(enttaskconfig.FieldEntID),
		).
		AppendSelect(
			t1.C(enttaskconfig.FieldEntID),
			t1.C(enttaskconfig.FieldAppID),
			t1.C(enttaskconfig.FieldEventID),
			t1.C(enttaskconfig.FieldTaskType),
			t1.C(enttaskconfig.FieldName),
			t1.C(enttaskconfig.FieldTaskDesc),
			t1.C(enttaskconfig.FieldStepGuide),
			t1.C(enttaskconfig.FieldRecommendMessage),
			t1.C(enttaskconfig.FieldIndex),
			t1.C(enttaskconfig.FieldLastTaskID),
			t1.C(enttaskconfig.FieldMaxRewardCount),
			t1.C(enttaskconfig.FieldCooldownSecord),
			t1.C(enttaskconfig.FieldCreatedAt),
			t1.C(enttaskconfig.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TaskType = basetypes.TaskType(basetypes.TaskType_value[info.TaskTypeStr])
	}
}

func (h *Handler) GetTaskConfig(ctx context.Context) (*npool.TaskConfig, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTaskConfig(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetTaskConfigs(ctx context.Context) ([]*npool.TaskConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTaskConfigs(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
