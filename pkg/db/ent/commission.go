// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Commission is the model entity for the Commission schema.
type Commission struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// AmountOrPercent holds the value of the "amount_or_percent" field.
	AmountOrPercent decimal.Decimal `json:"amount_or_percent,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt uint32 `json:"end_at,omitempty"`
	// SettleType holds the value of the "settle_type" field.
	SettleType string `json:"settle_type,omitempty"`
	// SettleMode holds the value of the "settle_mode" field.
	SettleMode string `json:"settle_mode,omitempty"`
	// SettleInterval holds the value of the "settle_interval" field.
	SettleInterval string `json:"settle_interval,omitempty"`
	// SettleAmountType holds the value of the "settle_amount_type" field.
	SettleAmountType string `json:"settle_amount_type,omitempty"`
	// Threshold holds the value of the "threshold" field.
	Threshold decimal.Decimal `json:"threshold,omitempty"`
	// OrderLimit holds the value of the "order_limit" field.
	OrderLimit uint32 `json:"order_limit,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Commission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case commission.FieldAmountOrPercent, commission.FieldThreshold:
			values[i] = new(decimal.Decimal)
		case commission.FieldCreatedAt, commission.FieldUpdatedAt, commission.FieldDeletedAt, commission.FieldStartAt, commission.FieldEndAt, commission.FieldOrderLimit:
			values[i] = new(sql.NullInt64)
		case commission.FieldSettleType, commission.FieldSettleMode, commission.FieldSettleInterval, commission.FieldSettleAmountType:
			values[i] = new(sql.NullString)
		case commission.FieldID, commission.FieldAppID, commission.FieldUserID, commission.FieldGoodID, commission.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Commission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Commission fields.
func (c *Commission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commission.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case commission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = uint32(value.Int64)
			}
		case commission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = uint32(value.Int64)
			}
		case commission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = uint32(value.Int64)
			}
		case commission.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				c.AppID = *value
			}
		case commission.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				c.UserID = *value
			}
		case commission.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				c.GoodID = *value
			}
		case commission.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				c.AppGoodID = *value
			}
		case commission.FieldAmountOrPercent:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount_or_percent", values[i])
			} else if value != nil {
				c.AmountOrPercent = *value
			}
		case commission.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				c.StartAt = uint32(value.Int64)
			}
		case commission.FieldEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				c.EndAt = uint32(value.Int64)
			}
		case commission.FieldSettleType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settle_type", values[i])
			} else if value.Valid {
				c.SettleType = value.String
			}
		case commission.FieldSettleMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settle_mode", values[i])
			} else if value.Valid {
				c.SettleMode = value.String
			}
		case commission.FieldSettleInterval:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settle_interval", values[i])
			} else if value.Valid {
				c.SettleInterval = value.String
			}
		case commission.FieldSettleAmountType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settle_amount_type", values[i])
			} else if value.Valid {
				c.SettleAmountType = value.String
			}
		case commission.FieldThreshold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field threshold", values[i])
			} else if value != nil {
				c.Threshold = *value
			}
		case commission.FieldOrderLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_limit", values[i])
			} else if value.Valid {
				c.OrderLimit = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Commission.
// Note that you need to call Commission.Unwrap() before calling this method if this Commission
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Commission) Update() *CommissionUpdateOne {
	return (&CommissionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Commission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Commission) Unwrap() *Commission {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Commission is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Commission) String() string {
	var builder strings.Builder
	builder.WriteString("Commission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", c.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", c.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("amount_or_percent=")
	builder.WriteString(fmt.Sprintf("%v", c.AmountOrPercent))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", c.StartAt))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(fmt.Sprintf("%v", c.EndAt))
	builder.WriteString(", ")
	builder.WriteString("settle_type=")
	builder.WriteString(c.SettleType)
	builder.WriteString(", ")
	builder.WriteString("settle_mode=")
	builder.WriteString(c.SettleMode)
	builder.WriteString(", ")
	builder.WriteString("settle_interval=")
	builder.WriteString(c.SettleInterval)
	builder.WriteString(", ")
	builder.WriteString("settle_amount_type=")
	builder.WriteString(c.SettleAmountType)
	builder.WriteString(", ")
	builder.WriteString("threshold=")
	builder.WriteString(fmt.Sprintf("%v", c.Threshold))
	builder.WriteString(", ")
	builder.WriteString("order_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.OrderLimit))
	builder.WriteByte(')')
	return builder.String()
}

// Commissions is a parsable slice of Commission.
type Commissions []*Commission

func (c Commissions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
