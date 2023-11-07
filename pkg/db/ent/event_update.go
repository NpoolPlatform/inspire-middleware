// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks     []Hook
	mutation  *EventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EventUpdate) SetCreatedAt(u uint32) *EventUpdate {
	eu.mutation.ResetCreatedAt()
	eu.mutation.SetCreatedAt(u)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedAt(u *uint32) *EventUpdate {
	if u != nil {
		eu.SetCreatedAt(*u)
	}
	return eu
}

// AddCreatedAt adds u to the "created_at" field.
func (eu *EventUpdate) AddCreatedAt(u int32) *EventUpdate {
	eu.mutation.AddCreatedAt(u)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EventUpdate) SetUpdatedAt(u uint32) *EventUpdate {
	eu.mutation.ResetUpdatedAt()
	eu.mutation.SetUpdatedAt(u)
	return eu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (eu *EventUpdate) AddUpdatedAt(u int32) *EventUpdate {
	eu.mutation.AddUpdatedAt(u)
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EventUpdate) SetDeletedAt(u uint32) *EventUpdate {
	eu.mutation.ResetDeletedAt()
	eu.mutation.SetDeletedAt(u)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableDeletedAt(u *uint32) *EventUpdate {
	if u != nil {
		eu.SetDeletedAt(*u)
	}
	return eu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (eu *EventUpdate) AddDeletedAt(u int32) *EventUpdate {
	eu.mutation.AddDeletedAt(u)
	return eu
}

// SetEntID sets the "ent_id" field.
func (eu *EventUpdate) SetEntID(u uuid.UUID) *EventUpdate {
	eu.mutation.SetEntID(u)
	return eu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEntID(u *uuid.UUID) *EventUpdate {
	if u != nil {
		eu.SetEntID(*u)
	}
	return eu
}

// SetAppID sets the "app_id" field.
func (eu *EventUpdate) SetAppID(u uuid.UUID) *EventUpdate {
	eu.mutation.SetAppID(u)
	return eu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableAppID(u *uuid.UUID) *EventUpdate {
	if u != nil {
		eu.SetAppID(*u)
	}
	return eu
}

// ClearAppID clears the value of the "app_id" field.
func (eu *EventUpdate) ClearAppID() *EventUpdate {
	eu.mutation.ClearAppID()
	return eu
}

// SetEventType sets the "event_type" field.
func (eu *EventUpdate) SetEventType(s string) *EventUpdate {
	eu.mutation.SetEventType(s)
	return eu
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventType(s *string) *EventUpdate {
	if s != nil {
		eu.SetEventType(*s)
	}
	return eu
}

// ClearEventType clears the value of the "event_type" field.
func (eu *EventUpdate) ClearEventType() *EventUpdate {
	eu.mutation.ClearEventType()
	return eu
}

// SetCouponIds sets the "coupon_ids" field.
func (eu *EventUpdate) SetCouponIds(u []uuid.UUID) *EventUpdate {
	eu.mutation.SetCouponIds(u)
	return eu
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (eu *EventUpdate) ClearCouponIds() *EventUpdate {
	eu.mutation.ClearCouponIds()
	return eu
}

// SetCredits sets the "credits" field.
func (eu *EventUpdate) SetCredits(d decimal.Decimal) *EventUpdate {
	eu.mutation.SetCredits(d)
	return eu
}

// SetNillableCredits sets the "credits" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCredits(d *decimal.Decimal) *EventUpdate {
	if d != nil {
		eu.SetCredits(*d)
	}
	return eu
}

// ClearCredits clears the value of the "credits" field.
func (eu *EventUpdate) ClearCredits() *EventUpdate {
	eu.mutation.ClearCredits()
	return eu
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (eu *EventUpdate) SetCreditsPerUsd(d decimal.Decimal) *EventUpdate {
	eu.mutation.SetCreditsPerUsd(d)
	return eu
}

// SetNillableCreditsPerUsd sets the "credits_per_usd" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreditsPerUsd(d *decimal.Decimal) *EventUpdate {
	if d != nil {
		eu.SetCreditsPerUsd(*d)
	}
	return eu
}

// ClearCreditsPerUsd clears the value of the "credits_per_usd" field.
func (eu *EventUpdate) ClearCreditsPerUsd() *EventUpdate {
	eu.mutation.ClearCreditsPerUsd()
	return eu
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (eu *EventUpdate) SetMaxConsecutive(u uint32) *EventUpdate {
	eu.mutation.ResetMaxConsecutive()
	eu.mutation.SetMaxConsecutive(u)
	return eu
}

// SetNillableMaxConsecutive sets the "max_consecutive" field if the given value is not nil.
func (eu *EventUpdate) SetNillableMaxConsecutive(u *uint32) *EventUpdate {
	if u != nil {
		eu.SetMaxConsecutive(*u)
	}
	return eu
}

// AddMaxConsecutive adds u to the "max_consecutive" field.
func (eu *EventUpdate) AddMaxConsecutive(u int32) *EventUpdate {
	eu.mutation.AddMaxConsecutive(u)
	return eu
}

// ClearMaxConsecutive clears the value of the "max_consecutive" field.
func (eu *EventUpdate) ClearMaxConsecutive() *EventUpdate {
	eu.mutation.ClearMaxConsecutive()
	return eu
}

// SetGoodID sets the "good_id" field.
func (eu *EventUpdate) SetGoodID(u uuid.UUID) *EventUpdate {
	eu.mutation.SetGoodID(u)
	return eu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableGoodID(u *uuid.UUID) *EventUpdate {
	if u != nil {
		eu.SetGoodID(*u)
	}
	return eu
}

// ClearGoodID clears the value of the "good_id" field.
func (eu *EventUpdate) ClearGoodID() *EventUpdate {
	eu.mutation.ClearGoodID()
	return eu
}

// SetAppGoodID sets the "app_good_id" field.
func (eu *EventUpdate) SetAppGoodID(u uuid.UUID) *EventUpdate {
	eu.mutation.SetAppGoodID(u)
	return eu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableAppGoodID(u *uuid.UUID) *EventUpdate {
	if u != nil {
		eu.SetAppGoodID(*u)
	}
	return eu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (eu *EventUpdate) ClearAppGoodID() *EventUpdate {
	eu.mutation.ClearAppGoodID()
	return eu
}

// SetInviterLayers sets the "inviter_layers" field.
func (eu *EventUpdate) SetInviterLayers(u uint32) *EventUpdate {
	eu.mutation.ResetInviterLayers()
	eu.mutation.SetInviterLayers(u)
	return eu
}

// SetNillableInviterLayers sets the "inviter_layers" field if the given value is not nil.
func (eu *EventUpdate) SetNillableInviterLayers(u *uint32) *EventUpdate {
	if u != nil {
		eu.SetInviterLayers(*u)
	}
	return eu
}

// AddInviterLayers adds u to the "inviter_layers" field.
func (eu *EventUpdate) AddInviterLayers(u int32) *EventUpdate {
	eu.mutation.AddInviterLayers(u)
	return eu
}

// ClearInviterLayers clears the value of the "inviter_layers" field.
func (eu *EventUpdate) ClearInviterLayers() *EventUpdate {
	eu.mutation.ClearInviterLayers()
	return eu
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EventUpdate) defaults() error {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		if event.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := event.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eu *EventUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventUpdate {
	eu.modifiers = append(eu.modifiers, modifiers...)
	return eu
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := eu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := eu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldDeletedAt,
		})
	}
	if value, ok := eu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldDeletedAt,
		})
	}
	if value, ok := eu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldEntID,
		})
	}
	if value, ok := eu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldAppID,
		})
	}
	if eu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldAppID,
		})
	}
	if value, ok := eu.mutation.EventType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventType,
		})
	}
	if eu.mutation.EventTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldEventType,
		})
	}
	if value, ok := eu.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldCouponIds,
		})
	}
	if eu.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: event.FieldCouponIds,
		})
	}
	if value, ok := eu.mutation.Credits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: event.FieldCredits,
		})
	}
	if eu.mutation.CreditsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: event.FieldCredits,
		})
	}
	if value, ok := eu.mutation.CreditsPerUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: event.FieldCreditsPerUsd,
		})
	}
	if eu.mutation.CreditsPerUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: event.FieldCreditsPerUsd,
		})
	}
	if value, ok := eu.mutation.MaxConsecutive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldMaxConsecutive,
		})
	}
	if value, ok := eu.mutation.AddedMaxConsecutive(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldMaxConsecutive,
		})
	}
	if eu.mutation.MaxConsecutiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: event.FieldMaxConsecutive,
		})
	}
	if value, ok := eu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldGoodID,
		})
	}
	if eu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldGoodID,
		})
	}
	if value, ok := eu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldAppGoodID,
		})
	}
	if eu.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldAppGoodID,
		})
	}
	if value, ok := eu.mutation.InviterLayers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldInviterLayers,
		})
	}
	if value, ok := eu.mutation.AddedInviterLayers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldInviterLayers,
		})
	}
	if eu.mutation.InviterLayersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: event.FieldInviterLayers,
		})
	}
	_spec.Modifiers = eu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (euo *EventUpdateOne) SetCreatedAt(u uint32) *EventUpdateOne {
	euo.mutation.ResetCreatedAt()
	euo.mutation.SetCreatedAt(u)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedAt(u *uint32) *EventUpdateOne {
	if u != nil {
		euo.SetCreatedAt(*u)
	}
	return euo
}

// AddCreatedAt adds u to the "created_at" field.
func (euo *EventUpdateOne) AddCreatedAt(u int32) *EventUpdateOne {
	euo.mutation.AddCreatedAt(u)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EventUpdateOne) SetUpdatedAt(u uint32) *EventUpdateOne {
	euo.mutation.ResetUpdatedAt()
	euo.mutation.SetUpdatedAt(u)
	return euo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (euo *EventUpdateOne) AddUpdatedAt(u int32) *EventUpdateOne {
	euo.mutation.AddUpdatedAt(u)
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EventUpdateOne) SetDeletedAt(u uint32) *EventUpdateOne {
	euo.mutation.ResetDeletedAt()
	euo.mutation.SetDeletedAt(u)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDeletedAt(u *uint32) *EventUpdateOne {
	if u != nil {
		euo.SetDeletedAt(*u)
	}
	return euo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (euo *EventUpdateOne) AddDeletedAt(u int32) *EventUpdateOne {
	euo.mutation.AddDeletedAt(u)
	return euo
}

// SetEntID sets the "ent_id" field.
func (euo *EventUpdateOne) SetEntID(u uuid.UUID) *EventUpdateOne {
	euo.mutation.SetEntID(u)
	return euo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEntID(u *uuid.UUID) *EventUpdateOne {
	if u != nil {
		euo.SetEntID(*u)
	}
	return euo
}

// SetAppID sets the "app_id" field.
func (euo *EventUpdateOne) SetAppID(u uuid.UUID) *EventUpdateOne {
	euo.mutation.SetAppID(u)
	return euo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableAppID(u *uuid.UUID) *EventUpdateOne {
	if u != nil {
		euo.SetAppID(*u)
	}
	return euo
}

// ClearAppID clears the value of the "app_id" field.
func (euo *EventUpdateOne) ClearAppID() *EventUpdateOne {
	euo.mutation.ClearAppID()
	return euo
}

// SetEventType sets the "event_type" field.
func (euo *EventUpdateOne) SetEventType(s string) *EventUpdateOne {
	euo.mutation.SetEventType(s)
	return euo
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventType(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetEventType(*s)
	}
	return euo
}

// ClearEventType clears the value of the "event_type" field.
func (euo *EventUpdateOne) ClearEventType() *EventUpdateOne {
	euo.mutation.ClearEventType()
	return euo
}

// SetCouponIds sets the "coupon_ids" field.
func (euo *EventUpdateOne) SetCouponIds(u []uuid.UUID) *EventUpdateOne {
	euo.mutation.SetCouponIds(u)
	return euo
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (euo *EventUpdateOne) ClearCouponIds() *EventUpdateOne {
	euo.mutation.ClearCouponIds()
	return euo
}

// SetCredits sets the "credits" field.
func (euo *EventUpdateOne) SetCredits(d decimal.Decimal) *EventUpdateOne {
	euo.mutation.SetCredits(d)
	return euo
}

// SetNillableCredits sets the "credits" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCredits(d *decimal.Decimal) *EventUpdateOne {
	if d != nil {
		euo.SetCredits(*d)
	}
	return euo
}

// ClearCredits clears the value of the "credits" field.
func (euo *EventUpdateOne) ClearCredits() *EventUpdateOne {
	euo.mutation.ClearCredits()
	return euo
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (euo *EventUpdateOne) SetCreditsPerUsd(d decimal.Decimal) *EventUpdateOne {
	euo.mutation.SetCreditsPerUsd(d)
	return euo
}

// SetNillableCreditsPerUsd sets the "credits_per_usd" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreditsPerUsd(d *decimal.Decimal) *EventUpdateOne {
	if d != nil {
		euo.SetCreditsPerUsd(*d)
	}
	return euo
}

// ClearCreditsPerUsd clears the value of the "credits_per_usd" field.
func (euo *EventUpdateOne) ClearCreditsPerUsd() *EventUpdateOne {
	euo.mutation.ClearCreditsPerUsd()
	return euo
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (euo *EventUpdateOne) SetMaxConsecutive(u uint32) *EventUpdateOne {
	euo.mutation.ResetMaxConsecutive()
	euo.mutation.SetMaxConsecutive(u)
	return euo
}

// SetNillableMaxConsecutive sets the "max_consecutive" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableMaxConsecutive(u *uint32) *EventUpdateOne {
	if u != nil {
		euo.SetMaxConsecutive(*u)
	}
	return euo
}

// AddMaxConsecutive adds u to the "max_consecutive" field.
func (euo *EventUpdateOne) AddMaxConsecutive(u int32) *EventUpdateOne {
	euo.mutation.AddMaxConsecutive(u)
	return euo
}

// ClearMaxConsecutive clears the value of the "max_consecutive" field.
func (euo *EventUpdateOne) ClearMaxConsecutive() *EventUpdateOne {
	euo.mutation.ClearMaxConsecutive()
	return euo
}

// SetGoodID sets the "good_id" field.
func (euo *EventUpdateOne) SetGoodID(u uuid.UUID) *EventUpdateOne {
	euo.mutation.SetGoodID(u)
	return euo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableGoodID(u *uuid.UUID) *EventUpdateOne {
	if u != nil {
		euo.SetGoodID(*u)
	}
	return euo
}

// ClearGoodID clears the value of the "good_id" field.
func (euo *EventUpdateOne) ClearGoodID() *EventUpdateOne {
	euo.mutation.ClearGoodID()
	return euo
}

// SetAppGoodID sets the "app_good_id" field.
func (euo *EventUpdateOne) SetAppGoodID(u uuid.UUID) *EventUpdateOne {
	euo.mutation.SetAppGoodID(u)
	return euo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *EventUpdateOne {
	if u != nil {
		euo.SetAppGoodID(*u)
	}
	return euo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (euo *EventUpdateOne) ClearAppGoodID() *EventUpdateOne {
	euo.mutation.ClearAppGoodID()
	return euo
}

// SetInviterLayers sets the "inviter_layers" field.
func (euo *EventUpdateOne) SetInviterLayers(u uint32) *EventUpdateOne {
	euo.mutation.ResetInviterLayers()
	euo.mutation.SetInviterLayers(u)
	return euo
}

// SetNillableInviterLayers sets the "inviter_layers" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableInviterLayers(u *uint32) *EventUpdateOne {
	if u != nil {
		euo.SetInviterLayers(*u)
	}
	return euo
}

// AddInviterLayers adds u to the "inviter_layers" field.
func (euo *EventUpdateOne) AddInviterLayers(u int32) *EventUpdateOne {
	euo.mutation.AddInviterLayers(u)
	return euo
}

// ClearInviterLayers clears the value of the "inviter_layers" field.
func (euo *EventUpdateOne) ClearInviterLayers() *EventUpdateOne {
	euo.mutation.ClearInviterLayers()
	return euo
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EventUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		if event.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := event.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (euo *EventUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventUpdateOne {
	euo.modifiers = append(euo.modifiers, modifiers...)
	return euo
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := euo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := euo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldDeletedAt,
		})
	}
	if value, ok := euo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldDeletedAt,
		})
	}
	if value, ok := euo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldEntID,
		})
	}
	if value, ok := euo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldAppID,
		})
	}
	if euo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldAppID,
		})
	}
	if value, ok := euo.mutation.EventType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventType,
		})
	}
	if euo.mutation.EventTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldEventType,
		})
	}
	if value, ok := euo.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldCouponIds,
		})
	}
	if euo.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: event.FieldCouponIds,
		})
	}
	if value, ok := euo.mutation.Credits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: event.FieldCredits,
		})
	}
	if euo.mutation.CreditsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: event.FieldCredits,
		})
	}
	if value, ok := euo.mutation.CreditsPerUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: event.FieldCreditsPerUsd,
		})
	}
	if euo.mutation.CreditsPerUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: event.FieldCreditsPerUsd,
		})
	}
	if value, ok := euo.mutation.MaxConsecutive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldMaxConsecutive,
		})
	}
	if value, ok := euo.mutation.AddedMaxConsecutive(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldMaxConsecutive,
		})
	}
	if euo.mutation.MaxConsecutiveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: event.FieldMaxConsecutive,
		})
	}
	if value, ok := euo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldGoodID,
		})
	}
	if euo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldGoodID,
		})
	}
	if value, ok := euo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: event.FieldAppGoodID,
		})
	}
	if euo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: event.FieldAppGoodID,
		})
	}
	if value, ok := euo.mutation.InviterLayers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldInviterLayers,
		})
	}
	if value, ok := euo.mutation.AddedInviterLayers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: event.FieldInviterLayers,
		})
	}
	if euo.mutation.InviterLayersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: event.FieldInviterLayers,
		})
	}
	_spec.Modifiers = euo.modifiers
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
