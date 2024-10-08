// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coinallocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinAllocatedUpdate is the builder for updating CoinAllocated entities.
type CoinAllocatedUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinAllocatedUpdate builder.
func (cau *CoinAllocatedUpdate) Where(ps ...predicate.CoinAllocated) *CoinAllocatedUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetCreatedAt sets the "created_at" field.
func (cau *CoinAllocatedUpdate) SetCreatedAt(u uint32) *CoinAllocatedUpdate {
	cau.mutation.ResetCreatedAt()
	cau.mutation.SetCreatedAt(u)
	return cau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableCreatedAt(u *uint32) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetCreatedAt(*u)
	}
	return cau
}

// AddCreatedAt adds u to the "created_at" field.
func (cau *CoinAllocatedUpdate) AddCreatedAt(u int32) *CoinAllocatedUpdate {
	cau.mutation.AddCreatedAt(u)
	return cau
}

// SetUpdatedAt sets the "updated_at" field.
func (cau *CoinAllocatedUpdate) SetUpdatedAt(u uint32) *CoinAllocatedUpdate {
	cau.mutation.ResetUpdatedAt()
	cau.mutation.SetUpdatedAt(u)
	return cau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cau *CoinAllocatedUpdate) AddUpdatedAt(u int32) *CoinAllocatedUpdate {
	cau.mutation.AddUpdatedAt(u)
	return cau
}

// SetDeletedAt sets the "deleted_at" field.
func (cau *CoinAllocatedUpdate) SetDeletedAt(u uint32) *CoinAllocatedUpdate {
	cau.mutation.ResetDeletedAt()
	cau.mutation.SetDeletedAt(u)
	return cau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableDeletedAt(u *uint32) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetDeletedAt(*u)
	}
	return cau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cau *CoinAllocatedUpdate) AddDeletedAt(u int32) *CoinAllocatedUpdate {
	cau.mutation.AddDeletedAt(u)
	return cau
}

// SetEntID sets the "ent_id" field.
func (cau *CoinAllocatedUpdate) SetEntID(u uuid.UUID) *CoinAllocatedUpdate {
	cau.mutation.SetEntID(u)
	return cau
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableEntID(u *uuid.UUID) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetEntID(*u)
	}
	return cau
}

// SetAppID sets the "app_id" field.
func (cau *CoinAllocatedUpdate) SetAppID(u uuid.UUID) *CoinAllocatedUpdate {
	cau.mutation.SetAppID(u)
	return cau
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableAppID(u *uuid.UUID) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetAppID(*u)
	}
	return cau
}

// ClearAppID clears the value of the "app_id" field.
func (cau *CoinAllocatedUpdate) ClearAppID() *CoinAllocatedUpdate {
	cau.mutation.ClearAppID()
	return cau
}

// SetCoinConfigID sets the "coin_config_id" field.
func (cau *CoinAllocatedUpdate) SetCoinConfigID(u uuid.UUID) *CoinAllocatedUpdate {
	cau.mutation.SetCoinConfigID(u)
	return cau
}

// SetNillableCoinConfigID sets the "coin_config_id" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableCoinConfigID(u *uuid.UUID) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetCoinConfigID(*u)
	}
	return cau
}

// ClearCoinConfigID clears the value of the "coin_config_id" field.
func (cau *CoinAllocatedUpdate) ClearCoinConfigID() *CoinAllocatedUpdate {
	cau.mutation.ClearCoinConfigID()
	return cau
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cau *CoinAllocatedUpdate) SetCoinTypeID(u uuid.UUID) *CoinAllocatedUpdate {
	cau.mutation.SetCoinTypeID(u)
	return cau
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetCoinTypeID(*u)
	}
	return cau
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cau *CoinAllocatedUpdate) ClearCoinTypeID() *CoinAllocatedUpdate {
	cau.mutation.ClearCoinTypeID()
	return cau
}

// SetUserID sets the "user_id" field.
func (cau *CoinAllocatedUpdate) SetUserID(u uuid.UUID) *CoinAllocatedUpdate {
	cau.mutation.SetUserID(u)
	return cau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableUserID(u *uuid.UUID) *CoinAllocatedUpdate {
	if u != nil {
		cau.SetUserID(*u)
	}
	return cau
}

// ClearUserID clears the value of the "user_id" field.
func (cau *CoinAllocatedUpdate) ClearUserID() *CoinAllocatedUpdate {
	cau.mutation.ClearUserID()
	return cau
}

// SetValue sets the "value" field.
func (cau *CoinAllocatedUpdate) SetValue(d decimal.Decimal) *CoinAllocatedUpdate {
	cau.mutation.SetValue(d)
	return cau
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableValue(d *decimal.Decimal) *CoinAllocatedUpdate {
	if d != nil {
		cau.SetValue(*d)
	}
	return cau
}

// ClearValue clears the value of the "value" field.
func (cau *CoinAllocatedUpdate) ClearValue() *CoinAllocatedUpdate {
	cau.mutation.ClearValue()
	return cau
}

// SetExtra sets the "extra" field.
func (cau *CoinAllocatedUpdate) SetExtra(s string) *CoinAllocatedUpdate {
	cau.mutation.SetExtra(s)
	return cau
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (cau *CoinAllocatedUpdate) SetNillableExtra(s *string) *CoinAllocatedUpdate {
	if s != nil {
		cau.SetExtra(*s)
	}
	return cau
}

// ClearExtra clears the value of the "extra" field.
func (cau *CoinAllocatedUpdate) ClearExtra() *CoinAllocatedUpdate {
	cau.mutation.ClearExtra()
	return cau
}

// Mutation returns the CoinAllocatedMutation object of the builder.
func (cau *CoinAllocatedUpdate) Mutation() *CoinAllocatedMutation {
	return cau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CoinAllocatedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cau.defaults(); err != nil {
		return 0, err
	}
	if len(cau.hooks) == 0 {
		if err = cau.check(); err != nil {
			return 0, err
		}
		affected, err = cau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cau.check(); err != nil {
				return 0, err
			}
			cau.mutation = mutation
			affected, err = cau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cau.hooks) - 1; i >= 0; i-- {
			if cau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cau *CoinAllocatedUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CoinAllocatedUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CoinAllocatedUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cau *CoinAllocatedUpdate) defaults() error {
	if _, ok := cau.mutation.UpdatedAt(); !ok {
		if coinallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinallocated.UpdateDefaultUpdatedAt()
		cau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cau *CoinAllocatedUpdate) check() error {
	if v, ok := cau.mutation.Extra(); ok {
		if err := coinallocated.ExtraValidator(v); err != nil {
			return &ValidationError{Name: "extra", err: fmt.Errorf(`ent: validator failed for field "CoinAllocated.extra": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cau *CoinAllocatedUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinAllocatedUpdate {
	cau.modifiers = append(cau.modifiers, modifiers...)
	return cau
}

func (cau *CoinAllocatedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinallocated.Table,
			Columns: coinallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinallocated.FieldID,
			},
		},
	}
	if ps := cau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldEntID,
		})
	}
	if value, ok := cau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldAppID,
		})
	}
	if cau.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldAppID,
		})
	}
	if value, ok := cau.mutation.CoinConfigID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinConfigID,
		})
	}
	if cau.mutation.CoinConfigIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldCoinConfigID,
		})
	}
	if value, ok := cau.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinTypeID,
		})
	}
	if cau.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldCoinTypeID,
		})
	}
	if value, ok := cau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldUserID,
		})
	}
	if cau.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldUserID,
		})
	}
	if value, ok := cau.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: coinallocated.FieldValue,
		})
	}
	if cau.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: coinallocated.FieldValue,
		})
	}
	if value, ok := cau.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinallocated.FieldExtra,
		})
	}
	if cau.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinallocated.FieldExtra,
		})
	}
	_spec.Modifiers = cau.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinAllocatedUpdateOne is the builder for updating a single CoinAllocated entity.
type CoinAllocatedUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cauo *CoinAllocatedUpdateOne) SetCreatedAt(u uint32) *CoinAllocatedUpdateOne {
	cauo.mutation.ResetCreatedAt()
	cauo.mutation.SetCreatedAt(u)
	return cauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableCreatedAt(u *uint32) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetCreatedAt(*u)
	}
	return cauo
}

// AddCreatedAt adds u to the "created_at" field.
func (cauo *CoinAllocatedUpdateOne) AddCreatedAt(u int32) *CoinAllocatedUpdateOne {
	cauo.mutation.AddCreatedAt(u)
	return cauo
}

// SetUpdatedAt sets the "updated_at" field.
func (cauo *CoinAllocatedUpdateOne) SetUpdatedAt(u uint32) *CoinAllocatedUpdateOne {
	cauo.mutation.ResetUpdatedAt()
	cauo.mutation.SetUpdatedAt(u)
	return cauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cauo *CoinAllocatedUpdateOne) AddUpdatedAt(u int32) *CoinAllocatedUpdateOne {
	cauo.mutation.AddUpdatedAt(u)
	return cauo
}

// SetDeletedAt sets the "deleted_at" field.
func (cauo *CoinAllocatedUpdateOne) SetDeletedAt(u uint32) *CoinAllocatedUpdateOne {
	cauo.mutation.ResetDeletedAt()
	cauo.mutation.SetDeletedAt(u)
	return cauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableDeletedAt(u *uint32) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetDeletedAt(*u)
	}
	return cauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cauo *CoinAllocatedUpdateOne) AddDeletedAt(u int32) *CoinAllocatedUpdateOne {
	cauo.mutation.AddDeletedAt(u)
	return cauo
}

// SetEntID sets the "ent_id" field.
func (cauo *CoinAllocatedUpdateOne) SetEntID(u uuid.UUID) *CoinAllocatedUpdateOne {
	cauo.mutation.SetEntID(u)
	return cauo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetEntID(*u)
	}
	return cauo
}

// SetAppID sets the "app_id" field.
func (cauo *CoinAllocatedUpdateOne) SetAppID(u uuid.UUID) *CoinAllocatedUpdateOne {
	cauo.mutation.SetAppID(u)
	return cauo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableAppID(u *uuid.UUID) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetAppID(*u)
	}
	return cauo
}

// ClearAppID clears the value of the "app_id" field.
func (cauo *CoinAllocatedUpdateOne) ClearAppID() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearAppID()
	return cauo
}

// SetCoinConfigID sets the "coin_config_id" field.
func (cauo *CoinAllocatedUpdateOne) SetCoinConfigID(u uuid.UUID) *CoinAllocatedUpdateOne {
	cauo.mutation.SetCoinConfigID(u)
	return cauo
}

// SetNillableCoinConfigID sets the "coin_config_id" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableCoinConfigID(u *uuid.UUID) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetCoinConfigID(*u)
	}
	return cauo
}

// ClearCoinConfigID clears the value of the "coin_config_id" field.
func (cauo *CoinAllocatedUpdateOne) ClearCoinConfigID() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearCoinConfigID()
	return cauo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cauo *CoinAllocatedUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinAllocatedUpdateOne {
	cauo.mutation.SetCoinTypeID(u)
	return cauo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetCoinTypeID(*u)
	}
	return cauo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cauo *CoinAllocatedUpdateOne) ClearCoinTypeID() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearCoinTypeID()
	return cauo
}

// SetUserID sets the "user_id" field.
func (cauo *CoinAllocatedUpdateOne) SetUserID(u uuid.UUID) *CoinAllocatedUpdateOne {
	cauo.mutation.SetUserID(u)
	return cauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableUserID(u *uuid.UUID) *CoinAllocatedUpdateOne {
	if u != nil {
		cauo.SetUserID(*u)
	}
	return cauo
}

// ClearUserID clears the value of the "user_id" field.
func (cauo *CoinAllocatedUpdateOne) ClearUserID() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearUserID()
	return cauo
}

// SetValue sets the "value" field.
func (cauo *CoinAllocatedUpdateOne) SetValue(d decimal.Decimal) *CoinAllocatedUpdateOne {
	cauo.mutation.SetValue(d)
	return cauo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableValue(d *decimal.Decimal) *CoinAllocatedUpdateOne {
	if d != nil {
		cauo.SetValue(*d)
	}
	return cauo
}

// ClearValue clears the value of the "value" field.
func (cauo *CoinAllocatedUpdateOne) ClearValue() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearValue()
	return cauo
}

// SetExtra sets the "extra" field.
func (cauo *CoinAllocatedUpdateOne) SetExtra(s string) *CoinAllocatedUpdateOne {
	cauo.mutation.SetExtra(s)
	return cauo
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (cauo *CoinAllocatedUpdateOne) SetNillableExtra(s *string) *CoinAllocatedUpdateOne {
	if s != nil {
		cauo.SetExtra(*s)
	}
	return cauo
}

// ClearExtra clears the value of the "extra" field.
func (cauo *CoinAllocatedUpdateOne) ClearExtra() *CoinAllocatedUpdateOne {
	cauo.mutation.ClearExtra()
	return cauo
}

// Mutation returns the CoinAllocatedMutation object of the builder.
func (cauo *CoinAllocatedUpdateOne) Mutation() *CoinAllocatedMutation {
	return cauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CoinAllocatedUpdateOne) Select(field string, fields ...string) *CoinAllocatedUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CoinAllocated entity.
func (cauo *CoinAllocatedUpdateOne) Save(ctx context.Context) (*CoinAllocated, error) {
	var (
		err  error
		node *CoinAllocated
	)
	if err := cauo.defaults(); err != nil {
		return nil, err
	}
	if len(cauo.hooks) == 0 {
		if err = cauo.check(); err != nil {
			return nil, err
		}
		node, err = cauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cauo.check(); err != nil {
				return nil, err
			}
			cauo.mutation = mutation
			node, err = cauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cauo.hooks) - 1; i >= 0; i-- {
			if cauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinAllocated)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinAllocatedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CoinAllocatedUpdateOne) SaveX(ctx context.Context) *CoinAllocated {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CoinAllocatedUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CoinAllocatedUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cauo *CoinAllocatedUpdateOne) defaults() error {
	if _, ok := cauo.mutation.UpdatedAt(); !ok {
		if coinallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinallocated.UpdateDefaultUpdatedAt()
		cauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cauo *CoinAllocatedUpdateOne) check() error {
	if v, ok := cauo.mutation.Extra(); ok {
		if err := coinallocated.ExtraValidator(v); err != nil {
			return &ValidationError{Name: "extra", err: fmt.Errorf(`ent: validator failed for field "CoinAllocated.extra": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cauo *CoinAllocatedUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinAllocatedUpdateOne {
	cauo.modifiers = append(cauo.modifiers, modifiers...)
	return cauo
}

func (cauo *CoinAllocatedUpdateOne) sqlSave(ctx context.Context) (_node *CoinAllocated, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinallocated.Table,
			Columns: coinallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinallocated.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinAllocated.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinallocated.FieldID)
		for _, f := range fields {
			if !coinallocated.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinallocated.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldEntID,
		})
	}
	if value, ok := cauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldAppID,
		})
	}
	if cauo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldAppID,
		})
	}
	if value, ok := cauo.mutation.CoinConfigID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinConfigID,
		})
	}
	if cauo.mutation.CoinConfigIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldCoinConfigID,
		})
	}
	if value, ok := cauo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinTypeID,
		})
	}
	if cauo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldCoinTypeID,
		})
	}
	if value, ok := cauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldUserID,
		})
	}
	if cauo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinallocated.FieldUserID,
		})
	}
	if value, ok := cauo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: coinallocated.FieldValue,
		})
	}
	if cauo.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: coinallocated.FieldValue,
		})
	}
	if value, ok := cauo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinallocated.FieldExtra,
		})
	}
	if cauo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinallocated.FieldExtra,
		})
	}
	_spec.Modifiers = cauo.modifiers
	_node = &CoinAllocated{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
