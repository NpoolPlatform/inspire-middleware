// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPaymentStatementUpdate is the builder for updating OrderPaymentStatement entities.
type OrderPaymentStatementUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderPaymentStatementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderPaymentStatementUpdate builder.
func (opsu *OrderPaymentStatementUpdate) Where(ps ...predicate.OrderPaymentStatement) *OrderPaymentStatementUpdate {
	opsu.mutation.Where(ps...)
	return opsu
}

// SetCreatedAt sets the "created_at" field.
func (opsu *OrderPaymentStatementUpdate) SetCreatedAt(u uint32) *OrderPaymentStatementUpdate {
	opsu.mutation.ResetCreatedAt()
	opsu.mutation.SetCreatedAt(u)
	return opsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableCreatedAt(u *uint32) *OrderPaymentStatementUpdate {
	if u != nil {
		opsu.SetCreatedAt(*u)
	}
	return opsu
}

// AddCreatedAt adds u to the "created_at" field.
func (opsu *OrderPaymentStatementUpdate) AddCreatedAt(u int32) *OrderPaymentStatementUpdate {
	opsu.mutation.AddCreatedAt(u)
	return opsu
}

// SetUpdatedAt sets the "updated_at" field.
func (opsu *OrderPaymentStatementUpdate) SetUpdatedAt(u uint32) *OrderPaymentStatementUpdate {
	opsu.mutation.ResetUpdatedAt()
	opsu.mutation.SetUpdatedAt(u)
	return opsu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (opsu *OrderPaymentStatementUpdate) AddUpdatedAt(u int32) *OrderPaymentStatementUpdate {
	opsu.mutation.AddUpdatedAt(u)
	return opsu
}

// SetDeletedAt sets the "deleted_at" field.
func (opsu *OrderPaymentStatementUpdate) SetDeletedAt(u uint32) *OrderPaymentStatementUpdate {
	opsu.mutation.ResetDeletedAt()
	opsu.mutation.SetDeletedAt(u)
	return opsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableDeletedAt(u *uint32) *OrderPaymentStatementUpdate {
	if u != nil {
		opsu.SetDeletedAt(*u)
	}
	return opsu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (opsu *OrderPaymentStatementUpdate) AddDeletedAt(u int32) *OrderPaymentStatementUpdate {
	opsu.mutation.AddDeletedAt(u)
	return opsu
}

// SetEntID sets the "ent_id" field.
func (opsu *OrderPaymentStatementUpdate) SetEntID(u uuid.UUID) *OrderPaymentStatementUpdate {
	opsu.mutation.SetEntID(u)
	return opsu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableEntID(u *uuid.UUID) *OrderPaymentStatementUpdate {
	if u != nil {
		opsu.SetEntID(*u)
	}
	return opsu
}

// SetStatementID sets the "statement_id" field.
func (opsu *OrderPaymentStatementUpdate) SetStatementID(u uuid.UUID) *OrderPaymentStatementUpdate {
	opsu.mutation.SetStatementID(u)
	return opsu
}

// SetNillableStatementID sets the "statement_id" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableStatementID(u *uuid.UUID) *OrderPaymentStatementUpdate {
	if u != nil {
		opsu.SetStatementID(*u)
	}
	return opsu
}

// ClearStatementID clears the value of the "statement_id" field.
func (opsu *OrderPaymentStatementUpdate) ClearStatementID() *OrderPaymentStatementUpdate {
	opsu.mutation.ClearStatementID()
	return opsu
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (opsu *OrderPaymentStatementUpdate) SetPaymentCoinTypeID(u uuid.UUID) *OrderPaymentStatementUpdate {
	opsu.mutation.SetPaymentCoinTypeID(u)
	return opsu
}

// SetNillablePaymentCoinTypeID sets the "payment_coin_type_id" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillablePaymentCoinTypeID(u *uuid.UUID) *OrderPaymentStatementUpdate {
	if u != nil {
		opsu.SetPaymentCoinTypeID(*u)
	}
	return opsu
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (opsu *OrderPaymentStatementUpdate) ClearPaymentCoinTypeID() *OrderPaymentStatementUpdate {
	opsu.mutation.ClearPaymentCoinTypeID()
	return opsu
}

// SetAmount sets the "amount" field.
func (opsu *OrderPaymentStatementUpdate) SetAmount(d decimal.Decimal) *OrderPaymentStatementUpdate {
	opsu.mutation.SetAmount(d)
	return opsu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableAmount(d *decimal.Decimal) *OrderPaymentStatementUpdate {
	if d != nil {
		opsu.SetAmount(*d)
	}
	return opsu
}

// ClearAmount clears the value of the "amount" field.
func (opsu *OrderPaymentStatementUpdate) ClearAmount() *OrderPaymentStatementUpdate {
	opsu.mutation.ClearAmount()
	return opsu
}

// SetCommissionAmount sets the "commission_amount" field.
func (opsu *OrderPaymentStatementUpdate) SetCommissionAmount(d decimal.Decimal) *OrderPaymentStatementUpdate {
	opsu.mutation.SetCommissionAmount(d)
	return opsu
}

// SetNillableCommissionAmount sets the "commission_amount" field if the given value is not nil.
func (opsu *OrderPaymentStatementUpdate) SetNillableCommissionAmount(d *decimal.Decimal) *OrderPaymentStatementUpdate {
	if d != nil {
		opsu.SetCommissionAmount(*d)
	}
	return opsu
}

// ClearCommissionAmount clears the value of the "commission_amount" field.
func (opsu *OrderPaymentStatementUpdate) ClearCommissionAmount() *OrderPaymentStatementUpdate {
	opsu.mutation.ClearCommissionAmount()
	return opsu
}

// Mutation returns the OrderPaymentStatementMutation object of the builder.
func (opsu *OrderPaymentStatementUpdate) Mutation() *OrderPaymentStatementMutation {
	return opsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opsu *OrderPaymentStatementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := opsu.defaults(); err != nil {
		return 0, err
	}
	if len(opsu.hooks) == 0 {
		affected, err = opsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPaymentStatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			opsu.mutation = mutation
			affected, err = opsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(opsu.hooks) - 1; i >= 0; i-- {
			if opsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, opsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (opsu *OrderPaymentStatementUpdate) SaveX(ctx context.Context) int {
	affected, err := opsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opsu *OrderPaymentStatementUpdate) Exec(ctx context.Context) error {
	_, err := opsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opsu *OrderPaymentStatementUpdate) ExecX(ctx context.Context) {
	if err := opsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opsu *OrderPaymentStatementUpdate) defaults() error {
	if _, ok := opsu.mutation.UpdatedAt(); !ok {
		if orderpaymentstatement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpaymentstatement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderpaymentstatement.UpdateDefaultUpdatedAt()
		opsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opsu *OrderPaymentStatementUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPaymentStatementUpdate {
	opsu.modifiers = append(opsu.modifiers, modifiers...)
	return opsu
}

func (opsu *OrderPaymentStatementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymentstatement.Table,
			Columns: orderpaymentstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymentstatement.FieldID,
			},
		},
	}
	if ps := opsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldCreatedAt,
		})
	}
	if value, ok := opsu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldCreatedAt,
		})
	}
	if value, ok := opsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldUpdatedAt,
		})
	}
	if value, ok := opsu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldUpdatedAt,
		})
	}
	if value, ok := opsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldDeletedAt,
		})
	}
	if value, ok := opsu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldDeletedAt,
		})
	}
	if value, ok := opsu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldEntID,
		})
	}
	if value, ok := opsu.mutation.StatementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldStatementID,
		})
	}
	if opsu.mutation.StatementIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymentstatement.FieldStatementID,
		})
	}
	if value, ok := opsu.mutation.PaymentCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldPaymentCoinTypeID,
		})
	}
	if opsu.mutation.PaymentCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymentstatement.FieldPaymentCoinTypeID,
		})
	}
	if value, ok := opsu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymentstatement.FieldAmount,
		})
	}
	if opsu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymentstatement.FieldAmount,
		})
	}
	if value, ok := opsu.mutation.CommissionAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymentstatement.FieldCommissionAmount,
		})
	}
	if opsu.mutation.CommissionAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymentstatement.FieldCommissionAmount,
		})
	}
	_spec.Modifiers = opsu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, opsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpaymentstatement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderPaymentStatementUpdateOne is the builder for updating a single OrderPaymentStatement entity.
type OrderPaymentStatementUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderPaymentStatementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetCreatedAt(u uint32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ResetCreatedAt()
	opsuo.mutation.SetCreatedAt(u)
	return opsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableCreatedAt(u *uint32) *OrderPaymentStatementUpdateOne {
	if u != nil {
		opsuo.SetCreatedAt(*u)
	}
	return opsuo
}

// AddCreatedAt adds u to the "created_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) AddCreatedAt(u int32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.AddCreatedAt(u)
	return opsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetUpdatedAt(u uint32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ResetUpdatedAt()
	opsuo.mutation.SetUpdatedAt(u)
	return opsuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) AddUpdatedAt(u int32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.AddUpdatedAt(u)
	return opsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetDeletedAt(u uint32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ResetDeletedAt()
	opsuo.mutation.SetDeletedAt(u)
	return opsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableDeletedAt(u *uint32) *OrderPaymentStatementUpdateOne {
	if u != nil {
		opsuo.SetDeletedAt(*u)
	}
	return opsuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (opsuo *OrderPaymentStatementUpdateOne) AddDeletedAt(u int32) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.AddDeletedAt(u)
	return opsuo
}

// SetEntID sets the "ent_id" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetEntID(u uuid.UUID) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.SetEntID(u)
	return opsuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderPaymentStatementUpdateOne {
	if u != nil {
		opsuo.SetEntID(*u)
	}
	return opsuo
}

// SetStatementID sets the "statement_id" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetStatementID(u uuid.UUID) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.SetStatementID(u)
	return opsuo
}

// SetNillableStatementID sets the "statement_id" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableStatementID(u *uuid.UUID) *OrderPaymentStatementUpdateOne {
	if u != nil {
		opsuo.SetStatementID(*u)
	}
	return opsuo
}

// ClearStatementID clears the value of the "statement_id" field.
func (opsuo *OrderPaymentStatementUpdateOne) ClearStatementID() *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ClearStatementID()
	return opsuo
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetPaymentCoinTypeID(u uuid.UUID) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.SetPaymentCoinTypeID(u)
	return opsuo
}

// SetNillablePaymentCoinTypeID sets the "payment_coin_type_id" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillablePaymentCoinTypeID(u *uuid.UUID) *OrderPaymentStatementUpdateOne {
	if u != nil {
		opsuo.SetPaymentCoinTypeID(*u)
	}
	return opsuo
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (opsuo *OrderPaymentStatementUpdateOne) ClearPaymentCoinTypeID() *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ClearPaymentCoinTypeID()
	return opsuo
}

// SetAmount sets the "amount" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetAmount(d decimal.Decimal) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.SetAmount(d)
	return opsuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableAmount(d *decimal.Decimal) *OrderPaymentStatementUpdateOne {
	if d != nil {
		opsuo.SetAmount(*d)
	}
	return opsuo
}

// ClearAmount clears the value of the "amount" field.
func (opsuo *OrderPaymentStatementUpdateOne) ClearAmount() *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ClearAmount()
	return opsuo
}

// SetCommissionAmount sets the "commission_amount" field.
func (opsuo *OrderPaymentStatementUpdateOne) SetCommissionAmount(d decimal.Decimal) *OrderPaymentStatementUpdateOne {
	opsuo.mutation.SetCommissionAmount(d)
	return opsuo
}

// SetNillableCommissionAmount sets the "commission_amount" field if the given value is not nil.
func (opsuo *OrderPaymentStatementUpdateOne) SetNillableCommissionAmount(d *decimal.Decimal) *OrderPaymentStatementUpdateOne {
	if d != nil {
		opsuo.SetCommissionAmount(*d)
	}
	return opsuo
}

// ClearCommissionAmount clears the value of the "commission_amount" field.
func (opsuo *OrderPaymentStatementUpdateOne) ClearCommissionAmount() *OrderPaymentStatementUpdateOne {
	opsuo.mutation.ClearCommissionAmount()
	return opsuo
}

// Mutation returns the OrderPaymentStatementMutation object of the builder.
func (opsuo *OrderPaymentStatementUpdateOne) Mutation() *OrderPaymentStatementMutation {
	return opsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opsuo *OrderPaymentStatementUpdateOne) Select(field string, fields ...string) *OrderPaymentStatementUpdateOne {
	opsuo.fields = append([]string{field}, fields...)
	return opsuo
}

// Save executes the query and returns the updated OrderPaymentStatement entity.
func (opsuo *OrderPaymentStatementUpdateOne) Save(ctx context.Context) (*OrderPaymentStatement, error) {
	var (
		err  error
		node *OrderPaymentStatement
	)
	if err := opsuo.defaults(); err != nil {
		return nil, err
	}
	if len(opsuo.hooks) == 0 {
		node, err = opsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPaymentStatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			opsuo.mutation = mutation
			node, err = opsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(opsuo.hooks) - 1; i >= 0; i-- {
			if opsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, opsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderPaymentStatement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderPaymentStatementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (opsuo *OrderPaymentStatementUpdateOne) SaveX(ctx context.Context) *OrderPaymentStatement {
	node, err := opsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opsuo *OrderPaymentStatementUpdateOne) Exec(ctx context.Context) error {
	_, err := opsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opsuo *OrderPaymentStatementUpdateOne) ExecX(ctx context.Context) {
	if err := opsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opsuo *OrderPaymentStatementUpdateOne) defaults() error {
	if _, ok := opsuo.mutation.UpdatedAt(); !ok {
		if orderpaymentstatement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpaymentstatement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderpaymentstatement.UpdateDefaultUpdatedAt()
		opsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opsuo *OrderPaymentStatementUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPaymentStatementUpdateOne {
	opsuo.modifiers = append(opsuo.modifiers, modifiers...)
	return opsuo
}

func (opsuo *OrderPaymentStatementUpdateOne) sqlSave(ctx context.Context) (_node *OrderPaymentStatement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymentstatement.Table,
			Columns: orderpaymentstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymentstatement.FieldID,
			},
		},
	}
	id, ok := opsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderPaymentStatement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpaymentstatement.FieldID)
		for _, f := range fields {
			if !orderpaymentstatement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderpaymentstatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldCreatedAt,
		})
	}
	if value, ok := opsuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldCreatedAt,
		})
	}
	if value, ok := opsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldUpdatedAt,
		})
	}
	if value, ok := opsuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldUpdatedAt,
		})
	}
	if value, ok := opsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldDeletedAt,
		})
	}
	if value, ok := opsuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymentstatement.FieldDeletedAt,
		})
	}
	if value, ok := opsuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldEntID,
		})
	}
	if value, ok := opsuo.mutation.StatementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldStatementID,
		})
	}
	if opsuo.mutation.StatementIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymentstatement.FieldStatementID,
		})
	}
	if value, ok := opsuo.mutation.PaymentCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymentstatement.FieldPaymentCoinTypeID,
		})
	}
	if opsuo.mutation.PaymentCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymentstatement.FieldPaymentCoinTypeID,
		})
	}
	if value, ok := opsuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymentstatement.FieldAmount,
		})
	}
	if opsuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymentstatement.FieldAmount,
		})
	}
	if value, ok := opsuo.mutation.CommissionAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymentstatement.FieldCommissionAmount,
		})
	}
	if opsuo.mutation.CommissionAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymentstatement.FieldCommissionAmount,
		})
	}
	_spec.Modifiers = opsuo.modifiers
	_node = &OrderPaymentStatement{config: opsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpaymentstatement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
