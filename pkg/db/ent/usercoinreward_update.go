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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/usercoinreward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UserCoinRewardUpdate is the builder for updating UserCoinReward entities.
type UserCoinRewardUpdate struct {
	config
	hooks     []Hook
	mutation  *UserCoinRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserCoinRewardUpdate builder.
func (ucru *UserCoinRewardUpdate) Where(ps ...predicate.UserCoinReward) *UserCoinRewardUpdate {
	ucru.mutation.Where(ps...)
	return ucru
}

// SetCreatedAt sets the "created_at" field.
func (ucru *UserCoinRewardUpdate) SetCreatedAt(u uint32) *UserCoinRewardUpdate {
	ucru.mutation.ResetCreatedAt()
	ucru.mutation.SetCreatedAt(u)
	return ucru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableCreatedAt(u *uint32) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetCreatedAt(*u)
	}
	return ucru
}

// AddCreatedAt adds u to the "created_at" field.
func (ucru *UserCoinRewardUpdate) AddCreatedAt(u int32) *UserCoinRewardUpdate {
	ucru.mutation.AddCreatedAt(u)
	return ucru
}

// SetUpdatedAt sets the "updated_at" field.
func (ucru *UserCoinRewardUpdate) SetUpdatedAt(u uint32) *UserCoinRewardUpdate {
	ucru.mutation.ResetUpdatedAt()
	ucru.mutation.SetUpdatedAt(u)
	return ucru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ucru *UserCoinRewardUpdate) AddUpdatedAt(u int32) *UserCoinRewardUpdate {
	ucru.mutation.AddUpdatedAt(u)
	return ucru
}

// SetDeletedAt sets the "deleted_at" field.
func (ucru *UserCoinRewardUpdate) SetDeletedAt(u uint32) *UserCoinRewardUpdate {
	ucru.mutation.ResetDeletedAt()
	ucru.mutation.SetDeletedAt(u)
	return ucru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableDeletedAt(u *uint32) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetDeletedAt(*u)
	}
	return ucru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ucru *UserCoinRewardUpdate) AddDeletedAt(u int32) *UserCoinRewardUpdate {
	ucru.mutation.AddDeletedAt(u)
	return ucru
}

// SetEntID sets the "ent_id" field.
func (ucru *UserCoinRewardUpdate) SetEntID(u uuid.UUID) *UserCoinRewardUpdate {
	ucru.mutation.SetEntID(u)
	return ucru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableEntID(u *uuid.UUID) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetEntID(*u)
	}
	return ucru
}

// SetAppID sets the "app_id" field.
func (ucru *UserCoinRewardUpdate) SetAppID(u uuid.UUID) *UserCoinRewardUpdate {
	ucru.mutation.SetAppID(u)
	return ucru
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableAppID(u *uuid.UUID) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetAppID(*u)
	}
	return ucru
}

// ClearAppID clears the value of the "app_id" field.
func (ucru *UserCoinRewardUpdate) ClearAppID() *UserCoinRewardUpdate {
	ucru.mutation.ClearAppID()
	return ucru
}

// SetUserID sets the "user_id" field.
func (ucru *UserCoinRewardUpdate) SetUserID(u uuid.UUID) *UserCoinRewardUpdate {
	ucru.mutation.SetUserID(u)
	return ucru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableUserID(u *uuid.UUID) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetUserID(*u)
	}
	return ucru
}

// ClearUserID clears the value of the "user_id" field.
func (ucru *UserCoinRewardUpdate) ClearUserID() *UserCoinRewardUpdate {
	ucru.mutation.ClearUserID()
	return ucru
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ucru *UserCoinRewardUpdate) SetCoinTypeID(u uuid.UUID) *UserCoinRewardUpdate {
	ucru.mutation.SetCoinTypeID(u)
	return ucru
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableCoinTypeID(u *uuid.UUID) *UserCoinRewardUpdate {
	if u != nil {
		ucru.SetCoinTypeID(*u)
	}
	return ucru
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ucru *UserCoinRewardUpdate) ClearCoinTypeID() *UserCoinRewardUpdate {
	ucru.mutation.ClearCoinTypeID()
	return ucru
}

// SetCoinRewards sets the "coin_rewards" field.
func (ucru *UserCoinRewardUpdate) SetCoinRewards(d decimal.Decimal) *UserCoinRewardUpdate {
	ucru.mutation.SetCoinRewards(d)
	return ucru
}

// SetNillableCoinRewards sets the "coin_rewards" field if the given value is not nil.
func (ucru *UserCoinRewardUpdate) SetNillableCoinRewards(d *decimal.Decimal) *UserCoinRewardUpdate {
	if d != nil {
		ucru.SetCoinRewards(*d)
	}
	return ucru
}

// ClearCoinRewards clears the value of the "coin_rewards" field.
func (ucru *UserCoinRewardUpdate) ClearCoinRewards() *UserCoinRewardUpdate {
	ucru.mutation.ClearCoinRewards()
	return ucru
}

// Mutation returns the UserCoinRewardMutation object of the builder.
func (ucru *UserCoinRewardUpdate) Mutation() *UserCoinRewardMutation {
	return ucru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucru *UserCoinRewardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ucru.defaults(); err != nil {
		return 0, err
	}
	if len(ucru.hooks) == 0 {
		affected, err = ucru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoinRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucru.mutation = mutation
			affected, err = ucru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucru.hooks) - 1; i >= 0; i-- {
			if ucru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucru *UserCoinRewardUpdate) SaveX(ctx context.Context) int {
	affected, err := ucru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucru *UserCoinRewardUpdate) Exec(ctx context.Context) error {
	_, err := ucru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucru *UserCoinRewardUpdate) ExecX(ctx context.Context) {
	if err := ucru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucru *UserCoinRewardUpdate) defaults() error {
	if _, ok := ucru.mutation.UpdatedAt(); !ok {
		if usercoinreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usercoinreward.UpdateDefaultUpdatedAt()
		ucru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucru *UserCoinRewardUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCoinRewardUpdate {
	ucru.modifiers = append(ucru.modifiers, modifiers...)
	return ucru
}

func (ucru *UserCoinRewardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoinreward.Table,
			Columns: usercoinreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercoinreward.FieldID,
			},
		},
	}
	if ps := ucru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldCreatedAt,
		})
	}
	if value, ok := ucru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldCreatedAt,
		})
	}
	if value, ok := ucru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := ucru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := ucru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldDeletedAt,
		})
	}
	if value, ok := ucru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldDeletedAt,
		})
	}
	if value, ok := ucru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldEntID,
		})
	}
	if value, ok := ucru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldAppID,
		})
	}
	if ucru.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldAppID,
		})
	}
	if value, ok := ucru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldUserID,
		})
	}
	if ucru.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldUserID,
		})
	}
	if value, ok := ucru.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldCoinTypeID,
		})
	}
	if ucru.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldCoinTypeID,
		})
	}
	if value, ok := ucru.mutation.CoinRewards(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: usercoinreward.FieldCoinRewards,
		})
	}
	if ucru.mutation.CoinRewardsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: usercoinreward.FieldCoinRewards,
		})
	}
	_spec.Modifiers = ucru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ucru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoinreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserCoinRewardUpdateOne is the builder for updating a single UserCoinReward entity.
type UserCoinRewardUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserCoinRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ucruo *UserCoinRewardUpdateOne) SetCreatedAt(u uint32) *UserCoinRewardUpdateOne {
	ucruo.mutation.ResetCreatedAt()
	ucruo.mutation.SetCreatedAt(u)
	return ucruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableCreatedAt(u *uint32) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetCreatedAt(*u)
	}
	return ucruo
}

// AddCreatedAt adds u to the "created_at" field.
func (ucruo *UserCoinRewardUpdateOne) AddCreatedAt(u int32) *UserCoinRewardUpdateOne {
	ucruo.mutation.AddCreatedAt(u)
	return ucruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ucruo *UserCoinRewardUpdateOne) SetUpdatedAt(u uint32) *UserCoinRewardUpdateOne {
	ucruo.mutation.ResetUpdatedAt()
	ucruo.mutation.SetUpdatedAt(u)
	return ucruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ucruo *UserCoinRewardUpdateOne) AddUpdatedAt(u int32) *UserCoinRewardUpdateOne {
	ucruo.mutation.AddUpdatedAt(u)
	return ucruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ucruo *UserCoinRewardUpdateOne) SetDeletedAt(u uint32) *UserCoinRewardUpdateOne {
	ucruo.mutation.ResetDeletedAt()
	ucruo.mutation.SetDeletedAt(u)
	return ucruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableDeletedAt(u *uint32) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetDeletedAt(*u)
	}
	return ucruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ucruo *UserCoinRewardUpdateOne) AddDeletedAt(u int32) *UserCoinRewardUpdateOne {
	ucruo.mutation.AddDeletedAt(u)
	return ucruo
}

// SetEntID sets the "ent_id" field.
func (ucruo *UserCoinRewardUpdateOne) SetEntID(u uuid.UUID) *UserCoinRewardUpdateOne {
	ucruo.mutation.SetEntID(u)
	return ucruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableEntID(u *uuid.UUID) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetEntID(*u)
	}
	return ucruo
}

// SetAppID sets the "app_id" field.
func (ucruo *UserCoinRewardUpdateOne) SetAppID(u uuid.UUID) *UserCoinRewardUpdateOne {
	ucruo.mutation.SetAppID(u)
	return ucruo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableAppID(u *uuid.UUID) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetAppID(*u)
	}
	return ucruo
}

// ClearAppID clears the value of the "app_id" field.
func (ucruo *UserCoinRewardUpdateOne) ClearAppID() *UserCoinRewardUpdateOne {
	ucruo.mutation.ClearAppID()
	return ucruo
}

// SetUserID sets the "user_id" field.
func (ucruo *UserCoinRewardUpdateOne) SetUserID(u uuid.UUID) *UserCoinRewardUpdateOne {
	ucruo.mutation.SetUserID(u)
	return ucruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableUserID(u *uuid.UUID) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetUserID(*u)
	}
	return ucruo
}

// ClearUserID clears the value of the "user_id" field.
func (ucruo *UserCoinRewardUpdateOne) ClearUserID() *UserCoinRewardUpdateOne {
	ucruo.mutation.ClearUserID()
	return ucruo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ucruo *UserCoinRewardUpdateOne) SetCoinTypeID(u uuid.UUID) *UserCoinRewardUpdateOne {
	ucruo.mutation.SetCoinTypeID(u)
	return ucruo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *UserCoinRewardUpdateOne {
	if u != nil {
		ucruo.SetCoinTypeID(*u)
	}
	return ucruo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ucruo *UserCoinRewardUpdateOne) ClearCoinTypeID() *UserCoinRewardUpdateOne {
	ucruo.mutation.ClearCoinTypeID()
	return ucruo
}

// SetCoinRewards sets the "coin_rewards" field.
func (ucruo *UserCoinRewardUpdateOne) SetCoinRewards(d decimal.Decimal) *UserCoinRewardUpdateOne {
	ucruo.mutation.SetCoinRewards(d)
	return ucruo
}

// SetNillableCoinRewards sets the "coin_rewards" field if the given value is not nil.
func (ucruo *UserCoinRewardUpdateOne) SetNillableCoinRewards(d *decimal.Decimal) *UserCoinRewardUpdateOne {
	if d != nil {
		ucruo.SetCoinRewards(*d)
	}
	return ucruo
}

// ClearCoinRewards clears the value of the "coin_rewards" field.
func (ucruo *UserCoinRewardUpdateOne) ClearCoinRewards() *UserCoinRewardUpdateOne {
	ucruo.mutation.ClearCoinRewards()
	return ucruo
}

// Mutation returns the UserCoinRewardMutation object of the builder.
func (ucruo *UserCoinRewardUpdateOne) Mutation() *UserCoinRewardMutation {
	return ucruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucruo *UserCoinRewardUpdateOne) Select(field string, fields ...string) *UserCoinRewardUpdateOne {
	ucruo.fields = append([]string{field}, fields...)
	return ucruo
}

// Save executes the query and returns the updated UserCoinReward entity.
func (ucruo *UserCoinRewardUpdateOne) Save(ctx context.Context) (*UserCoinReward, error) {
	var (
		err  error
		node *UserCoinReward
	)
	if err := ucruo.defaults(); err != nil {
		return nil, err
	}
	if len(ucruo.hooks) == 0 {
		node, err = ucruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoinRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucruo.mutation = mutation
			node, err = ucruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucruo.hooks) - 1; i >= 0; i-- {
			if ucruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ucruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserCoinReward)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserCoinRewardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucruo *UserCoinRewardUpdateOne) SaveX(ctx context.Context) *UserCoinReward {
	node, err := ucruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucruo *UserCoinRewardUpdateOne) Exec(ctx context.Context) error {
	_, err := ucruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucruo *UserCoinRewardUpdateOne) ExecX(ctx context.Context) {
	if err := ucruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucruo *UserCoinRewardUpdateOne) defaults() error {
	if _, ok := ucruo.mutation.UpdatedAt(); !ok {
		if usercoinreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usercoinreward.UpdateDefaultUpdatedAt()
		ucruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucruo *UserCoinRewardUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCoinRewardUpdateOne {
	ucruo.modifiers = append(ucruo.modifiers, modifiers...)
	return ucruo
}

func (ucruo *UserCoinRewardUpdateOne) sqlSave(ctx context.Context) (_node *UserCoinReward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoinreward.Table,
			Columns: usercoinreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercoinreward.FieldID,
			},
		},
	}
	id, ok := ucruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCoinReward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercoinreward.FieldID)
		for _, f := range fields {
			if !usercoinreward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercoinreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldCreatedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldCreatedAt,
		})
	}
	if value, ok := ucruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := ucruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldDeletedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldDeletedAt,
		})
	}
	if value, ok := ucruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldEntID,
		})
	}
	if value, ok := ucruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldAppID,
		})
	}
	if ucruo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldAppID,
		})
	}
	if value, ok := ucruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldUserID,
		})
	}
	if ucruo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldUserID,
		})
	}
	if value, ok := ucruo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldCoinTypeID,
		})
	}
	if ucruo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercoinreward.FieldCoinTypeID,
		})
	}
	if value, ok := ucruo.mutation.CoinRewards(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: usercoinreward.FieldCoinRewards,
		})
	}
	if ucruo.mutation.CoinRewardsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: usercoinreward.FieldCoinRewards,
		})
	}
	_spec.Modifiers = ucruo.modifiers
	_node = &UserCoinReward{config: ucruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoinreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}