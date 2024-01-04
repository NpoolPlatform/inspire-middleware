// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponcoin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CouponCoinUpdate is the builder for updating CouponCoin entities.
type CouponCoinUpdate struct {
	config
	hooks     []Hook
	mutation  *CouponCoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CouponCoinUpdate builder.
func (ccu *CouponCoinUpdate) Where(ps ...predicate.CouponCoin) *CouponCoinUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *CouponCoinUpdate) SetCreatedAt(u uint32) *CouponCoinUpdate {
	ccu.mutation.ResetCreatedAt()
	ccu.mutation.SetCreatedAt(u)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *CouponCoinUpdate) SetNillableCreatedAt(u *uint32) *CouponCoinUpdate {
	if u != nil {
		ccu.SetCreatedAt(*u)
	}
	return ccu
}

// AddCreatedAt adds u to the "created_at" field.
func (ccu *CouponCoinUpdate) AddCreatedAt(u int32) *CouponCoinUpdate {
	ccu.mutation.AddCreatedAt(u)
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *CouponCoinUpdate) SetUpdatedAt(u uint32) *CouponCoinUpdate {
	ccu.mutation.ResetUpdatedAt()
	ccu.mutation.SetUpdatedAt(u)
	return ccu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ccu *CouponCoinUpdate) AddUpdatedAt(u int32) *CouponCoinUpdate {
	ccu.mutation.AddUpdatedAt(u)
	return ccu
}

// SetDeletedAt sets the "deleted_at" field.
func (ccu *CouponCoinUpdate) SetDeletedAt(u uint32) *CouponCoinUpdate {
	ccu.mutation.ResetDeletedAt()
	ccu.mutation.SetDeletedAt(u)
	return ccu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ccu *CouponCoinUpdate) SetNillableDeletedAt(u *uint32) *CouponCoinUpdate {
	if u != nil {
		ccu.SetDeletedAt(*u)
	}
	return ccu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ccu *CouponCoinUpdate) AddDeletedAt(u int32) *CouponCoinUpdate {
	ccu.mutation.AddDeletedAt(u)
	return ccu
}

// SetEntID sets the "ent_id" field.
func (ccu *CouponCoinUpdate) SetEntID(u uuid.UUID) *CouponCoinUpdate {
	ccu.mutation.SetEntID(u)
	return ccu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ccu *CouponCoinUpdate) SetNillableEntID(u *uuid.UUID) *CouponCoinUpdate {
	if u != nil {
		ccu.SetEntID(*u)
	}
	return ccu
}

// SetAppID sets the "app_id" field.
func (ccu *CouponCoinUpdate) SetAppID(u uuid.UUID) *CouponCoinUpdate {
	ccu.mutation.SetAppID(u)
	return ccu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ccu *CouponCoinUpdate) SetNillableAppID(u *uuid.UUID) *CouponCoinUpdate {
	if u != nil {
		ccu.SetAppID(*u)
	}
	return ccu
}

// ClearAppID clears the value of the "app_id" field.
func (ccu *CouponCoinUpdate) ClearAppID() *CouponCoinUpdate {
	ccu.mutation.ClearAppID()
	return ccu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccu *CouponCoinUpdate) SetCoinTypeID(u uuid.UUID) *CouponCoinUpdate {
	ccu.mutation.SetCoinTypeID(u)
	return ccu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ccu *CouponCoinUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CouponCoinUpdate {
	if u != nil {
		ccu.SetCoinTypeID(*u)
	}
	return ccu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ccu *CouponCoinUpdate) ClearCoinTypeID() *CouponCoinUpdate {
	ccu.mutation.ClearCoinTypeID()
	return ccu
}

// Mutation returns the CouponCoinMutation object of the builder.
func (ccu *CouponCoinUpdate) Mutation() *CouponCoinMutation {
	return ccu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CouponCoinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ccu.defaults(); err != nil {
		return 0, err
	}
	if len(ccu.hooks) == 0 {
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			if ccu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CouponCoinUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CouponCoinUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CouponCoinUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *CouponCoinUpdate) defaults() error {
	if _, ok := ccu.mutation.UpdatedAt(); !ok {
		if couponcoin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponcoin.UpdateDefaultUpdatedAt()
		ccu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccu *CouponCoinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponCoinUpdate {
	ccu.modifiers = append(ccu.modifiers, modifiers...)
	return ccu
}

func (ccu *CouponCoinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponcoin.Table,
			Columns: couponcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: couponcoin.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldUpdatedAt,
		})
	}
	if value, ok := ccu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldUpdatedAt,
		})
	}
	if value, ok := ccu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldDeletedAt,
		})
	}
	if value, ok := ccu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldDeletedAt,
		})
	}
	if value, ok := ccu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldEntID,
		})
	}
	if value, ok := ccu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldAppID,
		})
	}
	if ccu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponcoin.FieldAppID,
		})
	}
	if value, ok := ccu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldCoinTypeID,
		})
	}
	if ccu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponcoin.FieldCoinTypeID,
		})
	}
	_spec.Modifiers = ccu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponcoin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponCoinUpdateOne is the builder for updating a single CouponCoin entity.
type CouponCoinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CouponCoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *CouponCoinUpdateOne) SetCreatedAt(u uint32) *CouponCoinUpdateOne {
	ccuo.mutation.ResetCreatedAt()
	ccuo.mutation.SetCreatedAt(u)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *CouponCoinUpdateOne) SetNillableCreatedAt(u *uint32) *CouponCoinUpdateOne {
	if u != nil {
		ccuo.SetCreatedAt(*u)
	}
	return ccuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ccuo *CouponCoinUpdateOne) AddCreatedAt(u int32) *CouponCoinUpdateOne {
	ccuo.mutation.AddCreatedAt(u)
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *CouponCoinUpdateOne) SetUpdatedAt(u uint32) *CouponCoinUpdateOne {
	ccuo.mutation.ResetUpdatedAt()
	ccuo.mutation.SetUpdatedAt(u)
	return ccuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ccuo *CouponCoinUpdateOne) AddUpdatedAt(u int32) *CouponCoinUpdateOne {
	ccuo.mutation.AddUpdatedAt(u)
	return ccuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ccuo *CouponCoinUpdateOne) SetDeletedAt(u uint32) *CouponCoinUpdateOne {
	ccuo.mutation.ResetDeletedAt()
	ccuo.mutation.SetDeletedAt(u)
	return ccuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ccuo *CouponCoinUpdateOne) SetNillableDeletedAt(u *uint32) *CouponCoinUpdateOne {
	if u != nil {
		ccuo.SetDeletedAt(*u)
	}
	return ccuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ccuo *CouponCoinUpdateOne) AddDeletedAt(u int32) *CouponCoinUpdateOne {
	ccuo.mutation.AddDeletedAt(u)
	return ccuo
}

// SetEntID sets the "ent_id" field.
func (ccuo *CouponCoinUpdateOne) SetEntID(u uuid.UUID) *CouponCoinUpdateOne {
	ccuo.mutation.SetEntID(u)
	return ccuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ccuo *CouponCoinUpdateOne) SetNillableEntID(u *uuid.UUID) *CouponCoinUpdateOne {
	if u != nil {
		ccuo.SetEntID(*u)
	}
	return ccuo
}

// SetAppID sets the "app_id" field.
func (ccuo *CouponCoinUpdateOne) SetAppID(u uuid.UUID) *CouponCoinUpdateOne {
	ccuo.mutation.SetAppID(u)
	return ccuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ccuo *CouponCoinUpdateOne) SetNillableAppID(u *uuid.UUID) *CouponCoinUpdateOne {
	if u != nil {
		ccuo.SetAppID(*u)
	}
	return ccuo
}

// ClearAppID clears the value of the "app_id" field.
func (ccuo *CouponCoinUpdateOne) ClearAppID() *CouponCoinUpdateOne {
	ccuo.mutation.ClearAppID()
	return ccuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccuo *CouponCoinUpdateOne) SetCoinTypeID(u uuid.UUID) *CouponCoinUpdateOne {
	ccuo.mutation.SetCoinTypeID(u)
	return ccuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ccuo *CouponCoinUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CouponCoinUpdateOne {
	if u != nil {
		ccuo.SetCoinTypeID(*u)
	}
	return ccuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ccuo *CouponCoinUpdateOne) ClearCoinTypeID() *CouponCoinUpdateOne {
	ccuo.mutation.ClearCoinTypeID()
	return ccuo
}

// Mutation returns the CouponCoinMutation object of the builder.
func (ccuo *CouponCoinUpdateOne) Mutation() *CouponCoinMutation {
	return ccuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CouponCoinUpdateOne) Select(field string, fields ...string) *CouponCoinUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CouponCoin entity.
func (ccuo *CouponCoinUpdateOne) Save(ctx context.Context) (*CouponCoin, error) {
	var (
		err  error
		node *CouponCoin
	)
	if err := ccuo.defaults(); err != nil {
		return nil, err
	}
	if len(ccuo.hooks) == 0 {
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			if ccuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CouponCoin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponCoinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CouponCoinUpdateOne) SaveX(ctx context.Context) *CouponCoin {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CouponCoinUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CouponCoinUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *CouponCoinUpdateOne) defaults() error {
	if _, ok := ccuo.mutation.UpdatedAt(); !ok {
		if couponcoin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponcoin.UpdateDefaultUpdatedAt()
		ccuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ccuo *CouponCoinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponCoinUpdateOne {
	ccuo.modifiers = append(ccuo.modifiers, modifiers...)
	return ccuo
}

func (ccuo *CouponCoinUpdateOne) sqlSave(ctx context.Context) (_node *CouponCoin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponcoin.Table,
			Columns: couponcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: couponcoin.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CouponCoin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponcoin.FieldID)
		for _, f := range fields {
			if !couponcoin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != couponcoin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldUpdatedAt,
		})
	}
	if value, ok := ccuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldUpdatedAt,
		})
	}
	if value, ok := ccuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldDeletedAt,
		})
	}
	if value, ok := ccuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldDeletedAt,
		})
	}
	if value, ok := ccuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldEntID,
		})
	}
	if value, ok := ccuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldAppID,
		})
	}
	if ccuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponcoin.FieldAppID,
		})
	}
	if value, ok := ccuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldCoinTypeID,
		})
	}
	if ccuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponcoin.FieldCoinTypeID,
		})
	}
	_spec.Modifiers = ccuo.modifiers
	_node = &CouponCoin{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponcoin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}