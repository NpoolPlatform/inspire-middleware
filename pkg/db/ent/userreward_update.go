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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/userreward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UserRewardUpdate is the builder for updating UserReward entities.
type UserRewardUpdate struct {
	config
	hooks     []Hook
	mutation  *UserRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserRewardUpdate builder.
func (uru *UserRewardUpdate) Where(ps ...predicate.UserReward) *UserRewardUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetCreatedAt sets the "created_at" field.
func (uru *UserRewardUpdate) SetCreatedAt(u uint32) *UserRewardUpdate {
	uru.mutation.ResetCreatedAt()
	uru.mutation.SetCreatedAt(u)
	return uru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableCreatedAt(u *uint32) *UserRewardUpdate {
	if u != nil {
		uru.SetCreatedAt(*u)
	}
	return uru
}

// AddCreatedAt adds u to the "created_at" field.
func (uru *UserRewardUpdate) AddCreatedAt(u int32) *UserRewardUpdate {
	uru.mutation.AddCreatedAt(u)
	return uru
}

// SetUpdatedAt sets the "updated_at" field.
func (uru *UserRewardUpdate) SetUpdatedAt(u uint32) *UserRewardUpdate {
	uru.mutation.ResetUpdatedAt()
	uru.mutation.SetUpdatedAt(u)
	return uru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (uru *UserRewardUpdate) AddUpdatedAt(u int32) *UserRewardUpdate {
	uru.mutation.AddUpdatedAt(u)
	return uru
}

// SetDeletedAt sets the "deleted_at" field.
func (uru *UserRewardUpdate) SetDeletedAt(u uint32) *UserRewardUpdate {
	uru.mutation.ResetDeletedAt()
	uru.mutation.SetDeletedAt(u)
	return uru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableDeletedAt(u *uint32) *UserRewardUpdate {
	if u != nil {
		uru.SetDeletedAt(*u)
	}
	return uru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (uru *UserRewardUpdate) AddDeletedAt(u int32) *UserRewardUpdate {
	uru.mutation.AddDeletedAt(u)
	return uru
}

// SetEntID sets the "ent_id" field.
func (uru *UserRewardUpdate) SetEntID(u uuid.UUID) *UserRewardUpdate {
	uru.mutation.SetEntID(u)
	return uru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableEntID(u *uuid.UUID) *UserRewardUpdate {
	if u != nil {
		uru.SetEntID(*u)
	}
	return uru
}

// SetAppID sets the "app_id" field.
func (uru *UserRewardUpdate) SetAppID(u uuid.UUID) *UserRewardUpdate {
	uru.mutation.SetAppID(u)
	return uru
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableAppID(u *uuid.UUID) *UserRewardUpdate {
	if u != nil {
		uru.SetAppID(*u)
	}
	return uru
}

// ClearAppID clears the value of the "app_id" field.
func (uru *UserRewardUpdate) ClearAppID() *UserRewardUpdate {
	uru.mutation.ClearAppID()
	return uru
}

// SetUserID sets the "user_id" field.
func (uru *UserRewardUpdate) SetUserID(u uuid.UUID) *UserRewardUpdate {
	uru.mutation.SetUserID(u)
	return uru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableUserID(u *uuid.UUID) *UserRewardUpdate {
	if u != nil {
		uru.SetUserID(*u)
	}
	return uru
}

// ClearUserID clears the value of the "user_id" field.
func (uru *UserRewardUpdate) ClearUserID() *UserRewardUpdate {
	uru.mutation.ClearUserID()
	return uru
}

// SetActionCredits sets the "action_credits" field.
func (uru *UserRewardUpdate) SetActionCredits(d decimal.Decimal) *UserRewardUpdate {
	uru.mutation.SetActionCredits(d)
	return uru
}

// SetNillableActionCredits sets the "action_credits" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableActionCredits(d *decimal.Decimal) *UserRewardUpdate {
	if d != nil {
		uru.SetActionCredits(*d)
	}
	return uru
}

// ClearActionCredits clears the value of the "action_credits" field.
func (uru *UserRewardUpdate) ClearActionCredits() *UserRewardUpdate {
	uru.mutation.ClearActionCredits()
	return uru
}

// SetCouponAmount sets the "coupon_amount" field.
func (uru *UserRewardUpdate) SetCouponAmount(d decimal.Decimal) *UserRewardUpdate {
	uru.mutation.SetCouponAmount(d)
	return uru
}

// SetNillableCouponAmount sets the "coupon_amount" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableCouponAmount(d *decimal.Decimal) *UserRewardUpdate {
	if d != nil {
		uru.SetCouponAmount(*d)
	}
	return uru
}

// ClearCouponAmount clears the value of the "coupon_amount" field.
func (uru *UserRewardUpdate) ClearCouponAmount() *UserRewardUpdate {
	uru.mutation.ClearCouponAmount()
	return uru
}

// SetCouponCashableAmount sets the "coupon_cashable_amount" field.
func (uru *UserRewardUpdate) SetCouponCashableAmount(d decimal.Decimal) *UserRewardUpdate {
	uru.mutation.SetCouponCashableAmount(d)
	return uru
}

// SetNillableCouponCashableAmount sets the "coupon_cashable_amount" field if the given value is not nil.
func (uru *UserRewardUpdate) SetNillableCouponCashableAmount(d *decimal.Decimal) *UserRewardUpdate {
	if d != nil {
		uru.SetCouponCashableAmount(*d)
	}
	return uru
}

// ClearCouponCashableAmount clears the value of the "coupon_cashable_amount" field.
func (uru *UserRewardUpdate) ClearCouponCashableAmount() *UserRewardUpdate {
	uru.mutation.ClearCouponCashableAmount()
	return uru
}

// Mutation returns the UserRewardMutation object of the builder.
func (uru *UserRewardUpdate) Mutation() *UserRewardMutation {
	return uru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRewardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := uru.defaults(); err != nil {
		return 0, err
	}
	if len(uru.hooks) == 0 {
		affected, err = uru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uru.mutation = mutation
			affected, err = uru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uru.hooks) - 1; i >= 0; i-- {
			if uru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRewardUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRewardUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRewardUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uru *UserRewardUpdate) defaults() error {
	if _, ok := uru.mutation.UpdatedAt(); !ok {
		if userreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userreward.UpdateDefaultUpdatedAt()
		uru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uru *UserRewardUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRewardUpdate {
	uru.modifiers = append(uru.modifiers, modifiers...)
	return uru
}

func (uru *UserRewardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userreward.Table,
			Columns: userreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: userreward.FieldID,
			},
		},
	}
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldCreatedAt,
		})
	}
	if value, ok := uru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldCreatedAt,
		})
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldUpdatedAt,
		})
	}
	if value, ok := uru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldUpdatedAt,
		})
	}
	if value, ok := uru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldDeletedAt,
		})
	}
	if value, ok := uru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldDeletedAt,
		})
	}
	if value, ok := uru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldEntID,
		})
	}
	if value, ok := uru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldAppID,
		})
	}
	if uru.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: userreward.FieldAppID,
		})
	}
	if value, ok := uru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldUserID,
		})
	}
	if uru.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: userreward.FieldUserID,
		})
	}
	if value, ok := uru.mutation.ActionCredits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldActionCredits,
		})
	}
	if uru.mutation.ActionCreditsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldActionCredits,
		})
	}
	if value, ok := uru.mutation.CouponAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldCouponAmount,
		})
	}
	if uru.mutation.CouponAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldCouponAmount,
		})
	}
	if value, ok := uru.mutation.CouponCashableAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldCouponCashableAmount,
		})
	}
	if uru.mutation.CouponCashableAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldCouponCashableAmount,
		})
	}
	_spec.Modifiers = uru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserRewardUpdateOne is the builder for updating a single UserReward entity.
type UserRewardUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (uruo *UserRewardUpdateOne) SetCreatedAt(u uint32) *UserRewardUpdateOne {
	uruo.mutation.ResetCreatedAt()
	uruo.mutation.SetCreatedAt(u)
	return uruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableCreatedAt(u *uint32) *UserRewardUpdateOne {
	if u != nil {
		uruo.SetCreatedAt(*u)
	}
	return uruo
}

// AddCreatedAt adds u to the "created_at" field.
func (uruo *UserRewardUpdateOne) AddCreatedAt(u int32) *UserRewardUpdateOne {
	uruo.mutation.AddCreatedAt(u)
	return uruo
}

// SetUpdatedAt sets the "updated_at" field.
func (uruo *UserRewardUpdateOne) SetUpdatedAt(u uint32) *UserRewardUpdateOne {
	uruo.mutation.ResetUpdatedAt()
	uruo.mutation.SetUpdatedAt(u)
	return uruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (uruo *UserRewardUpdateOne) AddUpdatedAt(u int32) *UserRewardUpdateOne {
	uruo.mutation.AddUpdatedAt(u)
	return uruo
}

// SetDeletedAt sets the "deleted_at" field.
func (uruo *UserRewardUpdateOne) SetDeletedAt(u uint32) *UserRewardUpdateOne {
	uruo.mutation.ResetDeletedAt()
	uruo.mutation.SetDeletedAt(u)
	return uruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableDeletedAt(u *uint32) *UserRewardUpdateOne {
	if u != nil {
		uruo.SetDeletedAt(*u)
	}
	return uruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (uruo *UserRewardUpdateOne) AddDeletedAt(u int32) *UserRewardUpdateOne {
	uruo.mutation.AddDeletedAt(u)
	return uruo
}

// SetEntID sets the "ent_id" field.
func (uruo *UserRewardUpdateOne) SetEntID(u uuid.UUID) *UserRewardUpdateOne {
	uruo.mutation.SetEntID(u)
	return uruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableEntID(u *uuid.UUID) *UserRewardUpdateOne {
	if u != nil {
		uruo.SetEntID(*u)
	}
	return uruo
}

// SetAppID sets the "app_id" field.
func (uruo *UserRewardUpdateOne) SetAppID(u uuid.UUID) *UserRewardUpdateOne {
	uruo.mutation.SetAppID(u)
	return uruo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableAppID(u *uuid.UUID) *UserRewardUpdateOne {
	if u != nil {
		uruo.SetAppID(*u)
	}
	return uruo
}

// ClearAppID clears the value of the "app_id" field.
func (uruo *UserRewardUpdateOne) ClearAppID() *UserRewardUpdateOne {
	uruo.mutation.ClearAppID()
	return uruo
}

// SetUserID sets the "user_id" field.
func (uruo *UserRewardUpdateOne) SetUserID(u uuid.UUID) *UserRewardUpdateOne {
	uruo.mutation.SetUserID(u)
	return uruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableUserID(u *uuid.UUID) *UserRewardUpdateOne {
	if u != nil {
		uruo.SetUserID(*u)
	}
	return uruo
}

// ClearUserID clears the value of the "user_id" field.
func (uruo *UserRewardUpdateOne) ClearUserID() *UserRewardUpdateOne {
	uruo.mutation.ClearUserID()
	return uruo
}

// SetActionCredits sets the "action_credits" field.
func (uruo *UserRewardUpdateOne) SetActionCredits(d decimal.Decimal) *UserRewardUpdateOne {
	uruo.mutation.SetActionCredits(d)
	return uruo
}

// SetNillableActionCredits sets the "action_credits" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableActionCredits(d *decimal.Decimal) *UserRewardUpdateOne {
	if d != nil {
		uruo.SetActionCredits(*d)
	}
	return uruo
}

// ClearActionCredits clears the value of the "action_credits" field.
func (uruo *UserRewardUpdateOne) ClearActionCredits() *UserRewardUpdateOne {
	uruo.mutation.ClearActionCredits()
	return uruo
}

// SetCouponAmount sets the "coupon_amount" field.
func (uruo *UserRewardUpdateOne) SetCouponAmount(d decimal.Decimal) *UserRewardUpdateOne {
	uruo.mutation.SetCouponAmount(d)
	return uruo
}

// SetNillableCouponAmount sets the "coupon_amount" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableCouponAmount(d *decimal.Decimal) *UserRewardUpdateOne {
	if d != nil {
		uruo.SetCouponAmount(*d)
	}
	return uruo
}

// ClearCouponAmount clears the value of the "coupon_amount" field.
func (uruo *UserRewardUpdateOne) ClearCouponAmount() *UserRewardUpdateOne {
	uruo.mutation.ClearCouponAmount()
	return uruo
}

// SetCouponCashableAmount sets the "coupon_cashable_amount" field.
func (uruo *UserRewardUpdateOne) SetCouponCashableAmount(d decimal.Decimal) *UserRewardUpdateOne {
	uruo.mutation.SetCouponCashableAmount(d)
	return uruo
}

// SetNillableCouponCashableAmount sets the "coupon_cashable_amount" field if the given value is not nil.
func (uruo *UserRewardUpdateOne) SetNillableCouponCashableAmount(d *decimal.Decimal) *UserRewardUpdateOne {
	if d != nil {
		uruo.SetCouponCashableAmount(*d)
	}
	return uruo
}

// ClearCouponCashableAmount clears the value of the "coupon_cashable_amount" field.
func (uruo *UserRewardUpdateOne) ClearCouponCashableAmount() *UserRewardUpdateOne {
	uruo.mutation.ClearCouponCashableAmount()
	return uruo
}

// Mutation returns the UserRewardMutation object of the builder.
func (uruo *UserRewardUpdateOne) Mutation() *UserRewardMutation {
	return uruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserRewardUpdateOne) Select(field string, fields ...string) *UserRewardUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserReward entity.
func (uruo *UserRewardUpdateOne) Save(ctx context.Context) (*UserReward, error) {
	var (
		err  error
		node *UserReward
	)
	if err := uruo.defaults(); err != nil {
		return nil, err
	}
	if len(uruo.hooks) == 0 {
		node, err = uruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uruo.mutation = mutation
			node, err = uruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uruo.hooks) - 1; i >= 0; i-- {
			if uruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserReward)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserRewardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRewardUpdateOne) SaveX(ctx context.Context) *UserReward {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRewardUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRewardUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uruo *UserRewardUpdateOne) defaults() error {
	if _, ok := uruo.mutation.UpdatedAt(); !ok {
		if userreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userreward.UpdateDefaultUpdatedAt()
		uruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uruo *UserRewardUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRewardUpdateOne {
	uruo.modifiers = append(uruo.modifiers, modifiers...)
	return uruo
}

func (uruo *UserRewardUpdateOne) sqlSave(ctx context.Context) (_node *UserReward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userreward.Table,
			Columns: userreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: userreward.FieldID,
			},
		},
	}
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserReward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userreward.FieldID)
		for _, f := range fields {
			if !userreward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldCreatedAt,
		})
	}
	if value, ok := uruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldCreatedAt,
		})
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldUpdatedAt,
		})
	}
	if value, ok := uruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldUpdatedAt,
		})
	}
	if value, ok := uruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldDeletedAt,
		})
	}
	if value, ok := uruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userreward.FieldDeletedAt,
		})
	}
	if value, ok := uruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldEntID,
		})
	}
	if value, ok := uruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldAppID,
		})
	}
	if uruo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: userreward.FieldAppID,
		})
	}
	if value, ok := uruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userreward.FieldUserID,
		})
	}
	if uruo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: userreward.FieldUserID,
		})
	}
	if value, ok := uruo.mutation.ActionCredits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldActionCredits,
		})
	}
	if uruo.mutation.ActionCreditsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldActionCredits,
		})
	}
	if value, ok := uruo.mutation.CouponAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldCouponAmount,
		})
	}
	if uruo.mutation.CouponAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldCouponAmount,
		})
	}
	if value, ok := uruo.mutation.CouponCashableAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: userreward.FieldCouponCashableAmount,
		})
	}
	if uruo.mutation.CouponCashableAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: userreward.FieldCouponCashableAmount,
		})
	}
	_spec.Modifiers = uruo.modifiers
	_node = &UserReward{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
