// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodcommissionconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppGoodCommissionConfig is the model entity for the AppGoodCommissionConfig schema.
type AppGoodCommissionConfig struct {
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
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// ThresholdAmount holds the value of the "threshold_amount" field.
	ThresholdAmount decimal.Decimal `json:"threshold_amount,omitempty"`
	// AmountOrPercent holds the value of the "amount_or_percent" field.
	AmountOrPercent decimal.Decimal `json:"amount_or_percent,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt uint32 `json:"end_at,omitempty"`
	// Invites holds the value of the "invites" field.
	Invites uint32 `json:"invites,omitempty"`
	// SettleType holds the value of the "settle_type" field.
	SettleType string `json:"settle_type,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppGoodCommissionConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appgoodcommissionconfig.FieldThresholdAmount, appgoodcommissionconfig.FieldAmountOrPercent:
			values[i] = new(decimal.Decimal)
		case appgoodcommissionconfig.FieldDisabled:
			values[i] = new(sql.NullBool)
		case appgoodcommissionconfig.FieldID, appgoodcommissionconfig.FieldCreatedAt, appgoodcommissionconfig.FieldUpdatedAt, appgoodcommissionconfig.FieldDeletedAt, appgoodcommissionconfig.FieldStartAt, appgoodcommissionconfig.FieldEndAt, appgoodcommissionconfig.FieldInvites:
			values[i] = new(sql.NullInt64)
		case appgoodcommissionconfig.FieldSettleType:
			values[i] = new(sql.NullString)
		case appgoodcommissionconfig.FieldEntID, appgoodcommissionconfig.FieldAppID, appgoodcommissionconfig.FieldGoodID, appgoodcommissionconfig.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppGoodCommissionConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppGoodCommissionConfig fields.
func (agcc *AppGoodCommissionConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appgoodcommissionconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			agcc.ID = uint32(value.Int64)
		case appgoodcommissionconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				agcc.CreatedAt = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				agcc.UpdatedAt = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				agcc.DeletedAt = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				agcc.EntID = *value
			}
		case appgoodcommissionconfig.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				agcc.AppID = *value
			}
		case appgoodcommissionconfig.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				agcc.GoodID = *value
			}
		case appgoodcommissionconfig.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				agcc.AppGoodID = *value
			}
		case appgoodcommissionconfig.FieldThresholdAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field threshold_amount", values[i])
			} else if value != nil {
				agcc.ThresholdAmount = *value
			}
		case appgoodcommissionconfig.FieldAmountOrPercent:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount_or_percent", values[i])
			} else if value != nil {
				agcc.AmountOrPercent = *value
			}
		case appgoodcommissionconfig.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				agcc.StartAt = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				agcc.EndAt = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldInvites:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invites", values[i])
			} else if value.Valid {
				agcc.Invites = uint32(value.Int64)
			}
		case appgoodcommissionconfig.FieldSettleType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settle_type", values[i])
			} else if value.Valid {
				agcc.SettleType = value.String
			}
		case appgoodcommissionconfig.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				agcc.Disabled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppGoodCommissionConfig.
// Note that you need to call AppGoodCommissionConfig.Unwrap() before calling this method if this AppGoodCommissionConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (agcc *AppGoodCommissionConfig) Update() *AppGoodCommissionConfigUpdateOne {
	return (&AppGoodCommissionConfigClient{config: agcc.config}).UpdateOne(agcc)
}

// Unwrap unwraps the AppGoodCommissionConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (agcc *AppGoodCommissionConfig) Unwrap() *AppGoodCommissionConfig {
	_tx, ok := agcc.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppGoodCommissionConfig is not a transactional entity")
	}
	agcc.config.driver = _tx.drv
	return agcc
}

// String implements the fmt.Stringer.
func (agcc *AppGoodCommissionConfig) String() string {
	var builder strings.Builder
	builder.WriteString("AppGoodCommissionConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", agcc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", agcc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", agcc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", agcc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", agcc.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", agcc.AppID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", agcc.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", agcc.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("threshold_amount=")
	builder.WriteString(fmt.Sprintf("%v", agcc.ThresholdAmount))
	builder.WriteString(", ")
	builder.WriteString("amount_or_percent=")
	builder.WriteString(fmt.Sprintf("%v", agcc.AmountOrPercent))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", agcc.StartAt))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(fmt.Sprintf("%v", agcc.EndAt))
	builder.WriteString(", ")
	builder.WriteString("invites=")
	builder.WriteString(fmt.Sprintf("%v", agcc.Invites))
	builder.WriteString(", ")
	builder.WriteString("settle_type=")
	builder.WriteString(agcc.SettleType)
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", agcc.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// AppGoodCommissionConfigs is a parsable slice of AppGoodCommissionConfig.
type AppGoodCommissionConfigs []*AppGoodCommissionConfig

func (agcc AppGoodCommissionConfigs) config(cfg config) {
	for _i := range agcc {
		agcc[_i].config = cfg
	}
}
