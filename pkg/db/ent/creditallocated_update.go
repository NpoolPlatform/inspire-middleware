// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/creditallocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CreditAllocatedUpdate is the builder for updating CreditAllocated entities.
type CreditAllocatedUpdate struct {
	config
	hooks     []Hook
	mutation  *CreditAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CreditAllocatedUpdate builder.
func (cau *CreditAllocatedUpdate) Where(ps ...predicate.CreditAllocated) *CreditAllocatedUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetCreatedAt sets the "created_at" field.
func (cau *CreditAllocatedUpdate) SetCreatedAt(u uint32) *CreditAllocatedUpdate {
	cau.mutation.ResetCreatedAt()
	cau.mutation.SetCreatedAt(u)
	return cau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableCreatedAt(u *uint32) *CreditAllocatedUpdate {
	if u != nil {
		cau.SetCreatedAt(*u)
	}
	return cau
}

// AddCreatedAt adds u to the "created_at" field.
func (cau *CreditAllocatedUpdate) AddCreatedAt(u int32) *CreditAllocatedUpdate {
	cau.mutation.AddCreatedAt(u)
	return cau
}

// SetUpdatedAt sets the "updated_at" field.
func (cau *CreditAllocatedUpdate) SetUpdatedAt(u uint32) *CreditAllocatedUpdate {
	cau.mutation.ResetUpdatedAt()
	cau.mutation.SetUpdatedAt(u)
	return cau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cau *CreditAllocatedUpdate) AddUpdatedAt(u int32) *CreditAllocatedUpdate {
	cau.mutation.AddUpdatedAt(u)
	return cau
}

// SetDeletedAt sets the "deleted_at" field.
func (cau *CreditAllocatedUpdate) SetDeletedAt(u uint32) *CreditAllocatedUpdate {
	cau.mutation.ResetDeletedAt()
	cau.mutation.SetDeletedAt(u)
	return cau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableDeletedAt(u *uint32) *CreditAllocatedUpdate {
	if u != nil {
		cau.SetDeletedAt(*u)
	}
	return cau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cau *CreditAllocatedUpdate) AddDeletedAt(u int32) *CreditAllocatedUpdate {
	cau.mutation.AddDeletedAt(u)
	return cau
}

// SetEntID sets the "ent_id" field.
func (cau *CreditAllocatedUpdate) SetEntID(u uuid.UUID) *CreditAllocatedUpdate {
	cau.mutation.SetEntID(u)
	return cau
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableEntID(u *uuid.UUID) *CreditAllocatedUpdate {
	if u != nil {
		cau.SetEntID(*u)
	}
	return cau
}

// SetAppID sets the "app_id" field.
func (cau *CreditAllocatedUpdate) SetAppID(u uuid.UUID) *CreditAllocatedUpdate {
	cau.mutation.SetAppID(u)
	return cau
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableAppID(u *uuid.UUID) *CreditAllocatedUpdate {
	if u != nil {
		cau.SetAppID(*u)
	}
	return cau
}

// ClearAppID clears the value of the "app_id" field.
func (cau *CreditAllocatedUpdate) ClearAppID() *CreditAllocatedUpdate {
	cau.mutation.ClearAppID()
	return cau
}

// SetUserID sets the "user_id" field.
func (cau *CreditAllocatedUpdate) SetUserID(u uuid.UUID) *CreditAllocatedUpdate {
	cau.mutation.SetUserID(u)
	return cau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableUserID(u *uuid.UUID) *CreditAllocatedUpdate {
	if u != nil {
		cau.SetUserID(*u)
	}
	return cau
}

// ClearUserID clears the value of the "user_id" field.
func (cau *CreditAllocatedUpdate) ClearUserID() *CreditAllocatedUpdate {
	cau.mutation.ClearUserID()
	return cau
}

// SetValue sets the "value" field.
func (cau *CreditAllocatedUpdate) SetValue(d decimal.Decimal) *CreditAllocatedUpdate {
	cau.mutation.SetValue(d)
	return cau
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableValue(d *decimal.Decimal) *CreditAllocatedUpdate {
	if d != nil {
		cau.SetValue(*d)
	}
	return cau
}

// ClearValue clears the value of the "value" field.
func (cau *CreditAllocatedUpdate) ClearValue() *CreditAllocatedUpdate {
	cau.mutation.ClearValue()
	return cau
}

// SetExtra sets the "extra" field.
func (cau *CreditAllocatedUpdate) SetExtra(s string) *CreditAllocatedUpdate {
	cau.mutation.SetExtra(s)
	return cau
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (cau *CreditAllocatedUpdate) SetNillableExtra(s *string) *CreditAllocatedUpdate {
	if s != nil {
		cau.SetExtra(*s)
	}
	return cau
}

// ClearExtra clears the value of the "extra" field.
func (cau *CreditAllocatedUpdate) ClearExtra() *CreditAllocatedUpdate {
	cau.mutation.ClearExtra()
	return cau
}

// Mutation returns the CreditAllocatedMutation object of the builder.
func (cau *CreditAllocatedUpdate) Mutation() *CreditAllocatedMutation {
	return cau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CreditAllocatedUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*CreditAllocatedMutation)
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
func (cau *CreditAllocatedUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CreditAllocatedUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CreditAllocatedUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cau *CreditAllocatedUpdate) defaults() error {
	if _, ok := cau.mutation.UpdatedAt(); !ok {
		if creditallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized creditallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := creditallocated.UpdateDefaultUpdatedAt()
		cau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cau *CreditAllocatedUpdate) check() error {
	if v, ok := cau.mutation.Extra(); ok {
		if err := creditallocated.ExtraValidator(v); err != nil {
			return &ValidationError{Name: "extra", err: fmt.Errorf(`ent: validator failed for field "CreditAllocated.extra": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cau *CreditAllocatedUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CreditAllocatedUpdate {
	cau.modifiers = append(cau.modifiers, modifiers...)
	return cau
}

func (cau *CreditAllocatedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   creditallocated.Table,
			Columns: creditallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: creditallocated.FieldID,
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
			Column: creditallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldEntID,
		})
	}
	if value, ok := cau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldAppID,
		})
	}
	if cau.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: creditallocated.FieldAppID,
		})
	}
	if value, ok := cau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldUserID,
		})
	}
	if cau.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: creditallocated.FieldUserID,
		})
	}
	if value, ok := cau.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: creditallocated.FieldValue,
		})
	}
	if cau.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: creditallocated.FieldValue,
		})
	}
	if value, ok := cau.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: creditallocated.FieldExtra,
		})
	}
	if cau.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: creditallocated.FieldExtra,
		})
	}
	_spec.Modifiers = cau.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creditallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CreditAllocatedUpdateOne is the builder for updating a single CreditAllocated entity.
type CreditAllocatedUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CreditAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cauo *CreditAllocatedUpdateOne) SetCreatedAt(u uint32) *CreditAllocatedUpdateOne {
	cauo.mutation.ResetCreatedAt()
	cauo.mutation.SetCreatedAt(u)
	return cauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableCreatedAt(u *uint32) *CreditAllocatedUpdateOne {
	if u != nil {
		cauo.SetCreatedAt(*u)
	}
	return cauo
}

// AddCreatedAt adds u to the "created_at" field.
func (cauo *CreditAllocatedUpdateOne) AddCreatedAt(u int32) *CreditAllocatedUpdateOne {
	cauo.mutation.AddCreatedAt(u)
	return cauo
}

// SetUpdatedAt sets the "updated_at" field.
func (cauo *CreditAllocatedUpdateOne) SetUpdatedAt(u uint32) *CreditAllocatedUpdateOne {
	cauo.mutation.ResetUpdatedAt()
	cauo.mutation.SetUpdatedAt(u)
	return cauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cauo *CreditAllocatedUpdateOne) AddUpdatedAt(u int32) *CreditAllocatedUpdateOne {
	cauo.mutation.AddUpdatedAt(u)
	return cauo
}

// SetDeletedAt sets the "deleted_at" field.
func (cauo *CreditAllocatedUpdateOne) SetDeletedAt(u uint32) *CreditAllocatedUpdateOne {
	cauo.mutation.ResetDeletedAt()
	cauo.mutation.SetDeletedAt(u)
	return cauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableDeletedAt(u *uint32) *CreditAllocatedUpdateOne {
	if u != nil {
		cauo.SetDeletedAt(*u)
	}
	return cauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cauo *CreditAllocatedUpdateOne) AddDeletedAt(u int32) *CreditAllocatedUpdateOne {
	cauo.mutation.AddDeletedAt(u)
	return cauo
}

// SetEntID sets the "ent_id" field.
func (cauo *CreditAllocatedUpdateOne) SetEntID(u uuid.UUID) *CreditAllocatedUpdateOne {
	cauo.mutation.SetEntID(u)
	return cauo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableEntID(u *uuid.UUID) *CreditAllocatedUpdateOne {
	if u != nil {
		cauo.SetEntID(*u)
	}
	return cauo
}

// SetAppID sets the "app_id" field.
func (cauo *CreditAllocatedUpdateOne) SetAppID(u uuid.UUID) *CreditAllocatedUpdateOne {
	cauo.mutation.SetAppID(u)
	return cauo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableAppID(u *uuid.UUID) *CreditAllocatedUpdateOne {
	if u != nil {
		cauo.SetAppID(*u)
	}
	return cauo
}

// ClearAppID clears the value of the "app_id" field.
func (cauo *CreditAllocatedUpdateOne) ClearAppID() *CreditAllocatedUpdateOne {
	cauo.mutation.ClearAppID()
	return cauo
}

// SetUserID sets the "user_id" field.
func (cauo *CreditAllocatedUpdateOne) SetUserID(u uuid.UUID) *CreditAllocatedUpdateOne {
	cauo.mutation.SetUserID(u)
	return cauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableUserID(u *uuid.UUID) *CreditAllocatedUpdateOne {
	if u != nil {
		cauo.SetUserID(*u)
	}
	return cauo
}

// ClearUserID clears the value of the "user_id" field.
func (cauo *CreditAllocatedUpdateOne) ClearUserID() *CreditAllocatedUpdateOne {
	cauo.mutation.ClearUserID()
	return cauo
}

// SetValue sets the "value" field.
func (cauo *CreditAllocatedUpdateOne) SetValue(d decimal.Decimal) *CreditAllocatedUpdateOne {
	cauo.mutation.SetValue(d)
	return cauo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableValue(d *decimal.Decimal) *CreditAllocatedUpdateOne {
	if d != nil {
		cauo.SetValue(*d)
	}
	return cauo
}

// ClearValue clears the value of the "value" field.
func (cauo *CreditAllocatedUpdateOne) ClearValue() *CreditAllocatedUpdateOne {
	cauo.mutation.ClearValue()
	return cauo
}

// SetExtra sets the "extra" field.
func (cauo *CreditAllocatedUpdateOne) SetExtra(s string) *CreditAllocatedUpdateOne {
	cauo.mutation.SetExtra(s)
	return cauo
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (cauo *CreditAllocatedUpdateOne) SetNillableExtra(s *string) *CreditAllocatedUpdateOne {
	if s != nil {
		cauo.SetExtra(*s)
	}
	return cauo
}

// ClearExtra clears the value of the "extra" field.
func (cauo *CreditAllocatedUpdateOne) ClearExtra() *CreditAllocatedUpdateOne {
	cauo.mutation.ClearExtra()
	return cauo
}

// Mutation returns the CreditAllocatedMutation object of the builder.
func (cauo *CreditAllocatedUpdateOne) Mutation() *CreditAllocatedMutation {
	return cauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CreditAllocatedUpdateOne) Select(field string, fields ...string) *CreditAllocatedUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CreditAllocated entity.
func (cauo *CreditAllocatedUpdateOne) Save(ctx context.Context) (*CreditAllocated, error) {
	var (
		err  error
		node *CreditAllocated
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
			mutation, ok := m.(*CreditAllocatedMutation)
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
		nv, ok := v.(*CreditAllocated)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CreditAllocatedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CreditAllocatedUpdateOne) SaveX(ctx context.Context) *CreditAllocated {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CreditAllocatedUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CreditAllocatedUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cauo *CreditAllocatedUpdateOne) defaults() error {
	if _, ok := cauo.mutation.UpdatedAt(); !ok {
		if creditallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized creditallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := creditallocated.UpdateDefaultUpdatedAt()
		cauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cauo *CreditAllocatedUpdateOne) check() error {
	if v, ok := cauo.mutation.Extra(); ok {
		if err := creditallocated.ExtraValidator(v); err != nil {
			return &ValidationError{Name: "extra", err: fmt.Errorf(`ent: validator failed for field "CreditAllocated.extra": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cauo *CreditAllocatedUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CreditAllocatedUpdateOne {
	cauo.modifiers = append(cauo.modifiers, modifiers...)
	return cauo
}

func (cauo *CreditAllocatedUpdateOne) sqlSave(ctx context.Context) (_node *CreditAllocated, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   creditallocated.Table,
			Columns: creditallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: creditallocated.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CreditAllocated.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creditallocated.FieldID)
		for _, f := range fields {
			if !creditallocated.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != creditallocated.FieldID {
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
			Column: creditallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: creditallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldEntID,
		})
	}
	if value, ok := cauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldAppID,
		})
	}
	if cauo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: creditallocated.FieldAppID,
		})
	}
	if value, ok := cauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: creditallocated.FieldUserID,
		})
	}
	if cauo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: creditallocated.FieldUserID,
		})
	}
	if value, ok := cauo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: creditallocated.FieldValue,
		})
	}
	if cauo.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: creditallocated.FieldValue,
		})
	}
	if value, ok := cauo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: creditallocated.FieldExtra,
		})
	}
	if cauo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: creditallocated.FieldExtra,
		})
	}
	_spec.Modifiers = cauo.modifiers
	_node = &CreditAllocated{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creditallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}