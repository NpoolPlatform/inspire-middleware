// Code generated by ent, DO NOT EDIT.

package taskconfig

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the taskconfig type in the database.
	Label = "task_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldTaskType holds the string denoting the task_type field in the database.
	FieldTaskType = "task_type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTaskDesc holds the string denoting the task_desc field in the database.
	FieldTaskDesc = "task_desc"
	// FieldStepGuide holds the string denoting the step_guide field in the database.
	FieldStepGuide = "step_guide"
	// FieldRecommendMessage holds the string denoting the recommend_message field in the database.
	FieldRecommendMessage = "recommend_message"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldLastTaskID holds the string denoting the last_task_id field in the database.
	FieldLastTaskID = "last_task_id"
	// FieldMaxRewardCount holds the string denoting the max_reward_count field in the database.
	FieldMaxRewardCount = "max_reward_count"
	// FieldCooldownSecond holds the string denoting the cooldown_second field in the database.
	FieldCooldownSecond = "cooldown_second"
	// FieldIntervalReset holds the string denoting the interval_reset field in the database.
	FieldIntervalReset = "interval_reset"
	// FieldIntervalResetSecond holds the string denoting the interval_reset_second field in the database.
	FieldIntervalResetSecond = "interval_reset_second"
	// FieldMaxIntervalRewardCount holds the string denoting the max_interval_reward_count field in the database.
	FieldMaxIntervalRewardCount = "max_interval_reward_count"
	// Table holds the table name of the taskconfig in the database.
	Table = "task_configs"
)

// Columns holds all SQL columns for taskconfig fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldEventID,
	FieldTaskType,
	FieldName,
	FieldTaskDesc,
	FieldStepGuide,
	FieldRecommendMessage,
	FieldIndex,
	FieldLastTaskID,
	FieldMaxRewardCount,
	FieldCooldownSecond,
	FieldIntervalReset,
	FieldIntervalResetSecond,
	FieldMaxIntervalRewardCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultEventID holds the default value on creation for the "event_id" field.
	DefaultEventID func() uuid.UUID
	// DefaultTaskType holds the default value on creation for the "task_type" field.
	DefaultTaskType string
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultTaskDesc holds the default value on creation for the "task_desc" field.
	DefaultTaskDesc string
	// DefaultStepGuide holds the default value on creation for the "step_guide" field.
	DefaultStepGuide string
	// DefaultRecommendMessage holds the default value on creation for the "recommend_message" field.
	DefaultRecommendMessage string
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint32
	// DefaultLastTaskID holds the default value on creation for the "last_task_id" field.
	DefaultLastTaskID func() uuid.UUID
	// DefaultMaxRewardCount holds the default value on creation for the "max_reward_count" field.
	DefaultMaxRewardCount uint32
	// DefaultCooldownSecond holds the default value on creation for the "cooldown_second" field.
	DefaultCooldownSecond uint32
	// DefaultIntervalReset holds the default value on creation for the "interval_reset" field.
	DefaultIntervalReset bool
	// DefaultIntervalResetSecond holds the default value on creation for the "interval_reset_second" field.
	DefaultIntervalResetSecond uint32
	// DefaultMaxIntervalRewardCount holds the default value on creation for the "max_interval_reward_count" field.
	DefaultMaxIntervalRewardCount uint32
)
