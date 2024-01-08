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
	"github.com/google/uuid"
)

// CouponCoinCreate is the builder for creating a CouponCoin entity.
type CouponCoinCreate struct {
	config
	mutation *CouponCoinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ccc *CouponCoinCreate) SetCreatedAt(u uint32) *CouponCoinCreate {
	ccc.mutation.SetCreatedAt(u)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableCreatedAt(u *uint32) *CouponCoinCreate {
	if u != nil {
		ccc.SetCreatedAt(*u)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *CouponCoinCreate) SetUpdatedAt(u uint32) *CouponCoinCreate {
	ccc.mutation.SetUpdatedAt(u)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableUpdatedAt(u *uint32) *CouponCoinCreate {
	if u != nil {
		ccc.SetUpdatedAt(*u)
	}
	return ccc
}

// SetDeletedAt sets the "deleted_at" field.
func (ccc *CouponCoinCreate) SetDeletedAt(u uint32) *CouponCoinCreate {
	ccc.mutation.SetDeletedAt(u)
	return ccc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableDeletedAt(u *uint32) *CouponCoinCreate {
	if u != nil {
		ccc.SetDeletedAt(*u)
	}
	return ccc
}

// SetEntID sets the "ent_id" field.
func (ccc *CouponCoinCreate) SetEntID(u uuid.UUID) *CouponCoinCreate {
	ccc.mutation.SetEntID(u)
	return ccc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableEntID(u *uuid.UUID) *CouponCoinCreate {
	if u != nil {
		ccc.SetEntID(*u)
	}
	return ccc
}

// SetAppID sets the "app_id" field.
func (ccc *CouponCoinCreate) SetAppID(u uuid.UUID) *CouponCoinCreate {
	ccc.mutation.SetAppID(u)
	return ccc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableAppID(u *uuid.UUID) *CouponCoinCreate {
	if u != nil {
		ccc.SetAppID(*u)
	}
	return ccc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccc *CouponCoinCreate) SetCoinTypeID(u uuid.UUID) *CouponCoinCreate {
	ccc.mutation.SetCoinTypeID(u)
	return ccc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ccc *CouponCoinCreate) SetNillableCoinTypeID(u *uuid.UUID) *CouponCoinCreate {
	if u != nil {
		ccc.SetCoinTypeID(*u)
	}
	return ccc
}

// SetID sets the "id" field.
func (ccc *CouponCoinCreate) SetID(u uint32) *CouponCoinCreate {
	ccc.mutation.SetID(u)
	return ccc
}

// Mutation returns the CouponCoinMutation object of the builder.
func (ccc *CouponCoinCreate) Mutation() *CouponCoinMutation {
	return ccc.mutation
}

// Save creates the CouponCoin in the database.
func (ccc *CouponCoinCreate) Save(ctx context.Context) (*CouponCoin, error) {
	var (
		err  error
		node *CouponCoin
	)
	if err := ccc.defaults(); err != nil {
		return nil, err
	}
	if len(ccc.hooks) == 0 {
		if err = ccc.check(); err != nil {
			return nil, err
		}
		node, err = ccc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccc.check(); err != nil {
				return nil, err
			}
			ccc.mutation = mutation
			if node, err = ccc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ccc.hooks) - 1; i >= 0; i-- {
			if ccc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ccc *CouponCoinCreate) SaveX(ctx context.Context) *CouponCoin {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CouponCoinCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CouponCoinCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *CouponCoinCreate) defaults() error {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		if couponcoin.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		if couponcoin.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ccc.mutation.DeletedAt(); !ok {
		if couponcoin.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultDeletedAt()
		ccc.mutation.SetDeletedAt(v)
	}
	if _, ok := ccc.mutation.EntID(); !ok {
		if couponcoin.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultEntID()
		ccc.mutation.SetEntID(v)
	}
	if _, ok := ccc.mutation.AppID(); !ok {
		if couponcoin.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultAppID()
		ccc.mutation.SetAppID(v)
	}
	if _, ok := ccc.mutation.CoinTypeID(); !ok {
		if couponcoin.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized couponcoin.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := couponcoin.DefaultCoinTypeID()
		ccc.mutation.SetCoinTypeID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CouponCoinCreate) check() error {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CouponCoin.created_at"`)}
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CouponCoin.updated_at"`)}
	}
	if _, ok := ccc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CouponCoin.deleted_at"`)}
	}
	if _, ok := ccc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "CouponCoin.ent_id"`)}
	}
	return nil
}

func (ccc *CouponCoinCreate) sqlSave(ctx context.Context) (*CouponCoin, error) {
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
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

func (ccc *CouponCoinCreate) createSpec() (*CouponCoin, *sqlgraph.CreateSpec) {
	var (
		_node = &CouponCoin{config: ccc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: couponcoin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: couponcoin.FieldID,
			},
		}
	)
	_spec.OnConflict = ccc.conflict
	if id, ok := ccc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ccc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponcoin.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ccc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := ccc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ccc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponcoin.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CouponCoin.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CouponCoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccc *CouponCoinCreate) OnConflict(opts ...sql.ConflictOption) *CouponCoinUpsertOne {
	ccc.conflict = opts
	return &CouponCoinUpsertOne{
		create: ccc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CouponCoin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccc *CouponCoinCreate) OnConflictColumns(columns ...string) *CouponCoinUpsertOne {
	ccc.conflict = append(ccc.conflict, sql.ConflictColumns(columns...))
	return &CouponCoinUpsertOne{
		create: ccc,
	}
}

type (
	// CouponCoinUpsertOne is the builder for "upsert"-ing
	//  one CouponCoin node.
	CouponCoinUpsertOne struct {
		create *CouponCoinCreate
	}

	// CouponCoinUpsert is the "OnConflict" setter.
	CouponCoinUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CouponCoinUpsert) SetCreatedAt(v uint32) *CouponCoinUpsert {
	u.Set(couponcoin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateCreatedAt() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CouponCoinUpsert) AddCreatedAt(v uint32) *CouponCoinUpsert {
	u.Add(couponcoin.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CouponCoinUpsert) SetUpdatedAt(v uint32) *CouponCoinUpsert {
	u.Set(couponcoin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateUpdatedAt() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CouponCoinUpsert) AddUpdatedAt(v uint32) *CouponCoinUpsert {
	u.Add(couponcoin.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CouponCoinUpsert) SetDeletedAt(v uint32) *CouponCoinUpsert {
	u.Set(couponcoin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateDeletedAt() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CouponCoinUpsert) AddDeletedAt(v uint32) *CouponCoinUpsert {
	u.Add(couponcoin.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *CouponCoinUpsert) SetEntID(v uuid.UUID) *CouponCoinUpsert {
	u.Set(couponcoin.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateEntID() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *CouponCoinUpsert) SetAppID(v uuid.UUID) *CouponCoinUpsert {
	u.Set(couponcoin.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateAppID() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *CouponCoinUpsert) ClearAppID() *CouponCoinUpsert {
	u.SetNull(couponcoin.FieldAppID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CouponCoinUpsert) SetCoinTypeID(v uuid.UUID) *CouponCoinUpsert {
	u.Set(couponcoin.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CouponCoinUpsert) UpdateCoinTypeID() *CouponCoinUpsert {
	u.SetExcluded(couponcoin.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CouponCoinUpsert) ClearCoinTypeID() *CouponCoinUpsert {
	u.SetNull(couponcoin.FieldCoinTypeID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CouponCoin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(couponcoin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CouponCoinUpsertOne) UpdateNewValues() *CouponCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(couponcoin.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CouponCoin.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CouponCoinUpsertOne) Ignore() *CouponCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CouponCoinUpsertOne) DoNothing() *CouponCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CouponCoinCreate.OnConflict
// documentation for more info.
func (u *CouponCoinUpsertOne) Update(set func(*CouponCoinUpsert)) *CouponCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CouponCoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CouponCoinUpsertOne) SetCreatedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CouponCoinUpsertOne) AddCreatedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateCreatedAt() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CouponCoinUpsertOne) SetUpdatedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CouponCoinUpsertOne) AddUpdatedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateUpdatedAt() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CouponCoinUpsertOne) SetDeletedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CouponCoinUpsertOne) AddDeletedAt(v uint32) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateDeletedAt() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CouponCoinUpsertOne) SetEntID(v uuid.UUID) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateEntID() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *CouponCoinUpsertOne) SetAppID(v uuid.UUID) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateAppID() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *CouponCoinUpsertOne) ClearAppID() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.ClearAppID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CouponCoinUpsertOne) SetCoinTypeID(v uuid.UUID) *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CouponCoinUpsertOne) UpdateCoinTypeID() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CouponCoinUpsertOne) ClearCoinTypeID() *CouponCoinUpsertOne {
	return u.Update(func(s *CouponCoinUpsert) {
		s.ClearCoinTypeID()
	})
}

// Exec executes the query.
func (u *CouponCoinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CouponCoinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CouponCoinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CouponCoinUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CouponCoinUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CouponCoinCreateBulk is the builder for creating many CouponCoin entities in bulk.
type CouponCoinCreateBulk struct {
	config
	builders []*CouponCoinCreate
	conflict []sql.ConflictOption
}

// Save creates the CouponCoin entities in the database.
func (cccb *CouponCoinCreateBulk) Save(ctx context.Context) ([]*CouponCoin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CouponCoin, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CouponCoinMutation)
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
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *CouponCoinCreateBulk) SaveX(ctx context.Context) []*CouponCoin {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CouponCoinCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CouponCoinCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CouponCoin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CouponCoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cccb *CouponCoinCreateBulk) OnConflict(opts ...sql.ConflictOption) *CouponCoinUpsertBulk {
	cccb.conflict = opts
	return &CouponCoinUpsertBulk{
		create: cccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CouponCoin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cccb *CouponCoinCreateBulk) OnConflictColumns(columns ...string) *CouponCoinUpsertBulk {
	cccb.conflict = append(cccb.conflict, sql.ConflictColumns(columns...))
	return &CouponCoinUpsertBulk{
		create: cccb,
	}
}

// CouponCoinUpsertBulk is the builder for "upsert"-ing
// a bulk of CouponCoin nodes.
type CouponCoinUpsertBulk struct {
	create *CouponCoinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CouponCoin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(couponcoin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CouponCoinUpsertBulk) UpdateNewValues() *CouponCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(couponcoin.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CouponCoin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CouponCoinUpsertBulk) Ignore() *CouponCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CouponCoinUpsertBulk) DoNothing() *CouponCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CouponCoinCreateBulk.OnConflict
// documentation for more info.
func (u *CouponCoinUpsertBulk) Update(set func(*CouponCoinUpsert)) *CouponCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CouponCoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CouponCoinUpsertBulk) SetCreatedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CouponCoinUpsertBulk) AddCreatedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateCreatedAt() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CouponCoinUpsertBulk) SetUpdatedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CouponCoinUpsertBulk) AddUpdatedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateUpdatedAt() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CouponCoinUpsertBulk) SetDeletedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CouponCoinUpsertBulk) AddDeletedAt(v uint32) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateDeletedAt() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CouponCoinUpsertBulk) SetEntID(v uuid.UUID) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateEntID() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *CouponCoinUpsertBulk) SetAppID(v uuid.UUID) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateAppID() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *CouponCoinUpsertBulk) ClearAppID() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.ClearAppID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CouponCoinUpsertBulk) SetCoinTypeID(v uuid.UUID) *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CouponCoinUpsertBulk) UpdateCoinTypeID() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CouponCoinUpsertBulk) ClearCoinTypeID() *CouponCoinUpsertBulk {
	return u.Update(func(s *CouponCoinUpsert) {
		s.ClearCoinTypeID()
	})
}

// Exec executes the query.
func (u *CouponCoinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CouponCoinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CouponCoinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CouponCoinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
