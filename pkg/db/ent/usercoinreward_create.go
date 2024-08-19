// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/usercoinreward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UserCoinRewardCreate is the builder for creating a UserCoinReward entity.
type UserCoinRewardCreate struct {
	config
	mutation *UserCoinRewardMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ucrc *UserCoinRewardCreate) SetCreatedAt(u uint32) *UserCoinRewardCreate {
	ucrc.mutation.SetCreatedAt(u)
	return ucrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableCreatedAt(u *uint32) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetCreatedAt(*u)
	}
	return ucrc
}

// SetUpdatedAt sets the "updated_at" field.
func (ucrc *UserCoinRewardCreate) SetUpdatedAt(u uint32) *UserCoinRewardCreate {
	ucrc.mutation.SetUpdatedAt(u)
	return ucrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableUpdatedAt(u *uint32) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetUpdatedAt(*u)
	}
	return ucrc
}

// SetDeletedAt sets the "deleted_at" field.
func (ucrc *UserCoinRewardCreate) SetDeletedAt(u uint32) *UserCoinRewardCreate {
	ucrc.mutation.SetDeletedAt(u)
	return ucrc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableDeletedAt(u *uint32) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetDeletedAt(*u)
	}
	return ucrc
}

// SetEntID sets the "ent_id" field.
func (ucrc *UserCoinRewardCreate) SetEntID(u uuid.UUID) *UserCoinRewardCreate {
	ucrc.mutation.SetEntID(u)
	return ucrc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableEntID(u *uuid.UUID) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetEntID(*u)
	}
	return ucrc
}

// SetAppID sets the "app_id" field.
func (ucrc *UserCoinRewardCreate) SetAppID(u uuid.UUID) *UserCoinRewardCreate {
	ucrc.mutation.SetAppID(u)
	return ucrc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableAppID(u *uuid.UUID) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetAppID(*u)
	}
	return ucrc
}

// SetUserID sets the "user_id" field.
func (ucrc *UserCoinRewardCreate) SetUserID(u uuid.UUID) *UserCoinRewardCreate {
	ucrc.mutation.SetUserID(u)
	return ucrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableUserID(u *uuid.UUID) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetUserID(*u)
	}
	return ucrc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ucrc *UserCoinRewardCreate) SetCoinTypeID(u uuid.UUID) *UserCoinRewardCreate {
	ucrc.mutation.SetCoinTypeID(u)
	return ucrc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableCoinTypeID(u *uuid.UUID) *UserCoinRewardCreate {
	if u != nil {
		ucrc.SetCoinTypeID(*u)
	}
	return ucrc
}

// SetCoinRewards sets the "coin_rewards" field.
func (ucrc *UserCoinRewardCreate) SetCoinRewards(d decimal.Decimal) *UserCoinRewardCreate {
	ucrc.mutation.SetCoinRewards(d)
	return ucrc
}

// SetNillableCoinRewards sets the "coin_rewards" field if the given value is not nil.
func (ucrc *UserCoinRewardCreate) SetNillableCoinRewards(d *decimal.Decimal) *UserCoinRewardCreate {
	if d != nil {
		ucrc.SetCoinRewards(*d)
	}
	return ucrc
}

// SetID sets the "id" field.
func (ucrc *UserCoinRewardCreate) SetID(u uint32) *UserCoinRewardCreate {
	ucrc.mutation.SetID(u)
	return ucrc
}

// Mutation returns the UserCoinRewardMutation object of the builder.
func (ucrc *UserCoinRewardCreate) Mutation() *UserCoinRewardMutation {
	return ucrc.mutation
}

// Save creates the UserCoinReward in the database.
func (ucrc *UserCoinRewardCreate) Save(ctx context.Context) (*UserCoinReward, error) {
	var (
		err  error
		node *UserCoinReward
	)
	if err := ucrc.defaults(); err != nil {
		return nil, err
	}
	if len(ucrc.hooks) == 0 {
		if err = ucrc.check(); err != nil {
			return nil, err
		}
		node, err = ucrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoinRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucrc.check(); err != nil {
				return nil, err
			}
			ucrc.mutation = mutation
			if node, err = ucrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucrc.hooks) - 1; i >= 0; i-- {
			if ucrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ucrc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ucrc *UserCoinRewardCreate) SaveX(ctx context.Context) *UserCoinReward {
	v, err := ucrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucrc *UserCoinRewardCreate) Exec(ctx context.Context) error {
	_, err := ucrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucrc *UserCoinRewardCreate) ExecX(ctx context.Context) {
	if err := ucrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucrc *UserCoinRewardCreate) defaults() error {
	if _, ok := ucrc.mutation.CreatedAt(); !ok {
		if usercoinreward.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultCreatedAt()
		ucrc.mutation.SetCreatedAt(v)
	}
	if _, ok := ucrc.mutation.UpdatedAt(); !ok {
		if usercoinreward.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultUpdatedAt()
		ucrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ucrc.mutation.DeletedAt(); !ok {
		if usercoinreward.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultDeletedAt()
		ucrc.mutation.SetDeletedAt(v)
	}
	if _, ok := ucrc.mutation.EntID(); !ok {
		if usercoinreward.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultEntID()
		ucrc.mutation.SetEntID(v)
	}
	if _, ok := ucrc.mutation.AppID(); !ok {
		if usercoinreward.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultAppID()
		ucrc.mutation.SetAppID(v)
	}
	if _, ok := ucrc.mutation.UserID(); !ok {
		if usercoinreward.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultUserID()
		ucrc.mutation.SetUserID(v)
	}
	if _, ok := ucrc.mutation.CoinTypeID(); !ok {
		if usercoinreward.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized usercoinreward.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := usercoinreward.DefaultCoinTypeID()
		ucrc.mutation.SetCoinTypeID(v)
	}
	if _, ok := ucrc.mutation.CoinRewards(); !ok {
		v := usercoinreward.DefaultCoinRewards
		ucrc.mutation.SetCoinRewards(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ucrc *UserCoinRewardCreate) check() error {
	if _, ok := ucrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserCoinReward.created_at"`)}
	}
	if _, ok := ucrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserCoinReward.updated_at"`)}
	}
	if _, ok := ucrc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "UserCoinReward.deleted_at"`)}
	}
	if _, ok := ucrc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "UserCoinReward.ent_id"`)}
	}
	return nil
}

func (ucrc *UserCoinRewardCreate) sqlSave(ctx context.Context) (*UserCoinReward, error) {
	_node, _spec := ucrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (ucrc *UserCoinRewardCreate) createSpec() (*UserCoinReward, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCoinReward{config: ucrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usercoinreward.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercoinreward.FieldID,
			},
		}
	)
	_spec.OnConflict = ucrc.conflict
	if id, ok := ucrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ucrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ucrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ucrc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercoinreward.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ucrc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := ucrc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ucrc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ucrc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercoinreward.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := ucrc.mutation.CoinRewards(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: usercoinreward.FieldCoinRewards,
		})
		_node.CoinRewards = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCoinReward.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCoinRewardUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ucrc *UserCoinRewardCreate) OnConflict(opts ...sql.ConflictOption) *UserCoinRewardUpsertOne {
	ucrc.conflict = opts
	return &UserCoinRewardUpsertOne{
		create: ucrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCoinReward.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ucrc *UserCoinRewardCreate) OnConflictColumns(columns ...string) *UserCoinRewardUpsertOne {
	ucrc.conflict = append(ucrc.conflict, sql.ConflictColumns(columns...))
	return &UserCoinRewardUpsertOne{
		create: ucrc,
	}
}

type (
	// UserCoinRewardUpsertOne is the builder for "upsert"-ing
	//  one UserCoinReward node.
	UserCoinRewardUpsertOne struct {
		create *UserCoinRewardCreate
	}

	// UserCoinRewardUpsert is the "OnConflict" setter.
	UserCoinRewardUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserCoinRewardUpsert) SetCreatedAt(v uint32) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateCreatedAt() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserCoinRewardUpsert) AddCreatedAt(v uint32) *UserCoinRewardUpsert {
	u.Add(usercoinreward.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserCoinRewardUpsert) SetUpdatedAt(v uint32) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateUpdatedAt() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserCoinRewardUpsert) AddUpdatedAt(v uint32) *UserCoinRewardUpsert {
	u.Add(usercoinreward.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserCoinRewardUpsert) SetDeletedAt(v uint32) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateDeletedAt() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserCoinRewardUpsert) AddDeletedAt(v uint32) *UserCoinRewardUpsert {
	u.Add(usercoinreward.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *UserCoinRewardUpsert) SetEntID(v uuid.UUID) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateEntID() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserCoinRewardUpsert) SetAppID(v uuid.UUID) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateAppID() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserCoinRewardUpsert) ClearAppID() *UserCoinRewardUpsert {
	u.SetNull(usercoinreward.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserCoinRewardUpsert) SetUserID(v uuid.UUID) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateUserID() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserCoinRewardUpsert) ClearUserID() *UserCoinRewardUpsert {
	u.SetNull(usercoinreward.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserCoinRewardUpsert) SetCoinTypeID(v uuid.UUID) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateCoinTypeID() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserCoinRewardUpsert) ClearCoinTypeID() *UserCoinRewardUpsert {
	u.SetNull(usercoinreward.FieldCoinTypeID)
	return u
}

// SetCoinRewards sets the "coin_rewards" field.
func (u *UserCoinRewardUpsert) SetCoinRewards(v decimal.Decimal) *UserCoinRewardUpsert {
	u.Set(usercoinreward.FieldCoinRewards, v)
	return u
}

// UpdateCoinRewards sets the "coin_rewards" field to the value that was provided on create.
func (u *UserCoinRewardUpsert) UpdateCoinRewards() *UserCoinRewardUpsert {
	u.SetExcluded(usercoinreward.FieldCoinRewards)
	return u
}

// ClearCoinRewards clears the value of the "coin_rewards" field.
func (u *UserCoinRewardUpsert) ClearCoinRewards() *UserCoinRewardUpsert {
	u.SetNull(usercoinreward.FieldCoinRewards)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserCoinReward.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usercoinreward.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserCoinRewardUpsertOne) UpdateNewValues() *UserCoinRewardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(usercoinreward.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserCoinReward.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserCoinRewardUpsertOne) Ignore() *UserCoinRewardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCoinRewardUpsertOne) DoNothing() *UserCoinRewardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCoinRewardCreate.OnConflict
// documentation for more info.
func (u *UserCoinRewardUpsertOne) Update(set func(*UserCoinRewardUpsert)) *UserCoinRewardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCoinRewardUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserCoinRewardUpsertOne) SetCreatedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserCoinRewardUpsertOne) AddCreatedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateCreatedAt() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserCoinRewardUpsertOne) SetUpdatedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserCoinRewardUpsertOne) AddUpdatedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateUpdatedAt() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserCoinRewardUpsertOne) SetDeletedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserCoinRewardUpsertOne) AddDeletedAt(v uint32) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateDeletedAt() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *UserCoinRewardUpsertOne) SetEntID(v uuid.UUID) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateEntID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserCoinRewardUpsertOne) SetAppID(v uuid.UUID) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateAppID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserCoinRewardUpsertOne) ClearAppID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserCoinRewardUpsertOne) SetUserID(v uuid.UUID) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateUserID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserCoinRewardUpsertOne) ClearUserID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserCoinRewardUpsertOne) SetCoinTypeID(v uuid.UUID) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateCoinTypeID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserCoinRewardUpsertOne) ClearCoinTypeID() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetCoinRewards sets the "coin_rewards" field.
func (u *UserCoinRewardUpsertOne) SetCoinRewards(v decimal.Decimal) *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCoinRewards(v)
	})
}

// UpdateCoinRewards sets the "coin_rewards" field to the value that was provided on create.
func (u *UserCoinRewardUpsertOne) UpdateCoinRewards() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCoinRewards()
	})
}

// ClearCoinRewards clears the value of the "coin_rewards" field.
func (u *UserCoinRewardUpsertOne) ClearCoinRewards() *UserCoinRewardUpsertOne {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearCoinRewards()
	})
}

// Exec executes the query.
func (u *UserCoinRewardUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCoinRewardCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCoinRewardUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserCoinRewardUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserCoinRewardUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCoinRewardCreateBulk is the builder for creating many UserCoinReward entities in bulk.
type UserCoinRewardCreateBulk struct {
	config
	builders []*UserCoinRewardCreate
	conflict []sql.ConflictOption
}

// Save creates the UserCoinReward entities in the database.
func (ucrcb *UserCoinRewardCreateBulk) Save(ctx context.Context) ([]*UserCoinReward, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucrcb.builders))
	nodes := make([]*UserCoinReward, len(ucrcb.builders))
	mutators := make([]Mutator, len(ucrcb.builders))
	for i := range ucrcb.builders {
		func(i int, root context.Context) {
			builder := ucrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCoinRewardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucrcb *UserCoinRewardCreateBulk) SaveX(ctx context.Context) []*UserCoinReward {
	v, err := ucrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucrcb *UserCoinRewardCreateBulk) Exec(ctx context.Context) error {
	_, err := ucrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucrcb *UserCoinRewardCreateBulk) ExecX(ctx context.Context) {
	if err := ucrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserCoinReward.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserCoinRewardUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ucrcb *UserCoinRewardCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserCoinRewardUpsertBulk {
	ucrcb.conflict = opts
	return &UserCoinRewardUpsertBulk{
		create: ucrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserCoinReward.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ucrcb *UserCoinRewardCreateBulk) OnConflictColumns(columns ...string) *UserCoinRewardUpsertBulk {
	ucrcb.conflict = append(ucrcb.conflict, sql.ConflictColumns(columns...))
	return &UserCoinRewardUpsertBulk{
		create: ucrcb,
	}
}

// UserCoinRewardUpsertBulk is the builder for "upsert"-ing
// a bulk of UserCoinReward nodes.
type UserCoinRewardUpsertBulk struct {
	create *UserCoinRewardCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserCoinReward.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usercoinreward.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserCoinRewardUpsertBulk) UpdateNewValues() *UserCoinRewardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(usercoinreward.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserCoinReward.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserCoinRewardUpsertBulk) Ignore() *UserCoinRewardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserCoinRewardUpsertBulk) DoNothing() *UserCoinRewardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCoinRewardCreateBulk.OnConflict
// documentation for more info.
func (u *UserCoinRewardUpsertBulk) Update(set func(*UserCoinRewardUpsert)) *UserCoinRewardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserCoinRewardUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserCoinRewardUpsertBulk) SetCreatedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserCoinRewardUpsertBulk) AddCreatedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateCreatedAt() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserCoinRewardUpsertBulk) SetUpdatedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserCoinRewardUpsertBulk) AddUpdatedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateUpdatedAt() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserCoinRewardUpsertBulk) SetDeletedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserCoinRewardUpsertBulk) AddDeletedAt(v uint32) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateDeletedAt() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *UserCoinRewardUpsertBulk) SetEntID(v uuid.UUID) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateEntID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserCoinRewardUpsertBulk) SetAppID(v uuid.UUID) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateAppID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserCoinRewardUpsertBulk) ClearAppID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserCoinRewardUpsertBulk) SetUserID(v uuid.UUID) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateUserID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserCoinRewardUpsertBulk) ClearUserID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UserCoinRewardUpsertBulk) SetCoinTypeID(v uuid.UUID) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateCoinTypeID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UserCoinRewardUpsertBulk) ClearCoinTypeID() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetCoinRewards sets the "coin_rewards" field.
func (u *UserCoinRewardUpsertBulk) SetCoinRewards(v decimal.Decimal) *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.SetCoinRewards(v)
	})
}

// UpdateCoinRewards sets the "coin_rewards" field to the value that was provided on create.
func (u *UserCoinRewardUpsertBulk) UpdateCoinRewards() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.UpdateCoinRewards()
	})
}

// ClearCoinRewards clears the value of the "coin_rewards" field.
func (u *UserCoinRewardUpsertBulk) ClearCoinRewards() *UserCoinRewardUpsertBulk {
	return u.Update(func(s *UserCoinRewardUpsert) {
		s.ClearCoinRewards()
	})
}

// Exec executes the query.
func (u *UserCoinRewardUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCoinRewardCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCoinRewardCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserCoinRewardUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}