// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

// The AchievementFunc type is an adapter to allow the use of ordinary
// function as Achievement mutator.
type AchievementFunc func(context.Context, *ent.AchievementMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AchievementFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AchievementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AchievementMutation", m)
	}
	return f(ctx, mv)
}

// The CommissionFunc type is an adapter to allow the use of ordinary
// function as Commission mutator.
type CommissionFunc func(context.Context, *ent.CommissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CommissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommissionMutation", m)
	}
	return f(ctx, mv)
}

// The CouponFunc type is an adapter to allow the use of ordinary
// function as Coupon mutator.
type CouponFunc func(context.Context, *ent.CouponMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CouponFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CouponMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CouponMutation", m)
	}
	return f(ctx, mv)
}

// The CouponAllocatedFunc type is an adapter to allow the use of ordinary
// function as CouponAllocated mutator.
type CouponAllocatedFunc func(context.Context, *ent.CouponAllocatedMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CouponAllocatedFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CouponAllocatedMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CouponAllocatedMutation", m)
	}
	return f(ctx, mv)
}

// The EventFunc type is an adapter to allow the use of ordinary
// function as Event mutator.
type EventFunc func(context.Context, *ent.EventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EventMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EventMutation", m)
	}
	return f(ctx, mv)
}

// The InvitationCodeFunc type is an adapter to allow the use of ordinary
// function as InvitationCode mutator.
type InvitationCodeFunc func(context.Context, *ent.InvitationCodeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InvitationCodeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InvitationCodeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InvitationCodeMutation", m)
	}
	return f(ctx, mv)
}

// The PubsubMessageFunc type is an adapter to allow the use of ordinary
// function as PubsubMessage mutator.
type PubsubMessageFunc func(context.Context, *ent.PubsubMessageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PubsubMessageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PubsubMessageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PubsubMessageMutation", m)
	}
	return f(ctx, mv)
}

// The RegistrationFunc type is an adapter to allow the use of ordinary
// function as Registration mutator.
type RegistrationFunc func(context.Context, *ent.RegistrationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegistrationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RegistrationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegistrationMutation", m)
	}
	return f(ctx, mv)
}

// The StatementFunc type is an adapter to allow the use of ordinary
// function as Statement mutator.
type StatementFunc func(context.Context, *ent.StatementMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatementFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatementMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatementMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
