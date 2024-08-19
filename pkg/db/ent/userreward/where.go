// Code generated by ent, DO NOT EDIT.

package userreward

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ActionCredits applies equality check predicate on the "action_credits" field. It's identical to ActionCreditsEQ.
func ActionCredits(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActionCredits), v))
	})
}

// CouponAmount applies equality check predicate on the "coupon_amount" field. It's identical to CouponAmountEQ.
func CouponAmount(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponAmount), v))
	})
}

// CouponCashableAmount applies equality check predicate on the "coupon_cashable_amount" field. It's identical to CouponCashableAmountEQ.
func CouponCashableAmount(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponCashableAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// ActionCreditsEQ applies the EQ predicate on the "action_credits" field.
func ActionCreditsEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsNEQ applies the NEQ predicate on the "action_credits" field.
func ActionCreditsNEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsIn applies the In predicate on the "action_credits" field.
func ActionCreditsIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldActionCredits), v...))
	})
}

// ActionCreditsNotIn applies the NotIn predicate on the "action_credits" field.
func ActionCreditsNotIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldActionCredits), v...))
	})
}

// ActionCreditsGT applies the GT predicate on the "action_credits" field.
func ActionCreditsGT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsGTE applies the GTE predicate on the "action_credits" field.
func ActionCreditsGTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsLT applies the LT predicate on the "action_credits" field.
func ActionCreditsLT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsLTE applies the LTE predicate on the "action_credits" field.
func ActionCreditsLTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActionCredits), v))
	})
}

// ActionCreditsIsNil applies the IsNil predicate on the "action_credits" field.
func ActionCreditsIsNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldActionCredits)))
	})
}

// ActionCreditsNotNil applies the NotNil predicate on the "action_credits" field.
func ActionCreditsNotNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldActionCredits)))
	})
}

// CouponAmountEQ applies the EQ predicate on the "coupon_amount" field.
func CouponAmountEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountNEQ applies the NEQ predicate on the "coupon_amount" field.
func CouponAmountNEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountIn applies the In predicate on the "coupon_amount" field.
func CouponAmountIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCouponAmount), v...))
	})
}

// CouponAmountNotIn applies the NotIn predicate on the "coupon_amount" field.
func CouponAmountNotIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCouponAmount), v...))
	})
}

// CouponAmountGT applies the GT predicate on the "coupon_amount" field.
func CouponAmountGT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountGTE applies the GTE predicate on the "coupon_amount" field.
func CouponAmountGTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountLT applies the LT predicate on the "coupon_amount" field.
func CouponAmountLT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountLTE applies the LTE predicate on the "coupon_amount" field.
func CouponAmountLTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCouponAmount), v))
	})
}

// CouponAmountIsNil applies the IsNil predicate on the "coupon_amount" field.
func CouponAmountIsNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCouponAmount)))
	})
}

// CouponAmountNotNil applies the NotNil predicate on the "coupon_amount" field.
func CouponAmountNotNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCouponAmount)))
	})
}

// CouponCashableAmountEQ applies the EQ predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountNEQ applies the NEQ predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountNEQ(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountIn applies the In predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCouponCashableAmount), v...))
	})
}

// CouponCashableAmountNotIn applies the NotIn predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountNotIn(vs ...decimal.Decimal) predicate.UserReward {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCouponCashableAmount), v...))
	})
}

// CouponCashableAmountGT applies the GT predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountGT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountGTE applies the GTE predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountGTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountLT applies the LT predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountLT(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountLTE applies the LTE predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountLTE(v decimal.Decimal) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCouponCashableAmount), v))
	})
}

// CouponCashableAmountIsNil applies the IsNil predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountIsNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCouponCashableAmount)))
	})
}

// CouponCashableAmountNotNil applies the NotNil predicate on the "coupon_cashable_amount" field.
func CouponCashableAmountNotNil() predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCouponCashableAmount)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserReward) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserReward) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserReward) predicate.UserReward {
	return predicate.UserReward(func(s *sql.Selector) {
		p(s.Not())
	})
}