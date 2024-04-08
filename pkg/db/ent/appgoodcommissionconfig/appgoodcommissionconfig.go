// Code generated by ent, DO NOT EDIT.

package appgoodcommissionconfig

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the appgoodcommissionconfig type in the database.
	Label = "app_good_commission_config"
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
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldThresholdAmount holds the string denoting the threshold_amount field in the database.
	FieldThresholdAmount = "threshold_amount"
	// FieldAmountOrPercent holds the string denoting the amount_or_percent field in the database.
	FieldAmountOrPercent = "amount_or_percent"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldInvites holds the string denoting the invites field in the database.
	FieldInvites = "invites"
	// FieldSettleType holds the string denoting the settle_type field in the database.
	FieldSettleType = "settle_type"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// Table holds the table name of the appgoodcommissionconfig in the database.
	Table = "app_good_commission_configs"
)

// Columns holds all SQL columns for appgoodcommissionconfig fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldGoodID,
	FieldAppGoodID,
	FieldThresholdAmount,
	FieldAmountOrPercent,
	FieldStartAt,
	FieldEndAt,
	FieldInvites,
	FieldSettleType,
	FieldDisabled,
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
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultThresholdAmount holds the default value on creation for the "threshold_amount" field.
	DefaultThresholdAmount decimal.Decimal
	// DefaultAmountOrPercent holds the default value on creation for the "amount_or_percent" field.
	DefaultAmountOrPercent decimal.Decimal
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt uint32
	// DefaultInvites holds the default value on creation for the "invites" field.
	DefaultInvites uint32
	// DefaultSettleType holds the default value on creation for the "settle_type" field.
	DefaultSettleType string
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
)
