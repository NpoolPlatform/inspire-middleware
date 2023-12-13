// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// CouponScopeDelete is the builder for deleting a CouponScope entity.
type CouponScopeDelete struct {
	config
	hooks    []Hook
	mutation *CouponScopeMutation
}

// Where appends a list predicates to the CouponScopeDelete builder.
func (csd *CouponScopeDelete) Where(ps ...predicate.CouponScope) *CouponScopeDelete {
	csd.mutation.Where(ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *CouponScopeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csd.hooks) == 0 {
		affected, err = csd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponScopeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csd.mutation = mutation
			affected, err = csd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csd.hooks) - 1; i >= 0; i-- {
			if csd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *CouponScopeDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *CouponScopeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: couponscope.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: couponscope.FieldID,
			},
		},
	}
	if ps := csd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CouponScopeDeleteOne is the builder for deleting a single CouponScope entity.
type CouponScopeDeleteOne struct {
	csd *CouponScopeDelete
}

// Exec executes the deletion query.
func (csdo *CouponScopeDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{couponscope.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *CouponScopeDeleteOne) ExecX(ctx context.Context) {
	csdo.csd.ExecX(ctx)
}
