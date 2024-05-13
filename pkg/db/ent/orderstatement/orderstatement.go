// Code generated by ent, DO NOT EDIT.

package orderstatement

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the orderstatement type in the database.
	Label = "order_statement"
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
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldOrderUserID holds the string denoting the order_user_id field in the database.
	FieldOrderUserID = "order_user_id"
	// FieldSelfOrder holds the string denoting the self_order field in the database.
	FieldSelfOrder = "self_order"
	// FieldGoodCoinTypeID holds the string denoting the good_coin_type_id field in the database.
	FieldGoodCoinTypeID = "good_coin_type_id"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldGoodValueUsd holds the string denoting the good_value_usd field in the database.
	FieldGoodValueUsd = "good_value_usd"
	// FieldPaymentAmountUsd holds the string denoting the payment_amount_usd field in the database.
	FieldPaymentAmountUsd = "payment_amount_usd"
	// FieldCommissionAmountUsd holds the string denoting the commission_amount_usd field in the database.
	FieldCommissionAmountUsd = "commission_amount_usd"
	// FieldAppConfigID holds the string denoting the app_config_id field in the database.
	FieldAppConfigID = "app_config_id"
	// FieldCommissionConfigID holds the string denoting the commission_config_id field in the database.
	FieldCommissionConfigID = "commission_config_id"
	// FieldCommissionConfigType holds the string denoting the commission_config_type field in the database.
	FieldCommissionConfigType = "commission_config_type"
	// Table holds the table name of the orderstatement in the database.
	Table = "order_statements"
)

// Columns holds all SQL columns for orderstatement fields.
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
	FieldOrderID,
	FieldOrderUserID,
	FieldSelfOrder,
	FieldGoodCoinTypeID,
	FieldUnits,
	FieldGoodValueUsd,
	FieldPaymentAmountUsd,
	FieldCommissionAmountUsd,
	FieldAppConfigID,
	FieldCommissionConfigID,
	FieldCommissionConfigType,
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
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultOrderUserID holds the default value on creation for the "order_user_id" field.
	DefaultOrderUserID func() uuid.UUID
	// DefaultSelfOrder holds the default value on creation for the "self_order" field.
	DefaultSelfOrder bool
	// DefaultGoodCoinTypeID holds the default value on creation for the "good_coin_type_id" field.
	DefaultGoodCoinTypeID func() uuid.UUID
	// DefaultUnits holds the default value on creation for the "units" field.
	DefaultUnits uint32
	// DefaultGoodValueUsd holds the default value on creation for the "good_value_usd" field.
	DefaultGoodValueUsd decimal.Decimal
	// DefaultPaymentAmountUsd holds the default value on creation for the "payment_amount_usd" field.
	DefaultPaymentAmountUsd decimal.Decimal
	// DefaultCommissionAmountUsd holds the default value on creation for the "commission_amount_usd" field.
	DefaultCommissionAmountUsd decimal.Decimal
	// DefaultAppConfigID holds the default value on creation for the "app_config_id" field.
	DefaultAppConfigID func() uuid.UUID
	// DefaultCommissionConfigID holds the default value on creation for the "commission_config_id" field.
	DefaultCommissionConfigID func() uuid.UUID
	// DefaultCommissionConfigType holds the default value on creation for the "commission_config_type" field.
	DefaultCommissionConfigType string
)
