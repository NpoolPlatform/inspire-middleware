// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coinallocated"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinAllocated is the model entity for the CoinAllocated schema.
type CoinAllocated struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CoinConfigID holds the value of the "coin_config_id" field.
	CoinConfigID uuid.UUID `json:"coin_config_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Value holds the value of the "value" field.
	Value decimal.Decimal `json:"value,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinAllocated) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinallocated.FieldValue:
			values[i] = new(decimal.Decimal)
		case coinallocated.FieldID, coinallocated.FieldCreatedAt, coinallocated.FieldUpdatedAt, coinallocated.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coinallocated.FieldEntID, coinallocated.FieldAppID, coinallocated.FieldCoinConfigID, coinallocated.FieldCoinTypeID, coinallocated.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinAllocated", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinAllocated fields.
func (ca *CoinAllocated) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinallocated.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ca.ID = uint32(value.Int64)
		case coinallocated.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ca.CreatedAt = uint32(value.Int64)
			}
		case coinallocated.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ca.UpdatedAt = uint32(value.Int64)
			}
		case coinallocated.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ca.DeletedAt = uint32(value.Int64)
			}
		case coinallocated.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ca.EntID = *value
			}
		case coinallocated.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ca.AppID = *value
			}
		case coinallocated.FieldCoinConfigID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_config_id", values[i])
			} else if value != nil {
				ca.CoinConfigID = *value
			}
		case coinallocated.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ca.CoinTypeID = *value
			}
		case coinallocated.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ca.UserID = *value
			}
		case coinallocated.FieldValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value != nil {
				ca.Value = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinAllocated.
// Note that you need to call CoinAllocated.Unwrap() before calling this method if this CoinAllocated
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CoinAllocated) Update() *CoinAllocatedUpdateOne {
	return (&CoinAllocatedClient{config: ca.config}).UpdateOne(ca)
}

// Unwrap unwraps the CoinAllocated entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CoinAllocated) Unwrap() *CoinAllocated {
	_tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinAllocated is not a transactional entity")
	}
	ca.config.driver = _tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CoinAllocated) String() string {
	var builder strings.Builder
	builder.WriteString("CoinAllocated(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ca.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ca.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.AppID))
	builder.WriteString(", ")
	builder.WriteString("coin_config_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.CoinConfigID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.UserID))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", ca.Value))
	builder.WriteByte(')')
	return builder.String()
}

// CoinAllocateds is a parsable slice of CoinAllocated.
type CoinAllocateds []*CoinAllocated

func (ca CoinAllocateds) config(cfg config) {
	for _i := range ca {
		ca[_i].config = cfg
	}
}
