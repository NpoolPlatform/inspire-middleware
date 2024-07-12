// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/taskconfig"
	"github.com/google/uuid"
)

// TaskConfigUpdate is the builder for updating TaskConfig entities.
type TaskConfigUpdate struct {
	config
	hooks     []Hook
	mutation  *TaskConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TaskConfigUpdate builder.
func (tcu *TaskConfigUpdate) Where(ps ...predicate.TaskConfig) *TaskConfigUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetCreatedAt sets the "created_at" field.
func (tcu *TaskConfigUpdate) SetCreatedAt(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetCreatedAt()
	tcu.mutation.SetCreatedAt(u)
	return tcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableCreatedAt(u *uint32) *TaskConfigUpdate {
	if u != nil {
		tcu.SetCreatedAt(*u)
	}
	return tcu
}

// AddCreatedAt adds u to the "created_at" field.
func (tcu *TaskConfigUpdate) AddCreatedAt(u int32) *TaskConfigUpdate {
	tcu.mutation.AddCreatedAt(u)
	return tcu
}

// SetUpdatedAt sets the "updated_at" field.
func (tcu *TaskConfigUpdate) SetUpdatedAt(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetUpdatedAt()
	tcu.mutation.SetUpdatedAt(u)
	return tcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tcu *TaskConfigUpdate) AddUpdatedAt(u int32) *TaskConfigUpdate {
	tcu.mutation.AddUpdatedAt(u)
	return tcu
}

// SetDeletedAt sets the "deleted_at" field.
func (tcu *TaskConfigUpdate) SetDeletedAt(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetDeletedAt()
	tcu.mutation.SetDeletedAt(u)
	return tcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableDeletedAt(u *uint32) *TaskConfigUpdate {
	if u != nil {
		tcu.SetDeletedAt(*u)
	}
	return tcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tcu *TaskConfigUpdate) AddDeletedAt(u int32) *TaskConfigUpdate {
	tcu.mutation.AddDeletedAt(u)
	return tcu
}

// SetEntID sets the "ent_id" field.
func (tcu *TaskConfigUpdate) SetEntID(u uuid.UUID) *TaskConfigUpdate {
	tcu.mutation.SetEntID(u)
	return tcu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableEntID(u *uuid.UUID) *TaskConfigUpdate {
	if u != nil {
		tcu.SetEntID(*u)
	}
	return tcu
}

// SetAppID sets the "app_id" field.
func (tcu *TaskConfigUpdate) SetAppID(u uuid.UUID) *TaskConfigUpdate {
	tcu.mutation.SetAppID(u)
	return tcu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableAppID(u *uuid.UUID) *TaskConfigUpdate {
	if u != nil {
		tcu.SetAppID(*u)
	}
	return tcu
}

// ClearAppID clears the value of the "app_id" field.
func (tcu *TaskConfigUpdate) ClearAppID() *TaskConfigUpdate {
	tcu.mutation.ClearAppID()
	return tcu
}

// SetEventID sets the "event_id" field.
func (tcu *TaskConfigUpdate) SetEventID(u uuid.UUID) *TaskConfigUpdate {
	tcu.mutation.SetEventID(u)
	return tcu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableEventID(u *uuid.UUID) *TaskConfigUpdate {
	if u != nil {
		tcu.SetEventID(*u)
	}
	return tcu
}

// ClearEventID clears the value of the "event_id" field.
func (tcu *TaskConfigUpdate) ClearEventID() *TaskConfigUpdate {
	tcu.mutation.ClearEventID()
	return tcu
}

// SetTaskType sets the "task_type" field.
func (tcu *TaskConfigUpdate) SetTaskType(s string) *TaskConfigUpdate {
	tcu.mutation.SetTaskType(s)
	return tcu
}

// SetNillableTaskType sets the "task_type" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableTaskType(s *string) *TaskConfigUpdate {
	if s != nil {
		tcu.SetTaskType(*s)
	}
	return tcu
}

// ClearTaskType clears the value of the "task_type" field.
func (tcu *TaskConfigUpdate) ClearTaskType() *TaskConfigUpdate {
	tcu.mutation.ClearTaskType()
	return tcu
}

// SetName sets the "name" field.
func (tcu *TaskConfigUpdate) SetName(s string) *TaskConfigUpdate {
	tcu.mutation.SetName(s)
	return tcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableName(s *string) *TaskConfigUpdate {
	if s != nil {
		tcu.SetName(*s)
	}
	return tcu
}

// ClearName clears the value of the "name" field.
func (tcu *TaskConfigUpdate) ClearName() *TaskConfigUpdate {
	tcu.mutation.ClearName()
	return tcu
}

// SetTaskDesc sets the "task_desc" field.
func (tcu *TaskConfigUpdate) SetTaskDesc(s string) *TaskConfigUpdate {
	tcu.mutation.SetTaskDesc(s)
	return tcu
}

// SetNillableTaskDesc sets the "task_desc" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableTaskDesc(s *string) *TaskConfigUpdate {
	if s != nil {
		tcu.SetTaskDesc(*s)
	}
	return tcu
}

// ClearTaskDesc clears the value of the "task_desc" field.
func (tcu *TaskConfigUpdate) ClearTaskDesc() *TaskConfigUpdate {
	tcu.mutation.ClearTaskDesc()
	return tcu
}

// SetStepGuide sets the "step_guide" field.
func (tcu *TaskConfigUpdate) SetStepGuide(s string) *TaskConfigUpdate {
	tcu.mutation.SetStepGuide(s)
	return tcu
}

// SetNillableStepGuide sets the "step_guide" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableStepGuide(s *string) *TaskConfigUpdate {
	if s != nil {
		tcu.SetStepGuide(*s)
	}
	return tcu
}

// ClearStepGuide clears the value of the "step_guide" field.
func (tcu *TaskConfigUpdate) ClearStepGuide() *TaskConfigUpdate {
	tcu.mutation.ClearStepGuide()
	return tcu
}

// SetRecommendMessage sets the "recommend_message" field.
func (tcu *TaskConfigUpdate) SetRecommendMessage(s string) *TaskConfigUpdate {
	tcu.mutation.SetRecommendMessage(s)
	return tcu
}

// SetNillableRecommendMessage sets the "recommend_message" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableRecommendMessage(s *string) *TaskConfigUpdate {
	if s != nil {
		tcu.SetRecommendMessage(*s)
	}
	return tcu
}

// ClearRecommendMessage clears the value of the "recommend_message" field.
func (tcu *TaskConfigUpdate) ClearRecommendMessage() *TaskConfigUpdate {
	tcu.mutation.ClearRecommendMessage()
	return tcu
}

// SetIndex sets the "index" field.
func (tcu *TaskConfigUpdate) SetIndex(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetIndex()
	tcu.mutation.SetIndex(u)
	return tcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableIndex(u *uint32) *TaskConfigUpdate {
	if u != nil {
		tcu.SetIndex(*u)
	}
	return tcu
}

// AddIndex adds u to the "index" field.
func (tcu *TaskConfigUpdate) AddIndex(u int32) *TaskConfigUpdate {
	tcu.mutation.AddIndex(u)
	return tcu
}

// ClearIndex clears the value of the "index" field.
func (tcu *TaskConfigUpdate) ClearIndex() *TaskConfigUpdate {
	tcu.mutation.ClearIndex()
	return tcu
}

// SetLastTaskID sets the "last_task_id" field.
func (tcu *TaskConfigUpdate) SetLastTaskID(u uuid.UUID) *TaskConfigUpdate {
	tcu.mutation.SetLastTaskID(u)
	return tcu
}

// SetNillableLastTaskID sets the "last_task_id" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableLastTaskID(u *uuid.UUID) *TaskConfigUpdate {
	if u != nil {
		tcu.SetLastTaskID(*u)
	}
	return tcu
}

// ClearLastTaskID clears the value of the "last_task_id" field.
func (tcu *TaskConfigUpdate) ClearLastTaskID() *TaskConfigUpdate {
	tcu.mutation.ClearLastTaskID()
	return tcu
}

// SetMaxRewardCount sets the "max_reward_count" field.
func (tcu *TaskConfigUpdate) SetMaxRewardCount(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetMaxRewardCount()
	tcu.mutation.SetMaxRewardCount(u)
	return tcu
}

// SetNillableMaxRewardCount sets the "max_reward_count" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableMaxRewardCount(u *uint32) *TaskConfigUpdate {
	if u != nil {
		tcu.SetMaxRewardCount(*u)
	}
	return tcu
}

// AddMaxRewardCount adds u to the "max_reward_count" field.
func (tcu *TaskConfigUpdate) AddMaxRewardCount(u int32) *TaskConfigUpdate {
	tcu.mutation.AddMaxRewardCount(u)
	return tcu
}

// ClearMaxRewardCount clears the value of the "max_reward_count" field.
func (tcu *TaskConfigUpdate) ClearMaxRewardCount() *TaskConfigUpdate {
	tcu.mutation.ClearMaxRewardCount()
	return tcu
}

// SetCooldownSecord sets the "cooldown_secord" field.
func (tcu *TaskConfigUpdate) SetCooldownSecord(u uint32) *TaskConfigUpdate {
	tcu.mutation.ResetCooldownSecord()
	tcu.mutation.SetCooldownSecord(u)
	return tcu
}

// SetNillableCooldownSecord sets the "cooldown_secord" field if the given value is not nil.
func (tcu *TaskConfigUpdate) SetNillableCooldownSecord(u *uint32) *TaskConfigUpdate {
	if u != nil {
		tcu.SetCooldownSecord(*u)
	}
	return tcu
}

// AddCooldownSecord adds u to the "cooldown_secord" field.
func (tcu *TaskConfigUpdate) AddCooldownSecord(u int32) *TaskConfigUpdate {
	tcu.mutation.AddCooldownSecord(u)
	return tcu
}

// ClearCooldownSecord clears the value of the "cooldown_secord" field.
func (tcu *TaskConfigUpdate) ClearCooldownSecord() *TaskConfigUpdate {
	tcu.mutation.ClearCooldownSecord()
	return tcu
}

// Mutation returns the TaskConfigMutation object of the builder.
func (tcu *TaskConfigUpdate) Mutation() *TaskConfigMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TaskConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tcu.defaults(); err != nil {
		return 0, err
	}
	if len(tcu.hooks) == 0 {
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TaskConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TaskConfigUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TaskConfigUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TaskConfigUpdate) defaults() error {
	if _, ok := tcu.mutation.UpdatedAt(); !ok {
		if taskconfig.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized taskconfig.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := taskconfig.UpdateDefaultUpdatedAt()
		tcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcu *TaskConfigUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskConfigUpdate {
	tcu.modifiers = append(tcu.modifiers, modifiers...)
	return tcu
}

func (tcu *TaskConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskconfig.Table,
			Columns: taskconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: taskconfig.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCreatedAt,
		})
	}
	if value, ok := tcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCreatedAt,
		})
	}
	if value, ok := tcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldUpdatedAt,
		})
	}
	if value, ok := tcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldUpdatedAt,
		})
	}
	if value, ok := tcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldDeletedAt,
		})
	}
	if value, ok := tcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldDeletedAt,
		})
	}
	if value, ok := tcu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldEntID,
		})
	}
	if value, ok := tcu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldAppID,
		})
	}
	if tcu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldAppID,
		})
	}
	if value, ok := tcu.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldEventID,
		})
	}
	if tcu.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldEventID,
		})
	}
	if value, ok := tcu.mutation.TaskType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldTaskType,
		})
	}
	if tcu.mutation.TaskTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldTaskType,
		})
	}
	if value, ok := tcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldName,
		})
	}
	if tcu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldName,
		})
	}
	if value, ok := tcu.mutation.TaskDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldTaskDesc,
		})
	}
	if tcu.mutation.TaskDescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldTaskDesc,
		})
	}
	if value, ok := tcu.mutation.StepGuide(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldStepGuide,
		})
	}
	if tcu.mutation.StepGuideCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldStepGuide,
		})
	}
	if value, ok := tcu.mutation.RecommendMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldRecommendMessage,
		})
	}
	if tcu.mutation.RecommendMessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldRecommendMessage,
		})
	}
	if value, ok := tcu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldIndex,
		})
	}
	if value, ok := tcu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldIndex,
		})
	}
	if tcu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldIndex,
		})
	}
	if value, ok := tcu.mutation.LastTaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldLastTaskID,
		})
	}
	if tcu.mutation.LastTaskIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldLastTaskID,
		})
	}
	if value, ok := tcu.mutation.MaxRewardCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if value, ok := tcu.mutation.AddedMaxRewardCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if tcu.mutation.MaxRewardCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if value, ok := tcu.mutation.CooldownSecord(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	if value, ok := tcu.mutation.AddedCooldownSecord(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	if tcu.mutation.CooldownSecordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	_spec.Modifiers = tcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TaskConfigUpdateOne is the builder for updating a single TaskConfig entity.
type TaskConfigUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TaskConfigMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tcuo *TaskConfigUpdateOne) SetCreatedAt(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetCreatedAt()
	tcuo.mutation.SetCreatedAt(u)
	return tcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableCreatedAt(u *uint32) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetCreatedAt(*u)
	}
	return tcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tcuo *TaskConfigUpdateOne) AddCreatedAt(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddCreatedAt(u)
	return tcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tcuo *TaskConfigUpdateOne) SetUpdatedAt(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetUpdatedAt()
	tcuo.mutation.SetUpdatedAt(u)
	return tcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tcuo *TaskConfigUpdateOne) AddUpdatedAt(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddUpdatedAt(u)
	return tcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tcuo *TaskConfigUpdateOne) SetDeletedAt(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetDeletedAt()
	tcuo.mutation.SetDeletedAt(u)
	return tcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableDeletedAt(u *uint32) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetDeletedAt(*u)
	}
	return tcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tcuo *TaskConfigUpdateOne) AddDeletedAt(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddDeletedAt(u)
	return tcuo
}

// SetEntID sets the "ent_id" field.
func (tcuo *TaskConfigUpdateOne) SetEntID(u uuid.UUID) *TaskConfigUpdateOne {
	tcuo.mutation.SetEntID(u)
	return tcuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableEntID(u *uuid.UUID) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetEntID(*u)
	}
	return tcuo
}

// SetAppID sets the "app_id" field.
func (tcuo *TaskConfigUpdateOne) SetAppID(u uuid.UUID) *TaskConfigUpdateOne {
	tcuo.mutation.SetAppID(u)
	return tcuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableAppID(u *uuid.UUID) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetAppID(*u)
	}
	return tcuo
}

// ClearAppID clears the value of the "app_id" field.
func (tcuo *TaskConfigUpdateOne) ClearAppID() *TaskConfigUpdateOne {
	tcuo.mutation.ClearAppID()
	return tcuo
}

// SetEventID sets the "event_id" field.
func (tcuo *TaskConfigUpdateOne) SetEventID(u uuid.UUID) *TaskConfigUpdateOne {
	tcuo.mutation.SetEventID(u)
	return tcuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableEventID(u *uuid.UUID) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetEventID(*u)
	}
	return tcuo
}

// ClearEventID clears the value of the "event_id" field.
func (tcuo *TaskConfigUpdateOne) ClearEventID() *TaskConfigUpdateOne {
	tcuo.mutation.ClearEventID()
	return tcuo
}

// SetTaskType sets the "task_type" field.
func (tcuo *TaskConfigUpdateOne) SetTaskType(s string) *TaskConfigUpdateOne {
	tcuo.mutation.SetTaskType(s)
	return tcuo
}

// SetNillableTaskType sets the "task_type" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableTaskType(s *string) *TaskConfigUpdateOne {
	if s != nil {
		tcuo.SetTaskType(*s)
	}
	return tcuo
}

// ClearTaskType clears the value of the "task_type" field.
func (tcuo *TaskConfigUpdateOne) ClearTaskType() *TaskConfigUpdateOne {
	tcuo.mutation.ClearTaskType()
	return tcuo
}

// SetName sets the "name" field.
func (tcuo *TaskConfigUpdateOne) SetName(s string) *TaskConfigUpdateOne {
	tcuo.mutation.SetName(s)
	return tcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableName(s *string) *TaskConfigUpdateOne {
	if s != nil {
		tcuo.SetName(*s)
	}
	return tcuo
}

// ClearName clears the value of the "name" field.
func (tcuo *TaskConfigUpdateOne) ClearName() *TaskConfigUpdateOne {
	tcuo.mutation.ClearName()
	return tcuo
}

// SetTaskDesc sets the "task_desc" field.
func (tcuo *TaskConfigUpdateOne) SetTaskDesc(s string) *TaskConfigUpdateOne {
	tcuo.mutation.SetTaskDesc(s)
	return tcuo
}

// SetNillableTaskDesc sets the "task_desc" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableTaskDesc(s *string) *TaskConfigUpdateOne {
	if s != nil {
		tcuo.SetTaskDesc(*s)
	}
	return tcuo
}

// ClearTaskDesc clears the value of the "task_desc" field.
func (tcuo *TaskConfigUpdateOne) ClearTaskDesc() *TaskConfigUpdateOne {
	tcuo.mutation.ClearTaskDesc()
	return tcuo
}

// SetStepGuide sets the "step_guide" field.
func (tcuo *TaskConfigUpdateOne) SetStepGuide(s string) *TaskConfigUpdateOne {
	tcuo.mutation.SetStepGuide(s)
	return tcuo
}

// SetNillableStepGuide sets the "step_guide" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableStepGuide(s *string) *TaskConfigUpdateOne {
	if s != nil {
		tcuo.SetStepGuide(*s)
	}
	return tcuo
}

// ClearStepGuide clears the value of the "step_guide" field.
func (tcuo *TaskConfigUpdateOne) ClearStepGuide() *TaskConfigUpdateOne {
	tcuo.mutation.ClearStepGuide()
	return tcuo
}

// SetRecommendMessage sets the "recommend_message" field.
func (tcuo *TaskConfigUpdateOne) SetRecommendMessage(s string) *TaskConfigUpdateOne {
	tcuo.mutation.SetRecommendMessage(s)
	return tcuo
}

// SetNillableRecommendMessage sets the "recommend_message" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableRecommendMessage(s *string) *TaskConfigUpdateOne {
	if s != nil {
		tcuo.SetRecommendMessage(*s)
	}
	return tcuo
}

// ClearRecommendMessage clears the value of the "recommend_message" field.
func (tcuo *TaskConfigUpdateOne) ClearRecommendMessage() *TaskConfigUpdateOne {
	tcuo.mutation.ClearRecommendMessage()
	return tcuo
}

// SetIndex sets the "index" field.
func (tcuo *TaskConfigUpdateOne) SetIndex(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetIndex()
	tcuo.mutation.SetIndex(u)
	return tcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableIndex(u *uint32) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetIndex(*u)
	}
	return tcuo
}

// AddIndex adds u to the "index" field.
func (tcuo *TaskConfigUpdateOne) AddIndex(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddIndex(u)
	return tcuo
}

// ClearIndex clears the value of the "index" field.
func (tcuo *TaskConfigUpdateOne) ClearIndex() *TaskConfigUpdateOne {
	tcuo.mutation.ClearIndex()
	return tcuo
}

// SetLastTaskID sets the "last_task_id" field.
func (tcuo *TaskConfigUpdateOne) SetLastTaskID(u uuid.UUID) *TaskConfigUpdateOne {
	tcuo.mutation.SetLastTaskID(u)
	return tcuo
}

// SetNillableLastTaskID sets the "last_task_id" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableLastTaskID(u *uuid.UUID) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetLastTaskID(*u)
	}
	return tcuo
}

// ClearLastTaskID clears the value of the "last_task_id" field.
func (tcuo *TaskConfigUpdateOne) ClearLastTaskID() *TaskConfigUpdateOne {
	tcuo.mutation.ClearLastTaskID()
	return tcuo
}

// SetMaxRewardCount sets the "max_reward_count" field.
func (tcuo *TaskConfigUpdateOne) SetMaxRewardCount(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetMaxRewardCount()
	tcuo.mutation.SetMaxRewardCount(u)
	return tcuo
}

// SetNillableMaxRewardCount sets the "max_reward_count" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableMaxRewardCount(u *uint32) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetMaxRewardCount(*u)
	}
	return tcuo
}

// AddMaxRewardCount adds u to the "max_reward_count" field.
func (tcuo *TaskConfigUpdateOne) AddMaxRewardCount(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddMaxRewardCount(u)
	return tcuo
}

// ClearMaxRewardCount clears the value of the "max_reward_count" field.
func (tcuo *TaskConfigUpdateOne) ClearMaxRewardCount() *TaskConfigUpdateOne {
	tcuo.mutation.ClearMaxRewardCount()
	return tcuo
}

// SetCooldownSecord sets the "cooldown_secord" field.
func (tcuo *TaskConfigUpdateOne) SetCooldownSecord(u uint32) *TaskConfigUpdateOne {
	tcuo.mutation.ResetCooldownSecord()
	tcuo.mutation.SetCooldownSecord(u)
	return tcuo
}

// SetNillableCooldownSecord sets the "cooldown_secord" field if the given value is not nil.
func (tcuo *TaskConfigUpdateOne) SetNillableCooldownSecord(u *uint32) *TaskConfigUpdateOne {
	if u != nil {
		tcuo.SetCooldownSecord(*u)
	}
	return tcuo
}

// AddCooldownSecord adds u to the "cooldown_secord" field.
func (tcuo *TaskConfigUpdateOne) AddCooldownSecord(u int32) *TaskConfigUpdateOne {
	tcuo.mutation.AddCooldownSecord(u)
	return tcuo
}

// ClearCooldownSecord clears the value of the "cooldown_secord" field.
func (tcuo *TaskConfigUpdateOne) ClearCooldownSecord() *TaskConfigUpdateOne {
	tcuo.mutation.ClearCooldownSecord()
	return tcuo
}

// Mutation returns the TaskConfigMutation object of the builder.
func (tcuo *TaskConfigUpdateOne) Mutation() *TaskConfigMutation {
	return tcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TaskConfigUpdateOne) Select(field string, fields ...string) *TaskConfigUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TaskConfig entity.
func (tcuo *TaskConfigUpdateOne) Save(ctx context.Context) (*TaskConfig, error) {
	var (
		err  error
		node *TaskConfig
	)
	if err := tcuo.defaults(); err != nil {
		return nil, err
	}
	if len(tcuo.hooks) == 0 {
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TaskConfig)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TaskConfigMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TaskConfigUpdateOne) SaveX(ctx context.Context) *TaskConfig {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TaskConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TaskConfigUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TaskConfigUpdateOne) defaults() error {
	if _, ok := tcuo.mutation.UpdatedAt(); !ok {
		if taskconfig.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized taskconfig.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := taskconfig.UpdateDefaultUpdatedAt()
		tcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcuo *TaskConfigUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskConfigUpdateOne {
	tcuo.modifiers = append(tcuo.modifiers, modifiers...)
	return tcuo
}

func (tcuo *TaskConfigUpdateOne) sqlSave(ctx context.Context) (_node *TaskConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskconfig.Table,
			Columns: taskconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: taskconfig.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskconfig.FieldID)
		for _, f := range fields {
			if !taskconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCreatedAt,
		})
	}
	if value, ok := tcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCreatedAt,
		})
	}
	if value, ok := tcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldUpdatedAt,
		})
	}
	if value, ok := tcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldUpdatedAt,
		})
	}
	if value, ok := tcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldDeletedAt,
		})
	}
	if value, ok := tcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldDeletedAt,
		})
	}
	if value, ok := tcuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldEntID,
		})
	}
	if value, ok := tcuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldAppID,
		})
	}
	if tcuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldAppID,
		})
	}
	if value, ok := tcuo.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldEventID,
		})
	}
	if tcuo.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldEventID,
		})
	}
	if value, ok := tcuo.mutation.TaskType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldTaskType,
		})
	}
	if tcuo.mutation.TaskTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldTaskType,
		})
	}
	if value, ok := tcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldName,
		})
	}
	if tcuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldName,
		})
	}
	if value, ok := tcuo.mutation.TaskDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldTaskDesc,
		})
	}
	if tcuo.mutation.TaskDescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldTaskDesc,
		})
	}
	if value, ok := tcuo.mutation.StepGuide(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldStepGuide,
		})
	}
	if tcuo.mutation.StepGuideCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldStepGuide,
		})
	}
	if value, ok := tcuo.mutation.RecommendMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskconfig.FieldRecommendMessage,
		})
	}
	if tcuo.mutation.RecommendMessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: taskconfig.FieldRecommendMessage,
		})
	}
	if value, ok := tcuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldIndex,
		})
	}
	if value, ok := tcuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldIndex,
		})
	}
	if tcuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldIndex,
		})
	}
	if value, ok := tcuo.mutation.LastTaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: taskconfig.FieldLastTaskID,
		})
	}
	if tcuo.mutation.LastTaskIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: taskconfig.FieldLastTaskID,
		})
	}
	if value, ok := tcuo.mutation.MaxRewardCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if value, ok := tcuo.mutation.AddedMaxRewardCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if tcuo.mutation.MaxRewardCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldMaxRewardCount,
		})
	}
	if value, ok := tcuo.mutation.CooldownSecord(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	if value, ok := tcuo.mutation.AddedCooldownSecord(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	if tcuo.mutation.CooldownSecordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: taskconfig.FieldCooldownSecord,
		})
	}
	_spec.Modifiers = tcuo.modifiers
	_node = &TaskConfig{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
