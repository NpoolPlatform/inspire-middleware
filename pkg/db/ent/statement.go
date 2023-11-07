// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Statement is the model entity for the Statement schema.
type Statement struct {
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
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// DirectContributorID holds the value of the "direct_contributor_id" field.
	DirectContributorID uuid.UUID `json:"direct_contributor_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// SelfOrder holds the value of the "self_order" field.
	SelfOrder bool `json:"self_order,omitempty"`
	// PaymentID holds the value of the "payment_id" field.
	PaymentID uuid.UUID `json:"payment_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// PaymentCoinTypeID holds the value of the "payment_coin_type_id" field.
	PaymentCoinTypeID uuid.UUID `json:"payment_coin_type_id,omitempty"`
	// PaymentCoinUsdCurrency holds the value of the "payment_coin_usd_currency" field.
	PaymentCoinUsdCurrency decimal.Decimal `json:"payment_coin_usd_currency,omitempty"`
	// Units holds the value of the "units" field.
	Units uint32 `json:"units,omitempty"`
	// UnitsV1 holds the value of the "units_v1" field.
	UnitsV1 decimal.Decimal `json:"units_v1,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// UsdAmount holds the value of the "usd_amount" field.
	UsdAmount decimal.Decimal `json:"usd_amount,omitempty"`
	// Commission holds the value of the "commission" field.
	Commission decimal.Decimal `json:"commission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Statement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case statement.FieldPaymentCoinUsdCurrency, statement.FieldUnitsV1, statement.FieldAmount, statement.FieldUsdAmount, statement.FieldCommission:
			values[i] = new(decimal.Decimal)
		case statement.FieldSelfOrder:
			values[i] = new(sql.NullBool)
		case statement.FieldID, statement.FieldCreatedAt, statement.FieldUpdatedAt, statement.FieldDeletedAt, statement.FieldUnits:
			values[i] = new(sql.NullInt64)
		case statement.FieldEntID, statement.FieldAppID, statement.FieldUserID, statement.FieldDirectContributorID, statement.FieldGoodID, statement.FieldAppGoodID, statement.FieldOrderID, statement.FieldPaymentID, statement.FieldCoinTypeID, statement.FieldPaymentCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Statement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Statement fields.
func (s *Statement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint32(value.Int64)
		case statement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = uint32(value.Int64)
			}
		case statement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = uint32(value.Int64)
			}
		case statement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = uint32(value.Int64)
			}
		case statement.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				s.EntID = *value
			}
		case statement.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				s.AppID = *value
			}
		case statement.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				s.UserID = *value
			}
		case statement.FieldDirectContributorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field direct_contributor_id", values[i])
			} else if value != nil {
				s.DirectContributorID = *value
			}
		case statement.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				s.GoodID = *value
			}
		case statement.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				s.AppGoodID = *value
			}
		case statement.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				s.OrderID = *value
			}
		case statement.FieldSelfOrder:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field self_order", values[i])
			} else if value.Valid {
				s.SelfOrder = value.Bool
			}
		case statement.FieldPaymentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value != nil {
				s.PaymentID = *value
			}
		case statement.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				s.CoinTypeID = *value
			}
		case statement.FieldPaymentCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_type_id", values[i])
			} else if value != nil {
				s.PaymentCoinTypeID = *value
			}
		case statement.FieldPaymentCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_usd_currency", values[i])
			} else if value != nil {
				s.PaymentCoinUsdCurrency = *value
			}
		case statement.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				s.Units = uint32(value.Int64)
			}
		case statement.FieldUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field units_v1", values[i])
			} else if value != nil {
				s.UnitsV1 = *value
			}
		case statement.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				s.Amount = *value
			}
		case statement.FieldUsdAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field usd_amount", values[i])
			} else if value != nil {
				s.UsdAmount = *value
			}
		case statement.FieldCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field commission", values[i])
			} else if value != nil {
				s.Commission = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Statement.
// Note that you need to call Statement.Unwrap() before calling this method if this Statement
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Statement) Update() *StatementUpdateOne {
	return (&StatementClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Statement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Statement) Unwrap() *Statement {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Statement is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Statement) String() string {
	var builder strings.Builder
	builder.WriteString("Statement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", s.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", s.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", s.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("direct_contributor_id=")
	builder.WriteString(fmt.Sprintf("%v", s.DirectContributorID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", s.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", s.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", s.OrderID))
	builder.WriteString(", ")
	builder.WriteString("self_order=")
	builder.WriteString(fmt.Sprintf("%v", s.SelfOrder))
	builder.WriteString(", ")
	builder.WriteString("payment_id=")
	builder.WriteString(fmt.Sprintf("%v", s.PaymentID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("payment_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.PaymentCoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("payment_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", s.PaymentCoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", s.Units))
	builder.WriteString(", ")
	builder.WriteString("units_v1=")
	builder.WriteString(fmt.Sprintf("%v", s.UnitsV1))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", s.Amount))
	builder.WriteString(", ")
	builder.WriteString("usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.UsdAmount))
	builder.WriteString(", ")
	builder.WriteString("commission=")
	builder.WriteString(fmt.Sprintf("%v", s.Commission))
	builder.WriteByte(')')
	return builder.String()
}

// Statements is a parsable slice of Statement.
type Statements []*Statement

func (s Statements) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
