// Code generated by ent, DO NOT EDIT.

package commission

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the commission type in the database.
	Label = "commission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldPercent holds the string denoting the percent field in the database.
	FieldPercent = "percent"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldSettleType holds the string denoting the settle_type field in the database.
	FieldSettleType = "settle_type"
	// FieldSettleMode holds the string denoting the settle_mode field in the database.
	FieldSettleMode = "settle_mode"
	// FieldSettleInterval holds the string denoting the settle_interval field in the database.
	FieldSettleInterval = "settle_interval"
	// FieldThreshold holds the string denoting the threshold field in the database.
	FieldThreshold = "threshold"
	// Table holds the table name of the commission in the database.
	Table = "commissions"
)

// Columns holds all SQL columns for commission fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldPercent,
	FieldStartAt,
	FieldEndAt,
	FieldSettleType,
	FieldSettleMode,
	FieldSettleInterval,
	FieldThreshold,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultPercent holds the default value on creation for the "percent" field.
	DefaultPercent decimal.Decimal
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt uint32
	// DefaultSettleType holds the default value on creation for the "settle_type" field.
	DefaultSettleType string
	// DefaultSettleMode holds the default value on creation for the "settle_mode" field.
	DefaultSettleMode string
	// DefaultSettleInterval holds the default value on creation for the "settle_interval" field.
	DefaultSettleInterval string
	// DefaultThreshold holds the default value on creation for the "threshold" field.
	DefaultThreshold decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
