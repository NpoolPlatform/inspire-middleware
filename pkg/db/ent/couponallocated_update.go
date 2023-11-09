// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CouponAllocatedUpdate is the builder for updating CouponAllocated entities.
type CouponAllocatedUpdate struct {
	config
	hooks     []Hook
	mutation  *CouponAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CouponAllocatedUpdate builder.
func (cau *CouponAllocatedUpdate) Where(ps ...predicate.CouponAllocated) *CouponAllocatedUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetCreatedAt sets the "created_at" field.
func (cau *CouponAllocatedUpdate) SetCreatedAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetCreatedAt()
	cau.mutation.SetCreatedAt(u)
	return cau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableCreatedAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetCreatedAt(*u)
	}
	return cau
}

// AddCreatedAt adds u to the "created_at" field.
func (cau *CouponAllocatedUpdate) AddCreatedAt(u int32) *CouponAllocatedUpdate {
	cau.mutation.AddCreatedAt(u)
	return cau
}

// SetUpdatedAt sets the "updated_at" field.
func (cau *CouponAllocatedUpdate) SetUpdatedAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetUpdatedAt()
	cau.mutation.SetUpdatedAt(u)
	return cau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cau *CouponAllocatedUpdate) AddUpdatedAt(u int32) *CouponAllocatedUpdate {
	cau.mutation.AddUpdatedAt(u)
	return cau
}

// SetDeletedAt sets the "deleted_at" field.
func (cau *CouponAllocatedUpdate) SetDeletedAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetDeletedAt()
	cau.mutation.SetDeletedAt(u)
	return cau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableDeletedAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetDeletedAt(*u)
	}
	return cau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cau *CouponAllocatedUpdate) AddDeletedAt(u int32) *CouponAllocatedUpdate {
	cau.mutation.AddDeletedAt(u)
	return cau
}

// SetAppID sets the "app_id" field.
func (cau *CouponAllocatedUpdate) SetAppID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetAppID(u)
	return cau
}

// SetUserID sets the "user_id" field.
func (cau *CouponAllocatedUpdate) SetUserID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetUserID(u)
	return cau
}

// SetCouponID sets the "coupon_id" field.
func (cau *CouponAllocatedUpdate) SetCouponID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetCouponID(u)
	return cau
}

// SetDenomination sets the "denomination" field.
func (cau *CouponAllocatedUpdate) SetDenomination(d decimal.Decimal) *CouponAllocatedUpdate {
	cau.mutation.SetDenomination(d)
	return cau
}

// SetNillableDenomination sets the "denomination" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableDenomination(d *decimal.Decimal) *CouponAllocatedUpdate {
	if d != nil {
		cau.SetDenomination(*d)
	}
	return cau
}

// ClearDenomination clears the value of the "denomination" field.
func (cau *CouponAllocatedUpdate) ClearDenomination() *CouponAllocatedUpdate {
	cau.mutation.ClearDenomination()
	return cau
}

// SetUsed sets the "used" field.
func (cau *CouponAllocatedUpdate) SetUsed(b bool) *CouponAllocatedUpdate {
	cau.mutation.SetUsed(b)
	return cau
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableUsed(b *bool) *CouponAllocatedUpdate {
	if b != nil {
		cau.SetUsed(*b)
	}
	return cau
}

// ClearUsed clears the value of the "used" field.
func (cau *CouponAllocatedUpdate) ClearUsed() *CouponAllocatedUpdate {
	cau.mutation.ClearUsed()
	return cau
}

// SetUsedAt sets the "used_at" field.
func (cau *CouponAllocatedUpdate) SetUsedAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetUsedAt()
	cau.mutation.SetUsedAt(u)
	return cau
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableUsedAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetUsedAt(*u)
	}
	return cau
}

// AddUsedAt adds u to the "used_at" field.
func (cau *CouponAllocatedUpdate) AddUsedAt(u int32) *CouponAllocatedUpdate {
	cau.mutation.AddUsedAt(u)
	return cau
}

// ClearUsedAt clears the value of the "used_at" field.
func (cau *CouponAllocatedUpdate) ClearUsedAt() *CouponAllocatedUpdate {
	cau.mutation.ClearUsedAt()
	return cau
}

// SetUsedByOrderID sets the "used_by_order_id" field.
func (cau *CouponAllocatedUpdate) SetUsedByOrderID(u uuid.UUID) *CouponAllocatedUpdate {
	cau.mutation.SetUsedByOrderID(u)
	return cau
}

// SetNillableUsedByOrderID sets the "used_by_order_id" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableUsedByOrderID(u *uuid.UUID) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetUsedByOrderID(*u)
	}
	return cau
}

// ClearUsedByOrderID clears the value of the "used_by_order_id" field.
func (cau *CouponAllocatedUpdate) ClearUsedByOrderID() *CouponAllocatedUpdate {
	cau.mutation.ClearUsedByOrderID()
	return cau
}

// SetStartAt sets the "start_at" field.
func (cau *CouponAllocatedUpdate) SetStartAt(u uint32) *CouponAllocatedUpdate {
	cau.mutation.ResetStartAt()
	cau.mutation.SetStartAt(u)
	return cau
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableStartAt(u *uint32) *CouponAllocatedUpdate {
	if u != nil {
		cau.SetStartAt(*u)
	}
	return cau
}

// AddStartAt adds u to the "start_at" field.
func (cau *CouponAllocatedUpdate) AddStartAt(u int32) *CouponAllocatedUpdate {
	cau.mutation.AddStartAt(u)
	return cau
}

// ClearStartAt clears the value of the "start_at" field.
func (cau *CouponAllocatedUpdate) ClearStartAt() *CouponAllocatedUpdate {
	cau.mutation.ClearStartAt()
	return cau
}

// SetCouponScope sets the "coupon_scope" field.
func (cau *CouponAllocatedUpdate) SetCouponScope(s string) *CouponAllocatedUpdate {
	cau.mutation.SetCouponScope(s)
	return cau
}

// SetNillableCouponScope sets the "coupon_scope" field if the given value is not nil.
func (cau *CouponAllocatedUpdate) SetNillableCouponScope(s *string) *CouponAllocatedUpdate {
	if s != nil {
		cau.SetCouponScope(*s)
	}
	return cau
}

// ClearCouponScope clears the value of the "coupon_scope" field.
func (cau *CouponAllocatedUpdate) ClearCouponScope() *CouponAllocatedUpdate {
	cau.mutation.ClearCouponScope()
	return cau
}

// Mutation returns the CouponAllocatedMutation object of the builder.
func (cau *CouponAllocatedUpdate) Mutation() *CouponAllocatedMutation {
	return cau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CouponAllocatedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cau.defaults(); err != nil {
		return 0, err
	}
	if len(cau.hooks) == 0 {
		affected, err = cau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (cau *CouponAllocatedUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CouponAllocatedUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CouponAllocatedUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cau *CouponAllocatedUpdate) defaults() error {
	if _, ok := cau.mutation.UpdatedAt(); !ok {
		if couponallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponallocated.UpdateDefaultUpdatedAt()
		cau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cau *CouponAllocatedUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponAllocatedUpdate {
	cau.modifiers = append(cau.modifiers, modifiers...)
	return cau
}

func (cau *CouponAllocatedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponallocated.Table,
			Columns: couponallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponallocated.FieldID,
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
			Column: couponallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreatedAt,
		})
	}
	if value, ok := cau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeletedAt,
		})
	}
	if value, ok := cau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldAppID,
		})
	}
	if value, ok := cau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUserID,
		})
	}
	if value, ok := cau.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldCouponID,
		})
	}
	if value, ok := cau.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponallocated.FieldDenomination,
		})
	}
	if cau.mutation.DenominationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponallocated.FieldDenomination,
		})
	}
	if value, ok := cau.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: couponallocated.FieldUsed,
		})
	}
	if cau.mutation.UsedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: couponallocated.FieldUsed,
		})
	}
	if value, ok := cau.mutation.UsedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if value, ok := cau.mutation.AddedUsedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if cau.mutation.UsedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if value, ok := cau.mutation.UsedByOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUsedByOrderID,
		})
	}
	if cau.mutation.UsedByOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponallocated.FieldUsedByOrderID,
		})
	}
	if value, ok := cau.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldStartAt,
		})
	}
	if value, ok := cau.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldStartAt,
		})
	}
	if cau.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponallocated.FieldStartAt,
		})
	}
	if value, ok := cau.mutation.CouponScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponallocated.FieldCouponScope,
		})
	}
	if cau.mutation.CouponScopeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponallocated.FieldCouponScope,
		})
	}
	_spec.Modifiers = cau.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponAllocatedUpdateOne is the builder for updating a single CouponAllocated entity.
type CouponAllocatedUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CouponAllocatedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cauo *CouponAllocatedUpdateOne) SetCreatedAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetCreatedAt()
	cauo.mutation.SetCreatedAt(u)
	return cauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableCreatedAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetCreatedAt(*u)
	}
	return cauo
}

// AddCreatedAt adds u to the "created_at" field.
func (cauo *CouponAllocatedUpdateOne) AddCreatedAt(u int32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddCreatedAt(u)
	return cauo
}

// SetUpdatedAt sets the "updated_at" field.
func (cauo *CouponAllocatedUpdateOne) SetUpdatedAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetUpdatedAt()
	cauo.mutation.SetUpdatedAt(u)
	return cauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cauo *CouponAllocatedUpdateOne) AddUpdatedAt(u int32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddUpdatedAt(u)
	return cauo
}

// SetDeletedAt sets the "deleted_at" field.
func (cauo *CouponAllocatedUpdateOne) SetDeletedAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetDeletedAt()
	cauo.mutation.SetDeletedAt(u)
	return cauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableDeletedAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetDeletedAt(*u)
	}
	return cauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cauo *CouponAllocatedUpdateOne) AddDeletedAt(u int32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddDeletedAt(u)
	return cauo
}

// SetAppID sets the "app_id" field.
func (cauo *CouponAllocatedUpdateOne) SetAppID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetAppID(u)
	return cauo
}

// SetUserID sets the "user_id" field.
func (cauo *CouponAllocatedUpdateOne) SetUserID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetUserID(u)
	return cauo
}

// SetCouponID sets the "coupon_id" field.
func (cauo *CouponAllocatedUpdateOne) SetCouponID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetCouponID(u)
	return cauo
}

// SetDenomination sets the "denomination" field.
func (cauo *CouponAllocatedUpdateOne) SetDenomination(d decimal.Decimal) *CouponAllocatedUpdateOne {
	cauo.mutation.SetDenomination(d)
	return cauo
}

// SetNillableDenomination sets the "denomination" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableDenomination(d *decimal.Decimal) *CouponAllocatedUpdateOne {
	if d != nil {
		cauo.SetDenomination(*d)
	}
	return cauo
}

// ClearDenomination clears the value of the "denomination" field.
func (cauo *CouponAllocatedUpdateOne) ClearDenomination() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearDenomination()
	return cauo
}

// SetUsed sets the "used" field.
func (cauo *CouponAllocatedUpdateOne) SetUsed(b bool) *CouponAllocatedUpdateOne {
	cauo.mutation.SetUsed(b)
	return cauo
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableUsed(b *bool) *CouponAllocatedUpdateOne {
	if b != nil {
		cauo.SetUsed(*b)
	}
	return cauo
}

// ClearUsed clears the value of the "used" field.
func (cauo *CouponAllocatedUpdateOne) ClearUsed() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearUsed()
	return cauo
}

// SetUsedAt sets the "used_at" field.
func (cauo *CouponAllocatedUpdateOne) SetUsedAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetUsedAt()
	cauo.mutation.SetUsedAt(u)
	return cauo
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableUsedAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetUsedAt(*u)
	}
	return cauo
}

// AddUsedAt adds u to the "used_at" field.
func (cauo *CouponAllocatedUpdateOne) AddUsedAt(u int32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddUsedAt(u)
	return cauo
}

// ClearUsedAt clears the value of the "used_at" field.
func (cauo *CouponAllocatedUpdateOne) ClearUsedAt() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearUsedAt()
	return cauo
}

// SetUsedByOrderID sets the "used_by_order_id" field.
func (cauo *CouponAllocatedUpdateOne) SetUsedByOrderID(u uuid.UUID) *CouponAllocatedUpdateOne {
	cauo.mutation.SetUsedByOrderID(u)
	return cauo
}

// SetNillableUsedByOrderID sets the "used_by_order_id" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableUsedByOrderID(u *uuid.UUID) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetUsedByOrderID(*u)
	}
	return cauo
}

// ClearUsedByOrderID clears the value of the "used_by_order_id" field.
func (cauo *CouponAllocatedUpdateOne) ClearUsedByOrderID() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearUsedByOrderID()
	return cauo
}

// SetStartAt sets the "start_at" field.
func (cauo *CouponAllocatedUpdateOne) SetStartAt(u uint32) *CouponAllocatedUpdateOne {
	cauo.mutation.ResetStartAt()
	cauo.mutation.SetStartAt(u)
	return cauo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableStartAt(u *uint32) *CouponAllocatedUpdateOne {
	if u != nil {
		cauo.SetStartAt(*u)
	}
	return cauo
}

// AddStartAt adds u to the "start_at" field.
func (cauo *CouponAllocatedUpdateOne) AddStartAt(u int32) *CouponAllocatedUpdateOne {
	cauo.mutation.AddStartAt(u)
	return cauo
}

// ClearStartAt clears the value of the "start_at" field.
func (cauo *CouponAllocatedUpdateOne) ClearStartAt() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearStartAt()
	return cauo
}

// SetCouponScope sets the "coupon_scope" field.
func (cauo *CouponAllocatedUpdateOne) SetCouponScope(s string) *CouponAllocatedUpdateOne {
	cauo.mutation.SetCouponScope(s)
	return cauo
}

// SetNillableCouponScope sets the "coupon_scope" field if the given value is not nil.
func (cauo *CouponAllocatedUpdateOne) SetNillableCouponScope(s *string) *CouponAllocatedUpdateOne {
	if s != nil {
		cauo.SetCouponScope(*s)
	}
	return cauo
}

// ClearCouponScope clears the value of the "coupon_scope" field.
func (cauo *CouponAllocatedUpdateOne) ClearCouponScope() *CouponAllocatedUpdateOne {
	cauo.mutation.ClearCouponScope()
	return cauo
}

// Mutation returns the CouponAllocatedMutation object of the builder.
func (cauo *CouponAllocatedUpdateOne) Mutation() *CouponAllocatedMutation {
	return cauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CouponAllocatedUpdateOne) Select(field string, fields ...string) *CouponAllocatedUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CouponAllocated entity.
func (cauo *CouponAllocatedUpdateOne) Save(ctx context.Context) (*CouponAllocated, error) {
	var (
		err  error
		node *CouponAllocated
	)
	if err := cauo.defaults(); err != nil {
		return nil, err
	}
	if len(cauo.hooks) == 0 {
		node, err = cauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponAllocatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
		nv, ok := v.(*CouponAllocated)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponAllocatedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CouponAllocatedUpdateOne) SaveX(ctx context.Context) *CouponAllocated {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CouponAllocatedUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CouponAllocatedUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cauo *CouponAllocatedUpdateOne) defaults() error {
	if _, ok := cauo.mutation.UpdatedAt(); !ok {
		if couponallocated.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponallocated.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponallocated.UpdateDefaultUpdatedAt()
		cauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cauo *CouponAllocatedUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponAllocatedUpdateOne {
	cauo.modifiers = append(cauo.modifiers, modifiers...)
	return cauo
}

func (cauo *CouponAllocatedUpdateOne) sqlSave(ctx context.Context) (_node *CouponAllocated, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponallocated.Table,
			Columns: couponallocated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponallocated.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CouponAllocated.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponallocated.FieldID)
		for _, f := range fields {
			if !couponallocated.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != couponallocated.FieldID {
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
			Column: couponallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldCreatedAt,
		})
	}
	if value, ok := cauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUpdatedAt,
		})
	}
	if value, ok := cauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldDeletedAt,
		})
	}
	if value, ok := cauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldAppID,
		})
	}
	if value, ok := cauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUserID,
		})
	}
	if value, ok := cauo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldCouponID,
		})
	}
	if value, ok := cauo.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponallocated.FieldDenomination,
		})
	}
	if cauo.mutation.DenominationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponallocated.FieldDenomination,
		})
	}
	if value, ok := cauo.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: couponallocated.FieldUsed,
		})
	}
	if cauo.mutation.UsedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: couponallocated.FieldUsed,
		})
	}
	if value, ok := cauo.mutation.UsedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if value, ok := cauo.mutation.AddedUsedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if cauo.mutation.UsedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponallocated.FieldUsedAt,
		})
	}
	if value, ok := cauo.mutation.UsedByOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponallocated.FieldUsedByOrderID,
		})
	}
	if cauo.mutation.UsedByOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: couponallocated.FieldUsedByOrderID,
		})
	}
	if value, ok := cauo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldStartAt,
		})
	}
	if value, ok := cauo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponallocated.FieldStartAt,
		})
	}
	if cauo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponallocated.FieldStartAt,
		})
	}
	if value, ok := cauo.mutation.CouponScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponallocated.FieldCouponScope,
		})
	}
	if cauo.mutation.CouponScopeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponallocated.FieldCouponScope,
		})
	}
	_spec.Modifiers = cauo.modifiers
	_node = &CouponAllocated{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponallocated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
