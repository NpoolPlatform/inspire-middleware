// Code generated by ent, DO NOT EDIT.

package appcommissionconfig

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// ThresholdAmount applies equality check predicate on the "threshold_amount" field. It's identical to ThresholdAmountEQ.
func ThresholdAmount(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThresholdAmount), v))
	})
}

// AmountOrPercent applies equality check predicate on the "amount_or_percent" field. It's identical to AmountOrPercentEQ.
func AmountOrPercent(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmountOrPercent), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// Invites applies equality check predicate on the "invites" field. It's identical to InvitesEQ.
func Invites(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvites), v))
	})
}

// SettleType applies equality check predicate on the "settle_type" field. It's identical to SettleTypeEQ.
func SettleType(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSettleType), v))
	})
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// ThresholdAmountEQ applies the EQ predicate on the "threshold_amount" field.
func ThresholdAmountEQ(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountNEQ applies the NEQ predicate on the "threshold_amount" field.
func ThresholdAmountNEQ(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountIn applies the In predicate on the "threshold_amount" field.
func ThresholdAmountIn(vs ...decimal.Decimal) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThresholdAmount), v...))
	})
}

// ThresholdAmountNotIn applies the NotIn predicate on the "threshold_amount" field.
func ThresholdAmountNotIn(vs ...decimal.Decimal) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThresholdAmount), v...))
	})
}

// ThresholdAmountGT applies the GT predicate on the "threshold_amount" field.
func ThresholdAmountGT(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountGTE applies the GTE predicate on the "threshold_amount" field.
func ThresholdAmountGTE(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountLT applies the LT predicate on the "threshold_amount" field.
func ThresholdAmountLT(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountLTE applies the LTE predicate on the "threshold_amount" field.
func ThresholdAmountLTE(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThresholdAmount), v))
	})
}

// ThresholdAmountIsNil applies the IsNil predicate on the "threshold_amount" field.
func ThresholdAmountIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThresholdAmount)))
	})
}

// ThresholdAmountNotNil applies the NotNil predicate on the "threshold_amount" field.
func ThresholdAmountNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThresholdAmount)))
	})
}

// AmountOrPercentEQ applies the EQ predicate on the "amount_or_percent" field.
func AmountOrPercentEQ(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentNEQ applies the NEQ predicate on the "amount_or_percent" field.
func AmountOrPercentNEQ(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentIn applies the In predicate on the "amount_or_percent" field.
func AmountOrPercentIn(vs ...decimal.Decimal) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmountOrPercent), v...))
	})
}

// AmountOrPercentNotIn applies the NotIn predicate on the "amount_or_percent" field.
func AmountOrPercentNotIn(vs ...decimal.Decimal) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmountOrPercent), v...))
	})
}

// AmountOrPercentGT applies the GT predicate on the "amount_or_percent" field.
func AmountOrPercentGT(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentGTE applies the GTE predicate on the "amount_or_percent" field.
func AmountOrPercentGTE(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentLT applies the LT predicate on the "amount_or_percent" field.
func AmountOrPercentLT(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentLTE applies the LTE predicate on the "amount_or_percent" field.
func AmountOrPercentLTE(v decimal.Decimal) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmountOrPercent), v))
	})
}

// AmountOrPercentIsNil applies the IsNil predicate on the "amount_or_percent" field.
func AmountOrPercentIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmountOrPercent)))
	})
}

// AmountOrPercentNotNil applies the NotNil predicate on the "amount_or_percent" field.
func AmountOrPercentNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmountOrPercent)))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// InvitesEQ applies the EQ predicate on the "invites" field.
func InvitesEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvites), v))
	})
}

// InvitesNEQ applies the NEQ predicate on the "invites" field.
func InvitesNEQ(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvites), v))
	})
}

// InvitesIn applies the In predicate on the "invites" field.
func InvitesIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInvites), v...))
	})
}

// InvitesNotIn applies the NotIn predicate on the "invites" field.
func InvitesNotIn(vs ...uint32) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInvites), v...))
	})
}

// InvitesGT applies the GT predicate on the "invites" field.
func InvitesGT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInvites), v))
	})
}

// InvitesGTE applies the GTE predicate on the "invites" field.
func InvitesGTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInvites), v))
	})
}

// InvitesLT applies the LT predicate on the "invites" field.
func InvitesLT(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInvites), v))
	})
}

// InvitesLTE applies the LTE predicate on the "invites" field.
func InvitesLTE(v uint32) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInvites), v))
	})
}

// InvitesIsNil applies the IsNil predicate on the "invites" field.
func InvitesIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInvites)))
	})
}

// InvitesNotNil applies the NotNil predicate on the "invites" field.
func InvitesNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInvites)))
	})
}

// SettleTypeEQ applies the EQ predicate on the "settle_type" field.
func SettleTypeEQ(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSettleType), v))
	})
}

// SettleTypeNEQ applies the NEQ predicate on the "settle_type" field.
func SettleTypeNEQ(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSettleType), v))
	})
}

// SettleTypeIn applies the In predicate on the "settle_type" field.
func SettleTypeIn(vs ...string) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSettleType), v...))
	})
}

// SettleTypeNotIn applies the NotIn predicate on the "settle_type" field.
func SettleTypeNotIn(vs ...string) predicate.AppCommissionConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSettleType), v...))
	})
}

// SettleTypeGT applies the GT predicate on the "settle_type" field.
func SettleTypeGT(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSettleType), v))
	})
}

// SettleTypeGTE applies the GTE predicate on the "settle_type" field.
func SettleTypeGTE(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSettleType), v))
	})
}

// SettleTypeLT applies the LT predicate on the "settle_type" field.
func SettleTypeLT(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSettleType), v))
	})
}

// SettleTypeLTE applies the LTE predicate on the "settle_type" field.
func SettleTypeLTE(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSettleType), v))
	})
}

// SettleTypeContains applies the Contains predicate on the "settle_type" field.
func SettleTypeContains(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSettleType), v))
	})
}

// SettleTypeHasPrefix applies the HasPrefix predicate on the "settle_type" field.
func SettleTypeHasPrefix(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSettleType), v))
	})
}

// SettleTypeHasSuffix applies the HasSuffix predicate on the "settle_type" field.
func SettleTypeHasSuffix(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSettleType), v))
	})
}

// SettleTypeIsNil applies the IsNil predicate on the "settle_type" field.
func SettleTypeIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSettleType)))
	})
}

// SettleTypeNotNil applies the NotNil predicate on the "settle_type" field.
func SettleTypeNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSettleType)))
	})
}

// SettleTypeEqualFold applies the EqualFold predicate on the "settle_type" field.
func SettleTypeEqualFold(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSettleType), v))
	})
}

// SettleTypeContainsFold applies the ContainsFold predicate on the "settle_type" field.
func SettleTypeContainsFold(v string) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSettleType), v))
	})
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisabled), v))
	})
}

// DisabledIsNil applies the IsNil predicate on the "disabled" field.
func DisabledIsNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDisabled)))
	})
}

// DisabledNotNil applies the NotNil predicate on the "disabled" field.
func DisabledNotNil() predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDisabled)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppCommissionConfig) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppCommissionConfig) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
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
func Not(p predicate.AppCommissionConfig) predicate.AppCommissionConfig {
	return predicate.AppCommissionConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
