// Code generated by ent, DO NOT EDIT.

package taskuser

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the taskuser type in the database.
	Label = "task_user"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldTaskState holds the string denoting the task_state field in the database.
	FieldTaskState = "task_state"
	// FieldRewardState holds the string denoting the reward_state field in the database.
	FieldRewardState = "reward_state"
	// Table holds the table name of the taskuser in the database.
	Table = "task_users"
)

// Columns holds all SQL columns for taskuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldTaskID,
	FieldEventID,
	FieldTaskState,
	FieldRewardState,
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
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultTaskID holds the default value on creation for the "task_id" field.
	DefaultTaskID func() uuid.UUID
	// DefaultEventID holds the default value on creation for the "event_id" field.
	DefaultEventID func() uuid.UUID
	// DefaultTaskState holds the default value on creation for the "task_state" field.
	DefaultTaskState string
	// DefaultRewardState holds the default value on creation for the "reward_state" field.
	DefaultRewardState string
)
