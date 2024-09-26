// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/taskconfig"
)

// TaskConfigDelete is the builder for deleting a TaskConfig entity.
type TaskConfigDelete struct {
	config
	hooks    []Hook
	mutation *TaskConfigMutation
}

// Where appends a list predicates to the TaskConfigDelete builder.
func (tcd *TaskConfigDelete) Where(ps ...predicate.TaskConfig) *TaskConfigDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TaskConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tcd.hooks) == 0 {
		affected, err = tcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcd.mutation = mutation
			affected, err = tcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcd.hooks) - 1; i >= 0; i-- {
			if tcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TaskConfigDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TaskConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: taskconfig.FieldID,
			},
		},
	}
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TaskConfigDeleteOne is the builder for deleting a single TaskConfig entity.
type TaskConfigDeleteOne struct {
	tcd *TaskConfigDelete
}

// Exec executes the deletion query.
func (tcdo *TaskConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TaskConfigDeleteOne) ExecX(ctx context.Context) {
	tcdo.tcd.ExecX(ctx)
}
