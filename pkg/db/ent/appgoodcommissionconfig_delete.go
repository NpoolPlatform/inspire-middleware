// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodcommissionconfig"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// AppGoodCommissionConfigDelete is the builder for deleting a AppGoodCommissionConfig entity.
type AppGoodCommissionConfigDelete struct {
	config
	hooks    []Hook
	mutation *AppGoodCommissionConfigMutation
}

// Where appends a list predicates to the AppGoodCommissionConfigDelete builder.
func (agccd *AppGoodCommissionConfigDelete) Where(ps ...predicate.AppGoodCommissionConfig) *AppGoodCommissionConfigDelete {
	agccd.mutation.Where(ps...)
	return agccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agccd *AppGoodCommissionConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agccd.hooks) == 0 {
		affected, err = agccd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodCommissionConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agccd.mutation = mutation
			affected, err = agccd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agccd.hooks) - 1; i >= 0; i-- {
			if agccd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agccd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agccd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agccd *AppGoodCommissionConfigDelete) ExecX(ctx context.Context) int {
	n, err := agccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agccd *AppGoodCommissionConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appgoodcommissionconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgoodcommissionconfig.FieldID,
			},
		},
	}
	if ps := agccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, agccd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppGoodCommissionConfigDeleteOne is the builder for deleting a single AppGoodCommissionConfig entity.
type AppGoodCommissionConfigDeleteOne struct {
	agccd *AppGoodCommissionConfigDelete
}

// Exec executes the deletion query.
func (agccdo *AppGoodCommissionConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := agccdo.agccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appgoodcommissionconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agccdo *AppGoodCommissionConfigDeleteOne) ExecX(ctx context.Context) {
	agccdo.agccd.ExecX(ctx)
}
