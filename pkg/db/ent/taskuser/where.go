// Code generated by ent, DO NOT EDIT.

package taskuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// TaskState applies equality check predicate on the "task_state" field. It's identical to TaskStateEQ.
func TaskState(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskState), v))
	})
}

// RewardInfo applies equality check predicate on the "reward_info" field. It's identical to RewardInfoEQ.
func RewardInfo(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardInfo), v))
	})
}

// RewardState applies equality check predicate on the "reward_state" field. It's identical to RewardStateEQ.
func RewardState(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardState), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskID), v))
	})
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTaskID), v...))
	})
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTaskID), v...))
	})
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskID), v))
	})
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskID), v))
	})
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskID), v))
	})
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskID), v))
	})
}

// TaskIDIsNil applies the IsNil predicate on the "task_id" field.
func TaskIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTaskID)))
	})
}

// TaskIDNotNil applies the NotNil predicate on the "task_id" field.
func TaskIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTaskID)))
	})
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventID), v))
	})
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEventID), v...))
	})
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEventID), v...))
	})
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventID), v))
	})
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventID), v))
	})
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventID), v))
	})
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventID), v))
	})
}

// EventIDIsNil applies the IsNil predicate on the "event_id" field.
func EventIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEventID)))
	})
}

// EventIDNotNil applies the NotNil predicate on the "event_id" field.
func EventIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEventID)))
	})
}

// TaskStateEQ applies the EQ predicate on the "task_state" field.
func TaskStateEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskState), v))
	})
}

// TaskStateNEQ applies the NEQ predicate on the "task_state" field.
func TaskStateNEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskState), v))
	})
}

// TaskStateIn applies the In predicate on the "task_state" field.
func TaskStateIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTaskState), v...))
	})
}

// TaskStateNotIn applies the NotIn predicate on the "task_state" field.
func TaskStateNotIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTaskState), v...))
	})
}

// TaskStateGT applies the GT predicate on the "task_state" field.
func TaskStateGT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskState), v))
	})
}

// TaskStateGTE applies the GTE predicate on the "task_state" field.
func TaskStateGTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskState), v))
	})
}

// TaskStateLT applies the LT predicate on the "task_state" field.
func TaskStateLT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskState), v))
	})
}

// TaskStateLTE applies the LTE predicate on the "task_state" field.
func TaskStateLTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskState), v))
	})
}

// TaskStateContains applies the Contains predicate on the "task_state" field.
func TaskStateContains(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskState), v))
	})
}

// TaskStateHasPrefix applies the HasPrefix predicate on the "task_state" field.
func TaskStateHasPrefix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskState), v))
	})
}

// TaskStateHasSuffix applies the HasSuffix predicate on the "task_state" field.
func TaskStateHasSuffix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskState), v))
	})
}

// TaskStateIsNil applies the IsNil predicate on the "task_state" field.
func TaskStateIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTaskState)))
	})
}

// TaskStateNotNil applies the NotNil predicate on the "task_state" field.
func TaskStateNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTaskState)))
	})
}

// TaskStateEqualFold applies the EqualFold predicate on the "task_state" field.
func TaskStateEqualFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskState), v))
	})
}

// TaskStateContainsFold applies the ContainsFold predicate on the "task_state" field.
func TaskStateContainsFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskState), v))
	})
}

// RewardInfoEQ applies the EQ predicate on the "reward_info" field.
func RewardInfoEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoNEQ applies the NEQ predicate on the "reward_info" field.
func RewardInfoNEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoIn applies the In predicate on the "reward_info" field.
func RewardInfoIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRewardInfo), v...))
	})
}

// RewardInfoNotIn applies the NotIn predicate on the "reward_info" field.
func RewardInfoNotIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRewardInfo), v...))
	})
}

// RewardInfoGT applies the GT predicate on the "reward_info" field.
func RewardInfoGT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoGTE applies the GTE predicate on the "reward_info" field.
func RewardInfoGTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoLT applies the LT predicate on the "reward_info" field.
func RewardInfoLT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoLTE applies the LTE predicate on the "reward_info" field.
func RewardInfoLTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoContains applies the Contains predicate on the "reward_info" field.
func RewardInfoContains(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoHasPrefix applies the HasPrefix predicate on the "reward_info" field.
func RewardInfoHasPrefix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoHasSuffix applies the HasSuffix predicate on the "reward_info" field.
func RewardInfoHasSuffix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoIsNil applies the IsNil predicate on the "reward_info" field.
func RewardInfoIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRewardInfo)))
	})
}

// RewardInfoNotNil applies the NotNil predicate on the "reward_info" field.
func RewardInfoNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRewardInfo)))
	})
}

// RewardInfoEqualFold applies the EqualFold predicate on the "reward_info" field.
func RewardInfoEqualFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRewardInfo), v))
	})
}

// RewardInfoContainsFold applies the ContainsFold predicate on the "reward_info" field.
func RewardInfoContainsFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRewardInfo), v))
	})
}

// RewardStateEQ applies the EQ predicate on the "reward_state" field.
func RewardStateEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardState), v))
	})
}

// RewardStateNEQ applies the NEQ predicate on the "reward_state" field.
func RewardStateNEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRewardState), v))
	})
}

// RewardStateIn applies the In predicate on the "reward_state" field.
func RewardStateIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRewardState), v...))
	})
}

// RewardStateNotIn applies the NotIn predicate on the "reward_state" field.
func RewardStateNotIn(vs ...string) predicate.TaskUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRewardState), v...))
	})
}

// RewardStateGT applies the GT predicate on the "reward_state" field.
func RewardStateGT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRewardState), v))
	})
}

// RewardStateGTE applies the GTE predicate on the "reward_state" field.
func RewardStateGTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRewardState), v))
	})
}

// RewardStateLT applies the LT predicate on the "reward_state" field.
func RewardStateLT(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRewardState), v))
	})
}

// RewardStateLTE applies the LTE predicate on the "reward_state" field.
func RewardStateLTE(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRewardState), v))
	})
}

// RewardStateContains applies the Contains predicate on the "reward_state" field.
func RewardStateContains(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRewardState), v))
	})
}

// RewardStateHasPrefix applies the HasPrefix predicate on the "reward_state" field.
func RewardStateHasPrefix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRewardState), v))
	})
}

// RewardStateHasSuffix applies the HasSuffix predicate on the "reward_state" field.
func RewardStateHasSuffix(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRewardState), v))
	})
}

// RewardStateIsNil applies the IsNil predicate on the "reward_state" field.
func RewardStateIsNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRewardState)))
	})
}

// RewardStateNotNil applies the NotNil predicate on the "reward_state" field.
func RewardStateNotNil() predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRewardState)))
	})
}

// RewardStateEqualFold applies the EqualFold predicate on the "reward_state" field.
func RewardStateEqualFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRewardState), v))
	})
}

// RewardStateContainsFold applies the ContainsFold predicate on the "reward_state" field.
func RewardStateContainsFold(v string) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRewardState), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
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
func Not(p predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
