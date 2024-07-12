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
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinAllocatedCreate is the builder for creating a CoinAllocated entity.
type CoinAllocatedCreate struct {
	config
	mutation *CoinAllocatedMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cac *CoinAllocatedCreate) SetCreatedAt(u uint32) *CoinAllocatedCreate {
	cac.mutation.SetCreatedAt(u)
	return cac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableCreatedAt(u *uint32) *CoinAllocatedCreate {
	if u != nil {
		cac.SetCreatedAt(*u)
	}
	return cac
}

// SetUpdatedAt sets the "updated_at" field.
func (cac *CoinAllocatedCreate) SetUpdatedAt(u uint32) *CoinAllocatedCreate {
	cac.mutation.SetUpdatedAt(u)
	return cac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableUpdatedAt(u *uint32) *CoinAllocatedCreate {
	if u != nil {
		cac.SetUpdatedAt(*u)
	}
	return cac
}

// SetDeletedAt sets the "deleted_at" field.
func (cac *CoinAllocatedCreate) SetDeletedAt(u uint32) *CoinAllocatedCreate {
	cac.mutation.SetDeletedAt(u)
	return cac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableDeletedAt(u *uint32) *CoinAllocatedCreate {
	if u != nil {
		cac.SetDeletedAt(*u)
	}
	return cac
}

// SetEntID sets the "ent_id" field.
func (cac *CoinAllocatedCreate) SetEntID(u uuid.UUID) *CoinAllocatedCreate {
	cac.mutation.SetEntID(u)
	return cac
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableEntID(u *uuid.UUID) *CoinAllocatedCreate {
	if u != nil {
		cac.SetEntID(*u)
	}
	return cac
}

// SetAppID sets the "app_id" field.
func (cac *CoinAllocatedCreate) SetAppID(u uuid.UUID) *CoinAllocatedCreate {
	cac.mutation.SetAppID(u)
	return cac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableAppID(u *uuid.UUID) *CoinAllocatedCreate {
	if u != nil {
		cac.SetAppID(*u)
	}
	return cac
}

// SetCoinConfigID sets the "coin_config_id" field.
func (cac *CoinAllocatedCreate) SetCoinConfigID(u uuid.UUID) *CoinAllocatedCreate {
	cac.mutation.SetCoinConfigID(u)
	return cac
}

// SetNillableCoinConfigID sets the "coin_config_id" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableCoinConfigID(u *uuid.UUID) *CoinAllocatedCreate {
	if u != nil {
		cac.SetCoinConfigID(*u)
	}
	return cac
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cac *CoinAllocatedCreate) SetCoinTypeID(u uuid.UUID) *CoinAllocatedCreate {
	cac.mutation.SetCoinTypeID(u)
	return cac
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableCoinTypeID(u *uuid.UUID) *CoinAllocatedCreate {
	if u != nil {
		cac.SetCoinTypeID(*u)
	}
	return cac
}

// SetUserID sets the "user_id" field.
func (cac *CoinAllocatedCreate) SetUserID(u uuid.UUID) *CoinAllocatedCreate {
	cac.mutation.SetUserID(u)
	return cac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableUserID(u *uuid.UUID) *CoinAllocatedCreate {
	if u != nil {
		cac.SetUserID(*u)
	}
	return cac
}

// SetValue sets the "value" field.
func (cac *CoinAllocatedCreate) SetValue(d decimal.Decimal) *CoinAllocatedCreate {
	cac.mutation.SetValue(d)
	return cac
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cac *CoinAllocatedCreate) SetNillableValue(d *decimal.Decimal) *CoinAllocatedCreate {
	if d != nil {
		cac.SetValue(*d)
	}
	return cac
}

// SetID sets the "id" field.
func (cac *CoinAllocatedCreate) SetID(u uint32) *CoinAllocatedCreate {
	cac.mutation.SetID(u)
	return cac
}

// Mutation returns the CoinAllocatedMutation object of the builder.
func (cac *CoinAllocatedCreate) Mutation() *CoinAllocatedMutation {
	return cac.mutation
}

// Save creates the CoinAllocated in the database.
func (cac *CoinAllocatedCreate) Save(ctx context.Context) (*CoinAllocated, error) {
	var (
		err  error
		node *CoinAllocated
	)
	if err := cac.defaults(); err != nil {
		return nil, err
	}
	if len(cac.hooks) == 0 {
		if err = cac.check(); err != nil {
			return nil, err
		}
		node, err = cac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cac.check(); err != nil {
				return nil, err
			}
			cac.mutation = mutation
			if node, err = cac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cac.hooks) - 1; i >= 0; i-- {
			if cac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cac *CoinAllocatedCreate) SaveX(ctx context.Context) *CoinAllocated {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cac *CoinAllocatedCreate) Exec(ctx context.Context) error {
	_, err := cac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cac *CoinAllocatedCreate) ExecX(ctx context.Context) {
	if err := cac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cac *CoinAllocatedCreate) defaults() error {
	if _, ok := cac.mutation.CreatedAt(); !ok {
		if coinallocated.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultCreatedAt()
		cac.mutation.SetCreatedAt(v)
	}
	if _, ok := cac.mutation.UpdatedAt(); !ok {
		if coinallocated.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultUpdatedAt()
		cac.mutation.SetUpdatedAt(v)
	}
	if _, ok := cac.mutation.DeletedAt(); !ok {
		if coinallocated.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultDeletedAt()
		cac.mutation.SetDeletedAt(v)
	}
	if _, ok := cac.mutation.EntID(); !ok {
		if coinallocated.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultEntID()
		cac.mutation.SetEntID(v)
	}
	if _, ok := cac.mutation.AppID(); !ok {
		if coinallocated.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultAppID()
		cac.mutation.SetAppID(v)
	}
	if _, ok := cac.mutation.CoinConfigID(); !ok {
		if coinallocated.DefaultCoinConfigID == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultCoinConfigID (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultCoinConfigID()
		cac.mutation.SetCoinConfigID(v)
	}
	if _, ok := cac.mutation.CoinTypeID(); !ok {
		if coinallocated.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultCoinTypeID()
		cac.mutation.SetCoinTypeID(v)
	}
	if _, ok := cac.mutation.UserID(); !ok {
		if coinallocated.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized coinallocated.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := coinallocated.DefaultUserID()
		cac.mutation.SetUserID(v)
	}
	if _, ok := cac.mutation.Value(); !ok {
		v := coinallocated.DefaultValue
		cac.mutation.SetValue(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cac *CoinAllocatedCreate) check() error {
	if _, ok := cac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CoinAllocated.created_at"`)}
	}
	if _, ok := cac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CoinAllocated.updated_at"`)}
	}
	if _, ok := cac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CoinAllocated.deleted_at"`)}
	}
	if _, ok := cac.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "CoinAllocated.ent_id"`)}
	}
	return nil
}

func (cac *CoinAllocatedCreate) sqlSave(ctx context.Context) (*CoinAllocated, error) {
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
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

func (cac *CoinAllocatedCreate) createSpec() (*CoinAllocated, *sqlgraph.CreateSpec) {
	var (
		_node = &CoinAllocated{config: cac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coinallocated.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinallocated.FieldID,
			},
		}
	)
	_spec.OnConflict = cac.conflict
	if id, ok := cac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinallocated.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cac.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := cac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := cac.mutation.CoinConfigID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinConfigID,
		})
		_node.CoinConfigID = value
	}
	if value, ok := cac.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := cac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinallocated.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := cac.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: coinallocated.FieldValue,
		})
		_node.Value = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinAllocated.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinAllocatedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cac *CoinAllocatedCreate) OnConflict(opts ...sql.ConflictOption) *CoinAllocatedUpsertOne {
	cac.conflict = opts
	return &CoinAllocatedUpsertOne{
		create: cac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinAllocated.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cac *CoinAllocatedCreate) OnConflictColumns(columns ...string) *CoinAllocatedUpsertOne {
	cac.conflict = append(cac.conflict, sql.ConflictColumns(columns...))
	return &CoinAllocatedUpsertOne{
		create: cac,
	}
}

type (
	// CoinAllocatedUpsertOne is the builder for "upsert"-ing
	//  one CoinAllocated node.
	CoinAllocatedUpsertOne struct {
		create *CoinAllocatedCreate
	}

	// CoinAllocatedUpsert is the "OnConflict" setter.
	CoinAllocatedUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CoinAllocatedUpsert) SetCreatedAt(v uint32) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateCreatedAt() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinAllocatedUpsert) AddCreatedAt(v uint32) *CoinAllocatedUpsert {
	u.Add(coinallocated.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinAllocatedUpsert) SetUpdatedAt(v uint32) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateUpdatedAt() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinAllocatedUpsert) AddUpdatedAt(v uint32) *CoinAllocatedUpsert {
	u.Add(coinallocated.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinAllocatedUpsert) SetDeletedAt(v uint32) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateDeletedAt() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinAllocatedUpsert) AddDeletedAt(v uint32) *CoinAllocatedUpsert {
	u.Add(coinallocated.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *CoinAllocatedUpsert) SetEntID(v uuid.UUID) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateEntID() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *CoinAllocatedUpsert) SetAppID(v uuid.UUID) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateAppID() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *CoinAllocatedUpsert) ClearAppID() *CoinAllocatedUpsert {
	u.SetNull(coinallocated.FieldAppID)
	return u
}

// SetCoinConfigID sets the "coin_config_id" field.
func (u *CoinAllocatedUpsert) SetCoinConfigID(v uuid.UUID) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldCoinConfigID, v)
	return u
}

// UpdateCoinConfigID sets the "coin_config_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateCoinConfigID() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldCoinConfigID)
	return u
}

// ClearCoinConfigID clears the value of the "coin_config_id" field.
func (u *CoinAllocatedUpsert) ClearCoinConfigID() *CoinAllocatedUpsert {
	u.SetNull(coinallocated.FieldCoinConfigID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinAllocatedUpsert) SetCoinTypeID(v uuid.UUID) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateCoinTypeID() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CoinAllocatedUpsert) ClearCoinTypeID() *CoinAllocatedUpsert {
	u.SetNull(coinallocated.FieldCoinTypeID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *CoinAllocatedUpsert) SetUserID(v uuid.UUID) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateUserID() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *CoinAllocatedUpsert) ClearUserID() *CoinAllocatedUpsert {
	u.SetNull(coinallocated.FieldUserID)
	return u
}

// SetValue sets the "value" field.
func (u *CoinAllocatedUpsert) SetValue(v decimal.Decimal) *CoinAllocatedUpsert {
	u.Set(coinallocated.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *CoinAllocatedUpsert) UpdateValue() *CoinAllocatedUpsert {
	u.SetExcluded(coinallocated.FieldValue)
	return u
}

// ClearValue clears the value of the "value" field.
func (u *CoinAllocatedUpsert) ClearValue() *CoinAllocatedUpsert {
	u.SetNull(coinallocated.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CoinAllocated.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coinallocated.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinAllocatedUpsertOne) UpdateNewValues() *CoinAllocatedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(coinallocated.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CoinAllocated.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CoinAllocatedUpsertOne) Ignore() *CoinAllocatedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinAllocatedUpsertOne) DoNothing() *CoinAllocatedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinAllocatedCreate.OnConflict
// documentation for more info.
func (u *CoinAllocatedUpsertOne) Update(set func(*CoinAllocatedUpsert)) *CoinAllocatedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinAllocatedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinAllocatedUpsertOne) SetCreatedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinAllocatedUpsertOne) AddCreatedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateCreatedAt() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinAllocatedUpsertOne) SetUpdatedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinAllocatedUpsertOne) AddUpdatedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateUpdatedAt() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinAllocatedUpsertOne) SetDeletedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinAllocatedUpsertOne) AddDeletedAt(v uint32) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateDeletedAt() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CoinAllocatedUpsertOne) SetEntID(v uuid.UUID) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateEntID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *CoinAllocatedUpsertOne) SetAppID(v uuid.UUID) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateAppID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *CoinAllocatedUpsertOne) ClearAppID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearAppID()
	})
}

// SetCoinConfigID sets the "coin_config_id" field.
func (u *CoinAllocatedUpsertOne) SetCoinConfigID(v uuid.UUID) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCoinConfigID(v)
	})
}

// UpdateCoinConfigID sets the "coin_config_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateCoinConfigID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCoinConfigID()
	})
}

// ClearCoinConfigID clears the value of the "coin_config_id" field.
func (u *CoinAllocatedUpsertOne) ClearCoinConfigID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearCoinConfigID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinAllocatedUpsertOne) SetCoinTypeID(v uuid.UUID) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateCoinTypeID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CoinAllocatedUpsertOne) ClearCoinTypeID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetUserID sets the "user_id" field.
func (u *CoinAllocatedUpsertOne) SetUserID(v uuid.UUID) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateUserID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *CoinAllocatedUpsertOne) ClearUserID() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearUserID()
	})
}

// SetValue sets the "value" field.
func (u *CoinAllocatedUpsertOne) SetValue(v decimal.Decimal) *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *CoinAllocatedUpsertOne) UpdateValue() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *CoinAllocatedUpsertOne) ClearValue() *CoinAllocatedUpsertOne {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearValue()
	})
}

// Exec executes the query.
func (u *CoinAllocatedUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinAllocatedCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinAllocatedUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CoinAllocatedUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CoinAllocatedUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CoinAllocatedCreateBulk is the builder for creating many CoinAllocated entities in bulk.
type CoinAllocatedCreateBulk struct {
	config
	builders []*CoinAllocatedCreate
	conflict []sql.ConflictOption
}

// Save creates the CoinAllocated entities in the database.
func (cacb *CoinAllocatedCreateBulk) Save(ctx context.Context) ([]*CoinAllocated, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*CoinAllocated, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinAllocatedMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *CoinAllocatedCreateBulk) SaveX(ctx context.Context) []*CoinAllocated {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cacb *CoinAllocatedCreateBulk) Exec(ctx context.Context) error {
	_, err := cacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cacb *CoinAllocatedCreateBulk) ExecX(ctx context.Context) {
	if err := cacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinAllocated.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinAllocatedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cacb *CoinAllocatedCreateBulk) OnConflict(opts ...sql.ConflictOption) *CoinAllocatedUpsertBulk {
	cacb.conflict = opts
	return &CoinAllocatedUpsertBulk{
		create: cacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinAllocated.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cacb *CoinAllocatedCreateBulk) OnConflictColumns(columns ...string) *CoinAllocatedUpsertBulk {
	cacb.conflict = append(cacb.conflict, sql.ConflictColumns(columns...))
	return &CoinAllocatedUpsertBulk{
		create: cacb,
	}
}

// CoinAllocatedUpsertBulk is the builder for "upsert"-ing
// a bulk of CoinAllocated nodes.
type CoinAllocatedUpsertBulk struct {
	create *CoinAllocatedCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CoinAllocated.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coinallocated.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinAllocatedUpsertBulk) UpdateNewValues() *CoinAllocatedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(coinallocated.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CoinAllocated.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CoinAllocatedUpsertBulk) Ignore() *CoinAllocatedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinAllocatedUpsertBulk) DoNothing() *CoinAllocatedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinAllocatedCreateBulk.OnConflict
// documentation for more info.
func (u *CoinAllocatedUpsertBulk) Update(set func(*CoinAllocatedUpsert)) *CoinAllocatedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinAllocatedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinAllocatedUpsertBulk) SetCreatedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinAllocatedUpsertBulk) AddCreatedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateCreatedAt() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinAllocatedUpsertBulk) SetUpdatedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinAllocatedUpsertBulk) AddUpdatedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateUpdatedAt() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinAllocatedUpsertBulk) SetDeletedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinAllocatedUpsertBulk) AddDeletedAt(v uint32) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateDeletedAt() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CoinAllocatedUpsertBulk) SetEntID(v uuid.UUID) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateEntID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *CoinAllocatedUpsertBulk) SetAppID(v uuid.UUID) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateAppID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *CoinAllocatedUpsertBulk) ClearAppID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearAppID()
	})
}

// SetCoinConfigID sets the "coin_config_id" field.
func (u *CoinAllocatedUpsertBulk) SetCoinConfigID(v uuid.UUID) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCoinConfigID(v)
	})
}

// UpdateCoinConfigID sets the "coin_config_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateCoinConfigID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCoinConfigID()
	})
}

// ClearCoinConfigID clears the value of the "coin_config_id" field.
func (u *CoinAllocatedUpsertBulk) ClearCoinConfigID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearCoinConfigID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinAllocatedUpsertBulk) SetCoinTypeID(v uuid.UUID) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateCoinTypeID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CoinAllocatedUpsertBulk) ClearCoinTypeID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetUserID sets the "user_id" field.
func (u *CoinAllocatedUpsertBulk) SetUserID(v uuid.UUID) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateUserID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *CoinAllocatedUpsertBulk) ClearUserID() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearUserID()
	})
}

// SetValue sets the "value" field.
func (u *CoinAllocatedUpsertBulk) SetValue(v decimal.Decimal) *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *CoinAllocatedUpsertBulk) UpdateValue() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *CoinAllocatedUpsertBulk) ClearValue() *CoinAllocatedUpsertBulk {
	return u.Update(func(s *CoinAllocatedUpsert) {
		s.ClearValue()
	})
}

// Exec executes the query.
func (u *CoinAllocatedUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CoinAllocatedCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinAllocatedCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinAllocatedUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
