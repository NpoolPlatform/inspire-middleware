// Code generated by ent, DO NOT EDIT.

package couponallocated

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CouponID applies equality check predicate on the "coupon_id" field. It's identical to CouponIDEQ.
func CouponID(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponID), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Used applies equality check predicate on the "used" field. It's identical to UsedEQ.
func Used(v bool) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsed), v))
	})
}

// UsedAt applies equality check predicate on the "used_at" field. It's identical to UsedAtEQ.
func UsedAt(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedAt), v))
	})
}

// UsedByOrderID applies equality check predicate on the "used_by_order_id" field. It's identical to UsedByOrderIDEQ.
func UsedByOrderID(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedByOrderID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// CouponIDEQ applies the EQ predicate on the "coupon_id" field.
func CouponIDEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponID), v))
	})
}

// CouponIDNEQ applies the NEQ predicate on the "coupon_id" field.
func CouponIDNEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCouponID), v))
	})
}

// CouponIDIn applies the In predicate on the "coupon_id" field.
func CouponIDIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCouponID), v...))
	})
}

// CouponIDNotIn applies the NotIn predicate on the "coupon_id" field.
func CouponIDNotIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCouponID), v...))
	})
}

// CouponIDGT applies the GT predicate on the "coupon_id" field.
func CouponIDGT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCouponID), v))
	})
}

// CouponIDGTE applies the GTE predicate on the "coupon_id" field.
func CouponIDGTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCouponID), v))
	})
}

// CouponIDLT applies the LT predicate on the "coupon_id" field.
func CouponIDLT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCouponID), v))
	})
}

// CouponIDLTE applies the LTE predicate on the "coupon_id" field.
func CouponIDLTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCouponID), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...decimal.Decimal) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...decimal.Decimal) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v decimal.Decimal) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValue)))
	})
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValue)))
	})
}

// UsedEQ applies the EQ predicate on the "used" field.
func UsedEQ(v bool) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsed), v))
	})
}

// UsedNEQ applies the NEQ predicate on the "used" field.
func UsedNEQ(v bool) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsed), v))
	})
}

// UsedIsNil applies the IsNil predicate on the "used" field.
func UsedIsNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsed)))
	})
}

// UsedNotNil applies the NotNil predicate on the "used" field.
func UsedNotNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsed)))
	})
}

// UsedAtEQ applies the EQ predicate on the "used_at" field.
func UsedAtEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedAt), v))
	})
}

// UsedAtNEQ applies the NEQ predicate on the "used_at" field.
func UsedAtNEQ(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedAt), v))
	})
}

// UsedAtIn applies the In predicate on the "used_at" field.
func UsedAtIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsedAt), v...))
	})
}

// UsedAtNotIn applies the NotIn predicate on the "used_at" field.
func UsedAtNotIn(vs ...uint32) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsedAt), v...))
	})
}

// UsedAtGT applies the GT predicate on the "used_at" field.
func UsedAtGT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedAt), v))
	})
}

// UsedAtGTE applies the GTE predicate on the "used_at" field.
func UsedAtGTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedAt), v))
	})
}

// UsedAtLT applies the LT predicate on the "used_at" field.
func UsedAtLT(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedAt), v))
	})
}

// UsedAtLTE applies the LTE predicate on the "used_at" field.
func UsedAtLTE(v uint32) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedAt), v))
	})
}

// UsedAtIsNil applies the IsNil predicate on the "used_at" field.
func UsedAtIsNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsedAt)))
	})
}

// UsedAtNotNil applies the NotNil predicate on the "used_at" field.
func UsedAtNotNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsedAt)))
	})
}

// UsedByOrderIDEQ applies the EQ predicate on the "used_by_order_id" field.
func UsedByOrderIDEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDNEQ applies the NEQ predicate on the "used_by_order_id" field.
func UsedByOrderIDNEQ(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDIn applies the In predicate on the "used_by_order_id" field.
func UsedByOrderIDIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsedByOrderID), v...))
	})
}

// UsedByOrderIDNotIn applies the NotIn predicate on the "used_by_order_id" field.
func UsedByOrderIDNotIn(vs ...uuid.UUID) predicate.CouponAllocated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsedByOrderID), v...))
	})
}

// UsedByOrderIDGT applies the GT predicate on the "used_by_order_id" field.
func UsedByOrderIDGT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDGTE applies the GTE predicate on the "used_by_order_id" field.
func UsedByOrderIDGTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDLT applies the LT predicate on the "used_by_order_id" field.
func UsedByOrderIDLT(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDLTE applies the LTE predicate on the "used_by_order_id" field.
func UsedByOrderIDLTE(v uuid.UUID) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedByOrderID), v))
	})
}

// UsedByOrderIDIsNil applies the IsNil predicate on the "used_by_order_id" field.
func UsedByOrderIDIsNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsedByOrderID)))
	})
}

// UsedByOrderIDNotNil applies the NotNil predicate on the "used_by_order_id" field.
func UsedByOrderIDNotNil() predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsedByOrderID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CouponAllocated) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CouponAllocated) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
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
func Not(p predicate.CouponAllocated) predicate.CouponAllocated {
	return predicate.CouponAllocated(func(s *sql.Selector) {
		p(s.Not())
	})
}