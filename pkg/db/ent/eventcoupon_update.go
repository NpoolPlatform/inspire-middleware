// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/eventcoupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// EventCouponUpdate is the builder for updating EventCoupon entities.
type EventCouponUpdate struct {
	config
	hooks     []Hook
	mutation  *EventCouponMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EventCouponUpdate builder.
func (ecu *EventCouponUpdate) Where(ps ...predicate.EventCoupon) *EventCouponUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetCreatedAt sets the "created_at" field.
func (ecu *EventCouponUpdate) SetCreatedAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetCreatedAt()
	ecu.mutation.SetCreatedAt(u)
	return ecu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableCreatedAt(u *uint32) *EventCouponUpdate {
	if u != nil {
		ecu.SetCreatedAt(*u)
	}
	return ecu
}

// AddCreatedAt adds u to the "created_at" field.
func (ecu *EventCouponUpdate) AddCreatedAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddCreatedAt(u)
	return ecu
}

// SetUpdatedAt sets the "updated_at" field.
func (ecu *EventCouponUpdate) SetUpdatedAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetUpdatedAt()
	ecu.mutation.SetUpdatedAt(u)
	return ecu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ecu *EventCouponUpdate) AddUpdatedAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddUpdatedAt(u)
	return ecu
}

// SetDeletedAt sets the "deleted_at" field.
func (ecu *EventCouponUpdate) SetDeletedAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetDeletedAt()
	ecu.mutation.SetDeletedAt(u)
	return ecu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableDeletedAt(u *uint32) *EventCouponUpdate {
	if u != nil {
		ecu.SetDeletedAt(*u)
	}
	return ecu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ecu *EventCouponUpdate) AddDeletedAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddDeletedAt(u)
	return ecu
}

// SetEntID sets the "ent_id" field.
func (ecu *EventCouponUpdate) SetEntID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetEntID(u)
	return ecu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableEntID(u *uuid.UUID) *EventCouponUpdate {
	if u != nil {
		ecu.SetEntID(*u)
	}
	return ecu
}

// SetAppID sets the "app_id" field.
func (ecu *EventCouponUpdate) SetAppID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetAppID(u)
	return ecu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableAppID(u *uuid.UUID) *EventCouponUpdate {
	if u != nil {
		ecu.SetAppID(*u)
	}
	return ecu
}

// ClearAppID clears the value of the "app_id" field.
func (ecu *EventCouponUpdate) ClearAppID() *EventCouponUpdate {
	ecu.mutation.ClearAppID()
	return ecu
}

// SetEventID sets the "event_id" field.
func (ecu *EventCouponUpdate) SetEventID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetEventID(u)
	return ecu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableEventID(u *uuid.UUID) *EventCouponUpdate {
	if u != nil {
		ecu.SetEventID(*u)
	}
	return ecu
}

// ClearEventID clears the value of the "event_id" field.
func (ecu *EventCouponUpdate) ClearEventID() *EventCouponUpdate {
	ecu.mutation.ClearEventID()
	return ecu
}

// SetCouponID sets the "coupon_id" field.
func (ecu *EventCouponUpdate) SetCouponID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetCouponID(u)
	return ecu
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableCouponID(u *uuid.UUID) *EventCouponUpdate {
	if u != nil {
		ecu.SetCouponID(*u)
	}
	return ecu
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ecu *EventCouponUpdate) ClearCouponID() *EventCouponUpdate {
	ecu.mutation.ClearCouponID()
	return ecu
}

// Mutation returns the EventCouponMutation object of the builder.
func (ecu *EventCouponUpdate) Mutation() *EventCouponMutation {
	return ecu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EventCouponUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ecu.defaults(); err != nil {
		return 0, err
	}
	if len(ecu.hooks) == 0 {
		affected, err = ecu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ecu.mutation = mutation
			affected, err = ecu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ecu.hooks) - 1; i >= 0; i-- {
			if ecu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EventCouponUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EventCouponUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EventCouponUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecu *EventCouponUpdate) defaults() error {
	if _, ok := ecu.mutation.UpdatedAt(); !ok {
		if eventcoupon.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized eventcoupon.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := eventcoupon.UpdateDefaultUpdatedAt()
		ecu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ecu *EventCouponUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventCouponUpdate {
	ecu.modifiers = append(ecu.modifiers, modifiers...)
	return ecu
}

func (ecu *EventCouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventcoupon.Table,
			Columns: eventcoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: eventcoupon.FieldID,
			},
		},
	}
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreatedAt,
		})
	}
	if value, ok := ecu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreatedAt,
		})
	}
	if value, ok := ecu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ecu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ecu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeletedAt,
		})
	}
	if value, ok := ecu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeletedAt,
		})
	}
	if value, ok := ecu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldEntID,
		})
	}
	if value, ok := ecu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldAppID,
		})
	}
	if ecu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldAppID,
		})
	}
	if value, ok := ecu.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldEventID,
		})
	}
	if ecu.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldEventID,
		})
	}
	if value, ok := ecu.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldCouponID,
		})
	}
	if ecu.mutation.CouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldCouponID,
		})
	}
	_spec.Modifiers = ecu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventcoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventCouponUpdateOne is the builder for updating a single EventCoupon entity.
type EventCouponUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EventCouponMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ecuo *EventCouponUpdateOne) SetCreatedAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetCreatedAt()
	ecuo.mutation.SetCreatedAt(u)
	return ecuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableCreatedAt(u *uint32) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetCreatedAt(*u)
	}
	return ecuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ecuo *EventCouponUpdateOne) AddCreatedAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddCreatedAt(u)
	return ecuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ecuo *EventCouponUpdateOne) SetUpdatedAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetUpdatedAt()
	ecuo.mutation.SetUpdatedAt(u)
	return ecuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ecuo *EventCouponUpdateOne) AddUpdatedAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddUpdatedAt(u)
	return ecuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ecuo *EventCouponUpdateOne) SetDeletedAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetDeletedAt()
	ecuo.mutation.SetDeletedAt(u)
	return ecuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableDeletedAt(u *uint32) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetDeletedAt(*u)
	}
	return ecuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ecuo *EventCouponUpdateOne) AddDeletedAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddDeletedAt(u)
	return ecuo
}

// SetEntID sets the "ent_id" field.
func (ecuo *EventCouponUpdateOne) SetEntID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetEntID(u)
	return ecuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableEntID(u *uuid.UUID) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetEntID(*u)
	}
	return ecuo
}

// SetAppID sets the "app_id" field.
func (ecuo *EventCouponUpdateOne) SetAppID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetAppID(u)
	return ecuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableAppID(u *uuid.UUID) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetAppID(*u)
	}
	return ecuo
}

// ClearAppID clears the value of the "app_id" field.
func (ecuo *EventCouponUpdateOne) ClearAppID() *EventCouponUpdateOne {
	ecuo.mutation.ClearAppID()
	return ecuo
}

// SetEventID sets the "event_id" field.
func (ecuo *EventCouponUpdateOne) SetEventID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetEventID(u)
	return ecuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableEventID(u *uuid.UUID) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetEventID(*u)
	}
	return ecuo
}

// ClearEventID clears the value of the "event_id" field.
func (ecuo *EventCouponUpdateOne) ClearEventID() *EventCouponUpdateOne {
	ecuo.mutation.ClearEventID()
	return ecuo
}

// SetCouponID sets the "coupon_id" field.
func (ecuo *EventCouponUpdateOne) SetCouponID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetCouponID(u)
	return ecuo
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableCouponID(u *uuid.UUID) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetCouponID(*u)
	}
	return ecuo
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ecuo *EventCouponUpdateOne) ClearCouponID() *EventCouponUpdateOne {
	ecuo.mutation.ClearCouponID()
	return ecuo
}

// Mutation returns the EventCouponMutation object of the builder.
func (ecuo *EventCouponUpdateOne) Mutation() *EventCouponMutation {
	return ecuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EventCouponUpdateOne) Select(field string, fields ...string) *EventCouponUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EventCoupon entity.
func (ecuo *EventCouponUpdateOne) Save(ctx context.Context) (*EventCoupon, error) {
	var (
		err  error
		node *EventCoupon
	)
	if err := ecuo.defaults(); err != nil {
		return nil, err
	}
	if len(ecuo.hooks) == 0 {
		node, err = ecuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ecuo.mutation = mutation
			node, err = ecuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ecuo.hooks) - 1; i >= 0; i-- {
			if ecuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ecuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EventCoupon)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventCouponMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EventCouponUpdateOne) SaveX(ctx context.Context) *EventCoupon {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EventCouponUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EventCouponUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecuo *EventCouponUpdateOne) defaults() error {
	if _, ok := ecuo.mutation.UpdatedAt(); !ok {
		if eventcoupon.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized eventcoupon.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := eventcoupon.UpdateDefaultUpdatedAt()
		ecuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ecuo *EventCouponUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventCouponUpdateOne {
	ecuo.modifiers = append(ecuo.modifiers, modifiers...)
	return ecuo
}

func (ecuo *EventCouponUpdateOne) sqlSave(ctx context.Context) (_node *EventCoupon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventcoupon.Table,
			Columns: eventcoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: eventcoupon.FieldID,
			},
		},
	}
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventCoupon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventcoupon.FieldID)
		for _, f := range fields {
			if !eventcoupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventcoupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreatedAt,
		})
	}
	if value, ok := ecuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreatedAt,
		})
	}
	if value, ok := ecuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ecuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ecuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeletedAt,
		})
	}
	if value, ok := ecuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeletedAt,
		})
	}
	if value, ok := ecuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldEntID,
		})
	}
	if value, ok := ecuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldAppID,
		})
	}
	if ecuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldAppID,
		})
	}
	if value, ok := ecuo.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldEventID,
		})
	}
	if ecuo.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldEventID,
		})
	}
	if value, ok := ecuo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldCouponID,
		})
	}
	if ecuo.mutation.CouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: eventcoupon.FieldCouponID,
		})
	}
	_spec.Modifiers = ecuo.modifiers
	_node = &EventCoupon{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventcoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
