// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Coupon is the model entity for the Coupon schema.
type Coupon struct {
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
	// Denomination holds the value of the "denomination" field.
	Denomination decimal.Decimal `json:"denomination,omitempty"`
	// Circulation holds the value of the "circulation" field.
	Circulation decimal.Decimal `json:"circulation,omitempty"`
	// Random holds the value of the "random" field.
	Random bool `json:"random,omitempty"`
	// IssuedBy holds the value of the "issued_by" field.
	IssuedBy uuid.UUID `json:"issued_by,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// DurationDays holds the value of the "duration_days" field.
	DurationDays uint32 `json:"duration_days,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Allocated holds the value of the "allocated" field.
	Allocated decimal.Decimal `json:"allocated,omitempty"`
	// CouponType holds the value of the "coupon_type" field.
	CouponType string `json:"coupon_type,omitempty"`
	// Threshold holds the value of the "threshold" field.
	Threshold decimal.Decimal `json:"threshold,omitempty"`
	// CouponConstraint holds the value of the "coupon_constraint" field.
	CouponConstraint string `json:"coupon_constraint,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Coupon) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coupon.FieldDenomination, coupon.FieldCirculation, coupon.FieldAllocated, coupon.FieldThreshold:
			values[i] = new(decimal.Decimal)
		case coupon.FieldRandom:
			values[i] = new(sql.NullBool)
		case coupon.FieldCreatedAt, coupon.FieldUpdatedAt, coupon.FieldDeletedAt, coupon.FieldStartAt, coupon.FieldDurationDays:
			values[i] = new(sql.NullInt64)
		case coupon.FieldMessage, coupon.FieldName, coupon.FieldCouponType, coupon.FieldCouponConstraint:
			values[i] = new(sql.NullString)
		case coupon.FieldID, coupon.FieldAppID, coupon.FieldUserID, coupon.FieldGoodID, coupon.FieldIssuedBy:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Coupon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Coupon fields.
func (c *Coupon) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coupon.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case coupon.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = uint32(value.Int64)
			}
		case coupon.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = uint32(value.Int64)
			}
		case coupon.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = uint32(value.Int64)
			}
		case coupon.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				c.AppID = *value
			}
		case coupon.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				c.UserID = *value
			}
		case coupon.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				c.GoodID = *value
			}
		case coupon.FieldDenomination:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field denomination", values[i])
			} else if value != nil {
				c.Denomination = *value
			}
		case coupon.FieldCirculation:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field circulation", values[i])
			} else if value != nil {
				c.Circulation = *value
			}
		case coupon.FieldRandom:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field random", values[i])
			} else if value.Valid {
				c.Random = value.Bool
			}
		case coupon.FieldIssuedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field issued_by", values[i])
			} else if value != nil {
				c.IssuedBy = *value
			}
		case coupon.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				c.StartAt = uint32(value.Int64)
			}
		case coupon.FieldDurationDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_days", values[i])
			} else if value.Valid {
				c.DurationDays = uint32(value.Int64)
			}
		case coupon.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				c.Message = value.String
			}
		case coupon.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case coupon.FieldAllocated:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field allocated", values[i])
			} else if value != nil {
				c.Allocated = *value
			}
		case coupon.FieldCouponType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_type", values[i])
			} else if value.Valid {
				c.CouponType = value.String
			}
		case coupon.FieldThreshold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field threshold", values[i])
			} else if value != nil {
				c.Threshold = *value
			}
		case coupon.FieldCouponConstraint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_constraint", values[i])
			} else if value.Valid {
				c.CouponConstraint = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Coupon.
// Note that you need to call Coupon.Unwrap() before calling this method if this Coupon
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Coupon) Update() *CouponUpdateOne {
	return (&CouponClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Coupon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Coupon) Unwrap() *Coupon {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Coupon is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Coupon) String() string {
	var builder strings.Builder
	builder.WriteString("Coupon(")
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
	builder.WriteString("denomination=")
	builder.WriteString(fmt.Sprintf("%v", c.Denomination))
	builder.WriteString(", ")
	builder.WriteString("circulation=")
	builder.WriteString(fmt.Sprintf("%v", c.Circulation))
	builder.WriteString(", ")
	builder.WriteString("random=")
	builder.WriteString(fmt.Sprintf("%v", c.Random))
	builder.WriteString(", ")
	builder.WriteString("issued_by=")
	builder.WriteString(fmt.Sprintf("%v", c.IssuedBy))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", c.StartAt))
	builder.WriteString(", ")
	builder.WriteString("duration_days=")
	builder.WriteString(fmt.Sprintf("%v", c.DurationDays))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(c.Message)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("allocated=")
	builder.WriteString(fmt.Sprintf("%v", c.Allocated))
	builder.WriteString(", ")
	builder.WriteString("coupon_type=")
	builder.WriteString(c.CouponType)
	builder.WriteString(", ")
	builder.WriteString("threshold=")
	builder.WriteString(fmt.Sprintf("%v", c.Threshold))
	builder.WriteString(", ")
	builder.WriteString("coupon_constraint=")
	builder.WriteString(c.CouponConstraint)
	builder.WriteByte(')')
	return builder.String()
}

// Coupons is a parsable slice of Coupon.
type Coupons []*Coupon

func (c Coupons) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}