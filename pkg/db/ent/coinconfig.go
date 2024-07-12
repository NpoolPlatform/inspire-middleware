// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coinconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinConfig is the model entity for the CoinConfig schema.
type CoinConfig struct {
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
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// MaxValue holds the value of the "max_value" field.
	MaxValue decimal.Decimal `json:"max_value,omitempty"`
	// Allocated holds the value of the "allocated" field.
	Allocated decimal.Decimal `json:"allocated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinconfig.FieldMaxValue, coinconfig.FieldAllocated:
			values[i] = new(decimal.Decimal)
		case coinconfig.FieldID, coinconfig.FieldCreatedAt, coinconfig.FieldUpdatedAt, coinconfig.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coinconfig.FieldEntID, coinconfig.FieldAppID, coinconfig.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinConfig fields.
func (cc *CoinConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = uint32(value.Int64)
		case coinconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = uint32(value.Int64)
			}
		case coinconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = uint32(value.Int64)
			}
		case coinconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cc.DeletedAt = uint32(value.Int64)
			}
		case coinconfig.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				cc.EntID = *value
			}
		case coinconfig.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				cc.AppID = *value
			}
		case coinconfig.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cc.CoinTypeID = *value
			}
		case coinconfig.FieldMaxValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field max_value", values[i])
			} else if value != nil {
				cc.MaxValue = *value
			}
		case coinconfig.FieldAllocated:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field allocated", values[i])
			} else if value != nil {
				cc.Allocated = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinConfig.
// Note that you need to call CoinConfig.Unwrap() before calling this method if this CoinConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CoinConfig) Update() *CoinConfigUpdateOne {
	return (&CoinConfigClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the CoinConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CoinConfig) Unwrap() *CoinConfig {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinConfig is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CoinConfig) String() string {
	var builder strings.Builder
	builder.WriteString("CoinConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.AppID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("max_value=")
	builder.WriteString(fmt.Sprintf("%v", cc.MaxValue))
	builder.WriteString(", ")
	builder.WriteString("allocated=")
	builder.WriteString(fmt.Sprintf("%v", cc.Allocated))
	builder.WriteByte(')')
	return builder.String()
}

// CoinConfigs is a parsable slice of CoinConfig.
type CoinConfigs []*CoinConfig

func (cc CoinConfigs) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}
