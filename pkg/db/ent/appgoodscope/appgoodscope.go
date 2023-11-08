// Code generated by ent, DO NOT EDIT.

package appgoodscope

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appgoodscope type in the database.
	Label = "app_good_scope"
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
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldCouponID holds the string denoting the coupon_id field in the database.
	FieldCouponID = "coupon_id"
	// FieldCouponScope holds the string denoting the coupon_scope field in the database.
	FieldCouponScope = "coupon_scope"
	// Table holds the table name of the appgoodscope in the database.
	Table = "app_good_scopes"
)

// Columns holds all SQL columns for appgoodscope fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldAppGoodID,
	FieldCouponID,
	FieldCouponScope,
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
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultCouponID holds the default value on creation for the "coupon_id" field.
	DefaultCouponID func() uuid.UUID
	// DefaultCouponScope holds the default value on creation for the "coupon_scope" field.
	DefaultCouponScope string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
