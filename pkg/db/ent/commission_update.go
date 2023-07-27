// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CommissionUpdate is the builder for updating Commission entities.
type CommissionUpdate struct {
	config
	hooks     []Hook
	mutation  *CommissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CommissionUpdate builder.
func (cu *CommissionUpdate) Where(ps ...predicate.Commission) *CommissionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommissionUpdate) SetCreatedAt(u uint32) *CommissionUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableCreatedAt(u *uint32) *CommissionUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CommissionUpdate) AddCreatedAt(u int32) *CommissionUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommissionUpdate) SetUpdatedAt(u uint32) *CommissionUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CommissionUpdate) AddUpdatedAt(u int32) *CommissionUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CommissionUpdate) SetDeletedAt(u uint32) *CommissionUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableDeletedAt(u *uint32) *CommissionUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CommissionUpdate) AddDeletedAt(u int32) *CommissionUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetAppID sets the "app_id" field.
func (cu *CommissionUpdate) SetAppID(u uuid.UUID) *CommissionUpdate {
	cu.mutation.SetAppID(u)
	return cu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableAppID(u *uuid.UUID) *CommissionUpdate {
	if u != nil {
		cu.SetAppID(*u)
	}
	return cu
}

// ClearAppID clears the value of the "app_id" field.
func (cu *CommissionUpdate) ClearAppID() *CommissionUpdate {
	cu.mutation.ClearAppID()
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommissionUpdate) SetUserID(u uuid.UUID) *CommissionUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableUserID(u *uuid.UUID) *CommissionUpdate {
	if u != nil {
		cu.SetUserID(*u)
	}
	return cu
}

// ClearUserID clears the value of the "user_id" field.
func (cu *CommissionUpdate) ClearUserID() *CommissionUpdate {
	cu.mutation.ClearUserID()
	return cu
}

// SetGoodID sets the "good_id" field.
func (cu *CommissionUpdate) SetGoodID(u uuid.UUID) *CommissionUpdate {
	cu.mutation.SetGoodID(u)
	return cu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableGoodID(u *uuid.UUID) *CommissionUpdate {
	if u != nil {
		cu.SetGoodID(*u)
	}
	return cu
}

// ClearGoodID clears the value of the "good_id" field.
func (cu *CommissionUpdate) ClearGoodID() *CommissionUpdate {
	cu.mutation.ClearGoodID()
	return cu
}

// SetPercent sets the "percent" field.
func (cu *CommissionUpdate) SetPercent(d decimal.Decimal) *CommissionUpdate {
	cu.mutation.SetPercent(d)
	return cu
}

// SetNillablePercent sets the "percent" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillablePercent(d *decimal.Decimal) *CommissionUpdate {
	if d != nil {
		cu.SetPercent(*d)
	}
	return cu
}

// ClearPercent clears the value of the "percent" field.
func (cu *CommissionUpdate) ClearPercent() *CommissionUpdate {
	cu.mutation.ClearPercent()
	return cu
}

// SetStartAt sets the "start_at" field.
func (cu *CommissionUpdate) SetStartAt(u uint32) *CommissionUpdate {
	cu.mutation.ResetStartAt()
	cu.mutation.SetStartAt(u)
	return cu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableStartAt(u *uint32) *CommissionUpdate {
	if u != nil {
		cu.SetStartAt(*u)
	}
	return cu
}

// AddStartAt adds u to the "start_at" field.
func (cu *CommissionUpdate) AddStartAt(u int32) *CommissionUpdate {
	cu.mutation.AddStartAt(u)
	return cu
}

// ClearStartAt clears the value of the "start_at" field.
func (cu *CommissionUpdate) ClearStartAt() *CommissionUpdate {
	cu.mutation.ClearStartAt()
	return cu
}

// SetEndAt sets the "end_at" field.
func (cu *CommissionUpdate) SetEndAt(u uint32) *CommissionUpdate {
	cu.mutation.ResetEndAt()
	cu.mutation.SetEndAt(u)
	return cu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableEndAt(u *uint32) *CommissionUpdate {
	if u != nil {
		cu.SetEndAt(*u)
	}
	return cu
}

// AddEndAt adds u to the "end_at" field.
func (cu *CommissionUpdate) AddEndAt(u int32) *CommissionUpdate {
	cu.mutation.AddEndAt(u)
	return cu
}

// ClearEndAt clears the value of the "end_at" field.
func (cu *CommissionUpdate) ClearEndAt() *CommissionUpdate {
	cu.mutation.ClearEndAt()
	return cu
}

// SetSettleType sets the "settle_type" field.
func (cu *CommissionUpdate) SetSettleType(s string) *CommissionUpdate {
	cu.mutation.SetSettleType(s)
	return cu
}

// SetNillableSettleType sets the "settle_type" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableSettleType(s *string) *CommissionUpdate {
	if s != nil {
		cu.SetSettleType(*s)
	}
	return cu
}

// ClearSettleType clears the value of the "settle_type" field.
func (cu *CommissionUpdate) ClearSettleType() *CommissionUpdate {
	cu.mutation.ClearSettleType()
	return cu
}

// SetSettleMode sets the "settle_mode" field.
func (cu *CommissionUpdate) SetSettleMode(s string) *CommissionUpdate {
	cu.mutation.SetSettleMode(s)
	return cu
}

// SetNillableSettleMode sets the "settle_mode" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableSettleMode(s *string) *CommissionUpdate {
	if s != nil {
		cu.SetSettleMode(*s)
	}
	return cu
}

// ClearSettleMode clears the value of the "settle_mode" field.
func (cu *CommissionUpdate) ClearSettleMode() *CommissionUpdate {
	cu.mutation.ClearSettleMode()
	return cu
}

// SetSettleInterval sets the "settle_interval" field.
func (cu *CommissionUpdate) SetSettleInterval(s string) *CommissionUpdate {
	cu.mutation.SetSettleInterval(s)
	return cu
}

// SetNillableSettleInterval sets the "settle_interval" field if the given value is not nil.
func (cu *CommissionUpdate) SetNillableSettleInterval(s *string) *CommissionUpdate {
	if s != nil {
		cu.SetSettleInterval(*s)
	}
	return cu
}

// ClearSettleInterval clears the value of the "settle_interval" field.
func (cu *CommissionUpdate) ClearSettleInterval() *CommissionUpdate {
	cu.mutation.ClearSettleInterval()
	return cu
}

// Mutation returns the CommissionMutation object of the builder.
func (cu *CommissionUpdate) Mutation() *CommissionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommissionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommissionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommissionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommissionUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if commission.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized commission.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := commission.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CommissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommissionUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CommissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commission.Table,
			Columns: commission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commission.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldAppID,
		})
	}
	if cu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldAppID,
		})
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldUserID,
		})
	}
	if cu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldUserID,
		})
	}
	if value, ok := cu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldGoodID,
		})
	}
	if cu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldGoodID,
		})
	}
	if value, ok := cu.mutation.Percent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: commission.FieldPercent,
		})
	}
	if cu.mutation.PercentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: commission.FieldPercent,
		})
	}
	if value, ok := cu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldStartAt,
		})
	}
	if value, ok := cu.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldStartAt,
		})
	}
	if cu.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: commission.FieldStartAt,
		})
	}
	if value, ok := cu.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldEndAt,
		})
	}
	if value, ok := cu.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldEndAt,
		})
	}
	if cu.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: commission.FieldEndAt,
		})
	}
	if value, ok := cu.mutation.SettleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleType,
		})
	}
	if cu.mutation.SettleTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleType,
		})
	}
	if value, ok := cu.mutation.SettleMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleMode,
		})
	}
	if cu.mutation.SettleModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleMode,
		})
	}
	if value, ok := cu.mutation.SettleInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleInterval,
		})
	}
	if cu.mutation.SettleIntervalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleInterval,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CommissionUpdateOne is the builder for updating a single Commission entity.
type CommissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CommissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommissionUpdateOne) SetCreatedAt(u uint32) *CommissionUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableCreatedAt(u *uint32) *CommissionUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CommissionUpdateOne) AddCreatedAt(u int32) *CommissionUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommissionUpdateOne) SetUpdatedAt(u uint32) *CommissionUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CommissionUpdateOne) AddUpdatedAt(u int32) *CommissionUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CommissionUpdateOne) SetDeletedAt(u uint32) *CommissionUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableDeletedAt(u *uint32) *CommissionUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CommissionUpdateOne) AddDeletedAt(u int32) *CommissionUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetAppID sets the "app_id" field.
func (cuo *CommissionUpdateOne) SetAppID(u uuid.UUID) *CommissionUpdateOne {
	cuo.mutation.SetAppID(u)
	return cuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableAppID(u *uuid.UUID) *CommissionUpdateOne {
	if u != nil {
		cuo.SetAppID(*u)
	}
	return cuo
}

// ClearAppID clears the value of the "app_id" field.
func (cuo *CommissionUpdateOne) ClearAppID() *CommissionUpdateOne {
	cuo.mutation.ClearAppID()
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CommissionUpdateOne) SetUserID(u uuid.UUID) *CommissionUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableUserID(u *uuid.UUID) *CommissionUpdateOne {
	if u != nil {
		cuo.SetUserID(*u)
	}
	return cuo
}

// ClearUserID clears the value of the "user_id" field.
func (cuo *CommissionUpdateOne) ClearUserID() *CommissionUpdateOne {
	cuo.mutation.ClearUserID()
	return cuo
}

// SetGoodID sets the "good_id" field.
func (cuo *CommissionUpdateOne) SetGoodID(u uuid.UUID) *CommissionUpdateOne {
	cuo.mutation.SetGoodID(u)
	return cuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableGoodID(u *uuid.UUID) *CommissionUpdateOne {
	if u != nil {
		cuo.SetGoodID(*u)
	}
	return cuo
}

// ClearGoodID clears the value of the "good_id" field.
func (cuo *CommissionUpdateOne) ClearGoodID() *CommissionUpdateOne {
	cuo.mutation.ClearGoodID()
	return cuo
}

// SetPercent sets the "percent" field.
func (cuo *CommissionUpdateOne) SetPercent(d decimal.Decimal) *CommissionUpdateOne {
	cuo.mutation.SetPercent(d)
	return cuo
}

// SetNillablePercent sets the "percent" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillablePercent(d *decimal.Decimal) *CommissionUpdateOne {
	if d != nil {
		cuo.SetPercent(*d)
	}
	return cuo
}

// ClearPercent clears the value of the "percent" field.
func (cuo *CommissionUpdateOne) ClearPercent() *CommissionUpdateOne {
	cuo.mutation.ClearPercent()
	return cuo
}

// SetStartAt sets the "start_at" field.
func (cuo *CommissionUpdateOne) SetStartAt(u uint32) *CommissionUpdateOne {
	cuo.mutation.ResetStartAt()
	cuo.mutation.SetStartAt(u)
	return cuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableStartAt(u *uint32) *CommissionUpdateOne {
	if u != nil {
		cuo.SetStartAt(*u)
	}
	return cuo
}

// AddStartAt adds u to the "start_at" field.
func (cuo *CommissionUpdateOne) AddStartAt(u int32) *CommissionUpdateOne {
	cuo.mutation.AddStartAt(u)
	return cuo
}

// ClearStartAt clears the value of the "start_at" field.
func (cuo *CommissionUpdateOne) ClearStartAt() *CommissionUpdateOne {
	cuo.mutation.ClearStartAt()
	return cuo
}

// SetEndAt sets the "end_at" field.
func (cuo *CommissionUpdateOne) SetEndAt(u uint32) *CommissionUpdateOne {
	cuo.mutation.ResetEndAt()
	cuo.mutation.SetEndAt(u)
	return cuo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableEndAt(u *uint32) *CommissionUpdateOne {
	if u != nil {
		cuo.SetEndAt(*u)
	}
	return cuo
}

// AddEndAt adds u to the "end_at" field.
func (cuo *CommissionUpdateOne) AddEndAt(u int32) *CommissionUpdateOne {
	cuo.mutation.AddEndAt(u)
	return cuo
}

// ClearEndAt clears the value of the "end_at" field.
func (cuo *CommissionUpdateOne) ClearEndAt() *CommissionUpdateOne {
	cuo.mutation.ClearEndAt()
	return cuo
}

// SetSettleType sets the "settle_type" field.
func (cuo *CommissionUpdateOne) SetSettleType(s string) *CommissionUpdateOne {
	cuo.mutation.SetSettleType(s)
	return cuo
}

// SetNillableSettleType sets the "settle_type" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableSettleType(s *string) *CommissionUpdateOne {
	if s != nil {
		cuo.SetSettleType(*s)
	}
	return cuo
}

// ClearSettleType clears the value of the "settle_type" field.
func (cuo *CommissionUpdateOne) ClearSettleType() *CommissionUpdateOne {
	cuo.mutation.ClearSettleType()
	return cuo
}

// SetSettleMode sets the "settle_mode" field.
func (cuo *CommissionUpdateOne) SetSettleMode(s string) *CommissionUpdateOne {
	cuo.mutation.SetSettleMode(s)
	return cuo
}

// SetNillableSettleMode sets the "settle_mode" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableSettleMode(s *string) *CommissionUpdateOne {
	if s != nil {
		cuo.SetSettleMode(*s)
	}
	return cuo
}

// ClearSettleMode clears the value of the "settle_mode" field.
func (cuo *CommissionUpdateOne) ClearSettleMode() *CommissionUpdateOne {
	cuo.mutation.ClearSettleMode()
	return cuo
}

// SetSettleInterval sets the "settle_interval" field.
func (cuo *CommissionUpdateOne) SetSettleInterval(s string) *CommissionUpdateOne {
	cuo.mutation.SetSettleInterval(s)
	return cuo
}

// SetNillableSettleInterval sets the "settle_interval" field if the given value is not nil.
func (cuo *CommissionUpdateOne) SetNillableSettleInterval(s *string) *CommissionUpdateOne {
	if s != nil {
		cuo.SetSettleInterval(*s)
	}
	return cuo
}

// ClearSettleInterval clears the value of the "settle_interval" field.
func (cuo *CommissionUpdateOne) ClearSettleInterval() *CommissionUpdateOne {
	cuo.mutation.ClearSettleInterval()
	return cuo
}

// Mutation returns the CommissionMutation object of the builder.
func (cuo *CommissionUpdateOne) Mutation() *CommissionMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommissionUpdateOne) Select(field string, fields ...string) *CommissionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Commission entity.
func (cuo *CommissionUpdateOne) Save(ctx context.Context) (*Commission, error) {
	var (
		err  error
		node *Commission
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Commission)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CommissionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommissionUpdateOne) SaveX(ctx context.Context) *Commission {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommissionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommissionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommissionUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if commission.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized commission.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := commission.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CommissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommissionUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CommissionUpdateOne) sqlSave(ctx context.Context) (_node *Commission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commission.Table,
			Columns: commission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commission.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Commission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commission.FieldID)
		for _, f := range fields {
			if !commission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldAppID,
		})
	}
	if cuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldAppID,
		})
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldUserID,
		})
	}
	if cuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commission.FieldGoodID,
		})
	}
	if cuo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: commission.FieldGoodID,
		})
	}
	if value, ok := cuo.mutation.Percent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: commission.FieldPercent,
		})
	}
	if cuo.mutation.PercentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: commission.FieldPercent,
		})
	}
	if value, ok := cuo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldStartAt,
		})
	}
	if value, ok := cuo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldStartAt,
		})
	}
	if cuo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: commission.FieldStartAt,
		})
	}
	if value, ok := cuo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldEndAt,
		})
	}
	if value, ok := cuo.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commission.FieldEndAt,
		})
	}
	if cuo.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: commission.FieldEndAt,
		})
	}
	if value, ok := cuo.mutation.SettleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleType,
		})
	}
	if cuo.mutation.SettleTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleType,
		})
	}
	if value, ok := cuo.mutation.SettleMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleMode,
		})
	}
	if cuo.mutation.SettleModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleMode,
		})
	}
	if value, ok := cuo.mutation.SettleInterval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commission.FieldSettleInterval,
		})
	}
	if cuo.mutation.SettleIntervalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: commission.FieldSettleInterval,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Commission{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
