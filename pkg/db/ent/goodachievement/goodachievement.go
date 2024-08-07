// Code generated by ent, DO NOT EDIT.

package goodachievement

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the goodachievement type in the database.
	Label = "good_achievement"
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
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldTotalUnits holds the string denoting the total_units field in the database.
	FieldTotalUnits = "total_units"
	// FieldSelfUnits holds the string denoting the self_units field in the database.
	FieldSelfUnits = "self_units"
	// FieldTotalAmountUsd holds the string denoting the total_amount_usd field in the database.
	FieldTotalAmountUsd = "total_amount_usd"
	// FieldSelfAmountUsd holds the string denoting the self_amount_usd field in the database.
	FieldSelfAmountUsd = "self_amount_usd"
	// FieldTotalCommissionUsd holds the string denoting the total_commission_usd field in the database.
	FieldTotalCommissionUsd = "total_commission_usd"
	// FieldSelfCommissionUsd holds the string denoting the self_commission_usd field in the database.
	FieldSelfCommissionUsd = "self_commission_usd"
	// Table holds the table name of the goodachievement in the database.
	Table = "good_achievements"
)

// Columns holds all SQL columns for goodachievement fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldAppGoodID,
	FieldTotalUnits,
	FieldSelfUnits,
	FieldTotalAmountUsd,
	FieldSelfAmountUsd,
	FieldTotalCommissionUsd,
	FieldSelfCommissionUsd,
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
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultTotalUnits holds the default value on creation for the "total_units" field.
	DefaultTotalUnits decimal.Decimal
	// DefaultSelfUnits holds the default value on creation for the "self_units" field.
	DefaultSelfUnits decimal.Decimal
	// DefaultTotalAmountUsd holds the default value on creation for the "total_amount_usd" field.
	DefaultTotalAmountUsd decimal.Decimal
	// DefaultSelfAmountUsd holds the default value on creation for the "self_amount_usd" field.
	DefaultSelfAmountUsd decimal.Decimal
	// DefaultTotalCommissionUsd holds the default value on creation for the "total_commission_usd" field.
	DefaultTotalCommissionUsd decimal.Decimal
	// DefaultSelfCommissionUsd holds the default value on creation for the "self_commission_usd" field.
	DefaultSelfCommissionUsd decimal.Decimal
)
