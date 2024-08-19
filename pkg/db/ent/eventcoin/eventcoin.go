// Code generated by ent, DO NOT EDIT.

package eventcoin

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the eventcoin type in the database.
	Label = "event_coin"
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
	// FieldCoinConfigID holds the string denoting the coin_config_id field in the database.
	FieldCoinConfigID = "coin_config_id"
	// FieldCoinValue holds the string denoting the coin_value field in the database.
	FieldCoinValue = "coin_value"
	// FieldCoinPreUsd holds the string denoting the coin_pre_usd field in the database.
	FieldCoinPreUsd = "coin_pre_usd"
	// Table holds the table name of the eventcoin in the database.
	Table = "event_coins"
)

// Columns holds all SQL columns for eventcoin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldEventID,
	FieldCoinConfigID,
	FieldCoinValue,
	FieldCoinPreUsd,
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
	// DefaultCoinConfigID holds the default value on creation for the "coin_config_id" field.
	DefaultCoinConfigID func() uuid.UUID
	// DefaultCoinValue holds the default value on creation for the "coin_value" field.
	DefaultCoinValue decimal.Decimal
	// DefaultCoinPreUsd holds the default value on creation for the "coin_pre_usd" field.
	DefaultCoinPreUsd decimal.Decimal
)