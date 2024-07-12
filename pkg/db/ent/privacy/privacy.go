// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AchievementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AchievementQueryRuleFunc func(context.Context, *ent.AchievementQuery) error

// EvalQuery return f(ctx, q).
func (f AchievementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AchievementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AchievementQuery", q)
}

// The AchievementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AchievementMutationRuleFunc func(context.Context, *ent.AchievementMutation) error

// EvalMutation calls f(ctx, m).
func (f AchievementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AchievementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AchievementMutation", m)
}

// The AchievementUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AchievementUserQueryRuleFunc func(context.Context, *ent.AchievementUserQuery) error

// EvalQuery return f(ctx, q).
func (f AchievementUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AchievementUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AchievementUserQuery", q)
}

// The AchievementUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AchievementUserMutationRuleFunc func(context.Context, *ent.AchievementUserMutation) error

// EvalMutation calls f(ctx, m).
func (f AchievementUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AchievementUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AchievementUserMutation", m)
}

// The AppCommissionConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppCommissionConfigQueryRuleFunc func(context.Context, *ent.AppCommissionConfigQuery) error

// EvalQuery return f(ctx, q).
func (f AppCommissionConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppCommissionConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppCommissionConfigQuery", q)
}

// The AppCommissionConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppCommissionConfigMutationRuleFunc func(context.Context, *ent.AppCommissionConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f AppCommissionConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppCommissionConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppCommissionConfigMutation", m)
}

// The AppConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppConfigQueryRuleFunc func(context.Context, *ent.AppConfigQuery) error

// EvalQuery return f(ctx, q).
func (f AppConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppConfigQuery", q)
}

// The AppConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppConfigMutationRuleFunc func(context.Context, *ent.AppConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f AppConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppConfigMutation", m)
}

// The AppGoodCommissionConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppGoodCommissionConfigQueryRuleFunc func(context.Context, *ent.AppGoodCommissionConfigQuery) error

// EvalQuery return f(ctx, q).
func (f AppGoodCommissionConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppGoodCommissionConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppGoodCommissionConfigQuery", q)
}

// The AppGoodCommissionConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppGoodCommissionConfigMutationRuleFunc func(context.Context, *ent.AppGoodCommissionConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f AppGoodCommissionConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppGoodCommissionConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppGoodCommissionConfigMutation", m)
}

// The AppGoodScopeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppGoodScopeQueryRuleFunc func(context.Context, *ent.AppGoodScopeQuery) error

// EvalQuery return f(ctx, q).
func (f AppGoodScopeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppGoodScopeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppGoodScopeQuery", q)
}

// The AppGoodScopeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppGoodScopeMutationRuleFunc func(context.Context, *ent.AppGoodScopeMutation) error

// EvalMutation calls f(ctx, m).
func (f AppGoodScopeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppGoodScopeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppGoodScopeMutation", m)
}

// The CashControlQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CashControlQueryRuleFunc func(context.Context, *ent.CashControlQuery) error

// EvalQuery return f(ctx, q).
func (f CashControlQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CashControlQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CashControlQuery", q)
}

// The CashControlMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CashControlMutationRuleFunc func(context.Context, *ent.CashControlMutation) error

// EvalMutation calls f(ctx, m).
func (f CashControlMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CashControlMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CashControlMutation", m)
}

// The CoinAllocatedQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinAllocatedQueryRuleFunc func(context.Context, *ent.CoinAllocatedQuery) error

// EvalQuery return f(ctx, q).
func (f CoinAllocatedQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CoinAllocatedQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CoinAllocatedQuery", q)
}

// The CoinAllocatedMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinAllocatedMutationRuleFunc func(context.Context, *ent.CoinAllocatedMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinAllocatedMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CoinAllocatedMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CoinAllocatedMutation", m)
}

// The CoinConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinConfigQueryRuleFunc func(context.Context, *ent.CoinConfigQuery) error

// EvalQuery return f(ctx, q).
func (f CoinConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CoinConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CoinConfigQuery", q)
}

// The CoinConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinConfigMutationRuleFunc func(context.Context, *ent.CoinConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CoinConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CoinConfigMutation", m)
}

// The CommissionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CommissionQueryRuleFunc func(context.Context, *ent.CommissionQuery) error

// EvalQuery return f(ctx, q).
func (f CommissionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CommissionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CommissionQuery", q)
}

// The CommissionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CommissionMutationRuleFunc func(context.Context, *ent.CommissionMutation) error

// EvalMutation calls f(ctx, m).
func (f CommissionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CommissionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CommissionMutation", m)
}

// The CouponQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CouponQueryRuleFunc func(context.Context, *ent.CouponQuery) error

// EvalQuery return f(ctx, q).
func (f CouponQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CouponQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CouponQuery", q)
}

// The CouponMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CouponMutationRuleFunc func(context.Context, *ent.CouponMutation) error

// EvalMutation calls f(ctx, m).
func (f CouponMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CouponMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CouponMutation", m)
}

// The CouponAllocatedQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CouponAllocatedQueryRuleFunc func(context.Context, *ent.CouponAllocatedQuery) error

// EvalQuery return f(ctx, q).
func (f CouponAllocatedQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CouponAllocatedQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CouponAllocatedQuery", q)
}

// The CouponAllocatedMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CouponAllocatedMutationRuleFunc func(context.Context, *ent.CouponAllocatedMutation) error

// EvalMutation calls f(ctx, m).
func (f CouponAllocatedMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CouponAllocatedMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CouponAllocatedMutation", m)
}

// The CouponScopeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CouponScopeQueryRuleFunc func(context.Context, *ent.CouponScopeQuery) error

// EvalQuery return f(ctx, q).
func (f CouponScopeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CouponScopeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CouponScopeQuery", q)
}

// The CouponScopeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CouponScopeMutationRuleFunc func(context.Context, *ent.CouponScopeMutation) error

// EvalMutation calls f(ctx, m).
func (f CouponScopeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CouponScopeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CouponScopeMutation", m)
}

// The EventQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EventQueryRuleFunc func(context.Context, *ent.EventQuery) error

// EvalQuery return f(ctx, q).
func (f EventQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EventQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EventQuery", q)
}

// The EventMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EventMutationRuleFunc func(context.Context, *ent.EventMutation) error

// EvalMutation calls f(ctx, m).
func (f EventMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EventMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EventMutation", m)
}

// The EventCoinQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EventCoinQueryRuleFunc func(context.Context, *ent.EventCoinQuery) error

// EvalQuery return f(ctx, q).
func (f EventCoinQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EventCoinQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EventCoinQuery", q)
}

// The EventCoinMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EventCoinMutationRuleFunc func(context.Context, *ent.EventCoinMutation) error

// EvalMutation calls f(ctx, m).
func (f EventCoinMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EventCoinMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EventCoinMutation", m)
}

// The EventCouponQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EventCouponQueryRuleFunc func(context.Context, *ent.EventCouponQuery) error

// EvalQuery return f(ctx, q).
func (f EventCouponQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EventCouponQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EventCouponQuery", q)
}

// The EventCouponMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EventCouponMutationRuleFunc func(context.Context, *ent.EventCouponMutation) error

// EvalMutation calls f(ctx, m).
func (f EventCouponMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EventCouponMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EventCouponMutation", m)
}

// The GoodAchievementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodAchievementQueryRuleFunc func(context.Context, *ent.GoodAchievementQuery) error

// EvalQuery return f(ctx, q).
func (f GoodAchievementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodAchievementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodAchievementQuery", q)
}

// The GoodAchievementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodAchievementMutationRuleFunc func(context.Context, *ent.GoodAchievementMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodAchievementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodAchievementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodAchievementMutation", m)
}

// The GoodCoinAchievementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodCoinAchievementQueryRuleFunc func(context.Context, *ent.GoodCoinAchievementQuery) error

// EvalQuery return f(ctx, q).
func (f GoodCoinAchievementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodCoinAchievementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodCoinAchievementQuery", q)
}

// The GoodCoinAchievementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodCoinAchievementMutationRuleFunc func(context.Context, *ent.GoodCoinAchievementMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodCoinAchievementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodCoinAchievementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodCoinAchievementMutation", m)
}

// The InvitationCodeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InvitationCodeQueryRuleFunc func(context.Context, *ent.InvitationCodeQuery) error

// EvalQuery return f(ctx, q).
func (f InvitationCodeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InvitationCodeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InvitationCodeQuery", q)
}

// The InvitationCodeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InvitationCodeMutationRuleFunc func(context.Context, *ent.InvitationCodeMutation) error

// EvalMutation calls f(ctx, m).
func (f InvitationCodeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InvitationCodeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InvitationCodeMutation", m)
}

// The OrderPaymentStatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrderPaymentStatementQueryRuleFunc func(context.Context, *ent.OrderPaymentStatementQuery) error

// EvalQuery return f(ctx, q).
func (f OrderPaymentStatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderPaymentStatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OrderPaymentStatementQuery", q)
}

// The OrderPaymentStatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrderPaymentStatementMutationRuleFunc func(context.Context, *ent.OrderPaymentStatementMutation) error

// EvalMutation calls f(ctx, m).
func (f OrderPaymentStatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OrderPaymentStatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OrderPaymentStatementMutation", m)
}

// The OrderStatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrderStatementQueryRuleFunc func(context.Context, *ent.OrderStatementQuery) error

// EvalQuery return f(ctx, q).
func (f OrderStatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderStatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OrderStatementQuery", q)
}

// The OrderStatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrderStatementMutationRuleFunc func(context.Context, *ent.OrderStatementMutation) error

// EvalMutation calls f(ctx, m).
func (f OrderStatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OrderStatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OrderStatementMutation", m)
}

// The PubsubMessageQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PubsubMessageQueryRuleFunc func(context.Context, *ent.PubsubMessageQuery) error

// EvalQuery return f(ctx, q).
func (f PubsubMessageQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PubsubMessageQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PubsubMessageQuery", q)
}

// The PubsubMessageMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PubsubMessageMutationRuleFunc func(context.Context, *ent.PubsubMessageMutation) error

// EvalMutation calls f(ctx, m).
func (f PubsubMessageMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PubsubMessageMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PubsubMessageMutation", m)
}

// The RegistrationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RegistrationQueryRuleFunc func(context.Context, *ent.RegistrationQuery) error

// EvalQuery return f(ctx, q).
func (f RegistrationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RegistrationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RegistrationQuery", q)
}

// The RegistrationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RegistrationMutationRuleFunc func(context.Context, *ent.RegistrationMutation) error

// EvalMutation calls f(ctx, m).
func (f RegistrationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RegistrationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RegistrationMutation", m)
}

// The StatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StatementQueryRuleFunc func(context.Context, *ent.StatementQuery) error

// EvalQuery return f(ctx, q).
func (f StatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StatementQuery", q)
}

// The StatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StatementMutationRuleFunc func(context.Context, *ent.StatementMutation) error

// EvalMutation calls f(ctx, m).
func (f StatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StatementMutation", m)
}

// The TaskConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TaskConfigQueryRuleFunc func(context.Context, *ent.TaskConfigQuery) error

// EvalQuery return f(ctx, q).
func (f TaskConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TaskConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TaskConfigQuery", q)
}

// The TaskConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TaskConfigMutationRuleFunc func(context.Context, *ent.TaskConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f TaskConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TaskConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TaskConfigMutation", m)
}

// The TaskUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TaskUserQueryRuleFunc func(context.Context, *ent.TaskUserQuery) error

// EvalQuery return f(ctx, q).
func (f TaskUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TaskUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TaskUserQuery", q)
}

// The TaskUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TaskUserMutationRuleFunc func(context.Context, *ent.TaskUserMutation) error

// EvalMutation calls f(ctx, m).
func (f TaskUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TaskUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TaskUserMutation", m)
}

// The UserCoinRewardQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserCoinRewardQueryRuleFunc func(context.Context, *ent.UserCoinRewardQuery) error

// EvalQuery return f(ctx, q).
func (f UserCoinRewardQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserCoinRewardQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserCoinRewardQuery", q)
}

// The UserCoinRewardMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserCoinRewardMutationRuleFunc func(context.Context, *ent.UserCoinRewardMutation) error

// EvalMutation calls f(ctx, m).
func (f UserCoinRewardMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserCoinRewardMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserCoinRewardMutation", m)
}

// The UserCreditHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserCreditHistoryQueryRuleFunc func(context.Context, *ent.UserCreditHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f UserCreditHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserCreditHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserCreditHistoryQuery", q)
}

// The UserCreditHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserCreditHistoryMutationRuleFunc func(context.Context, *ent.UserCreditHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f UserCreditHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserCreditHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserCreditHistoryMutation", m)
}

// The UserRewardQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserRewardQueryRuleFunc func(context.Context, *ent.UserRewardQuery) error

// EvalQuery return f(ctx, q).
func (f UserRewardQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserRewardQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserRewardQuery", q)
}

// The UserRewardMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserRewardMutationRuleFunc func(context.Context, *ent.UserRewardMutation) error

// EvalMutation calls f(ctx, m).
func (f UserRewardMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserRewardMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserRewardMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AchievementQuery:
		return q.Filter(), nil
	case *ent.AchievementUserQuery:
		return q.Filter(), nil
	case *ent.AppCommissionConfigQuery:
		return q.Filter(), nil
	case *ent.AppConfigQuery:
		return q.Filter(), nil
	case *ent.AppGoodCommissionConfigQuery:
		return q.Filter(), nil
	case *ent.AppGoodScopeQuery:
		return q.Filter(), nil
	case *ent.CashControlQuery:
		return q.Filter(), nil
	case *ent.CoinAllocatedQuery:
		return q.Filter(), nil
	case *ent.CoinConfigQuery:
		return q.Filter(), nil
	case *ent.CommissionQuery:
		return q.Filter(), nil
	case *ent.CouponQuery:
		return q.Filter(), nil
	case *ent.CouponAllocatedQuery:
		return q.Filter(), nil
	case *ent.CouponScopeQuery:
		return q.Filter(), nil
	case *ent.EventQuery:
		return q.Filter(), nil
	case *ent.EventCoinQuery:
		return q.Filter(), nil
	case *ent.EventCouponQuery:
		return q.Filter(), nil
	case *ent.GoodAchievementQuery:
		return q.Filter(), nil
	case *ent.GoodCoinAchievementQuery:
		return q.Filter(), nil
	case *ent.InvitationCodeQuery:
		return q.Filter(), nil
	case *ent.OrderPaymentStatementQuery:
		return q.Filter(), nil
	case *ent.OrderStatementQuery:
		return q.Filter(), nil
	case *ent.PubsubMessageQuery:
		return q.Filter(), nil
	case *ent.RegistrationQuery:
		return q.Filter(), nil
	case *ent.StatementQuery:
		return q.Filter(), nil
	case *ent.TaskConfigQuery:
		return q.Filter(), nil
	case *ent.TaskUserQuery:
		return q.Filter(), nil
	case *ent.UserCoinRewardQuery:
		return q.Filter(), nil
	case *ent.UserCreditHistoryQuery:
		return q.Filter(), nil
	case *ent.UserRewardQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AchievementMutation:
		return m.Filter(), nil
	case *ent.AchievementUserMutation:
		return m.Filter(), nil
	case *ent.AppCommissionConfigMutation:
		return m.Filter(), nil
	case *ent.AppConfigMutation:
		return m.Filter(), nil
	case *ent.AppGoodCommissionConfigMutation:
		return m.Filter(), nil
	case *ent.AppGoodScopeMutation:
		return m.Filter(), nil
	case *ent.CashControlMutation:
		return m.Filter(), nil
	case *ent.CoinAllocatedMutation:
		return m.Filter(), nil
	case *ent.CoinConfigMutation:
		return m.Filter(), nil
	case *ent.CommissionMutation:
		return m.Filter(), nil
	case *ent.CouponMutation:
		return m.Filter(), nil
	case *ent.CouponAllocatedMutation:
		return m.Filter(), nil
	case *ent.CouponScopeMutation:
		return m.Filter(), nil
	case *ent.EventMutation:
		return m.Filter(), nil
	case *ent.EventCoinMutation:
		return m.Filter(), nil
	case *ent.EventCouponMutation:
		return m.Filter(), nil
	case *ent.GoodAchievementMutation:
		return m.Filter(), nil
	case *ent.GoodCoinAchievementMutation:
		return m.Filter(), nil
	case *ent.InvitationCodeMutation:
		return m.Filter(), nil
	case *ent.OrderPaymentStatementMutation:
		return m.Filter(), nil
	case *ent.OrderStatementMutation:
		return m.Filter(), nil
	case *ent.PubsubMessageMutation:
		return m.Filter(), nil
	case *ent.RegistrationMutation:
		return m.Filter(), nil
	case *ent.StatementMutation:
		return m.Filter(), nil
	case *ent.TaskConfigMutation:
		return m.Filter(), nil
	case *ent.TaskUserMutation:
		return m.Filter(), nil
	case *ent.UserCoinRewardMutation:
		return m.Filter(), nil
	case *ent.UserCreditHistoryMutation:
		return m.Filter(), nil
	case *ent.UserRewardMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
