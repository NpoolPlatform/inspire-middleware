// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/usercoinreward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UserCoinReward is the model entity for the UserCoinReward schema.
type UserCoinReward struct {
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
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// CoinRewards holds the value of the "coin_rewards" field.
	CoinRewards decimal.Decimal `json:"coin_rewards,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCoinReward) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercoinreward.FieldCoinRewards:
			values[i] = new(decimal.Decimal)
		case usercoinreward.FieldID, usercoinreward.FieldCreatedAt, usercoinreward.FieldUpdatedAt, usercoinreward.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case usercoinreward.FieldEntID, usercoinreward.FieldAppID, usercoinreward.FieldUserID, usercoinreward.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserCoinReward", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCoinReward fields.
func (ucr *UserCoinReward) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercoinreward.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ucr.ID = uint32(value.Int64)
		case usercoinreward.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ucr.CreatedAt = uint32(value.Int64)
			}
		case usercoinreward.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ucr.UpdatedAt = uint32(value.Int64)
			}
		case usercoinreward.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ucr.DeletedAt = uint32(value.Int64)
			}
		case usercoinreward.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ucr.EntID = *value
			}
		case usercoinreward.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ucr.AppID = *value
			}
		case usercoinreward.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ucr.UserID = *value
			}
		case usercoinreward.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ucr.CoinTypeID = *value
			}
		case usercoinreward.FieldCoinRewards:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field coin_rewards", values[i])
			} else if value != nil {
				ucr.CoinRewards = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserCoinReward.
// Note that you need to call UserCoinReward.Unwrap() before calling this method if this UserCoinReward
// was returned from a transaction, and the transaction was committed or rolled back.
func (ucr *UserCoinReward) Update() *UserCoinRewardUpdateOne {
	return (&UserCoinRewardClient{config: ucr.config}).UpdateOne(ucr)
}

// Unwrap unwraps the UserCoinReward entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ucr *UserCoinReward) Unwrap() *UserCoinReward {
	_tx, ok := ucr.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserCoinReward is not a transactional entity")
	}
	ucr.config.driver = _tx.drv
	return ucr
}

// String implements the fmt.Stringer.
func (ucr *UserCoinReward) String() string {
	var builder strings.Builder
	builder.WriteString("UserCoinReward(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ucr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ucr.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ucr.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ucr.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ucr.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ucr.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ucr.UserID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ucr.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("coin_rewards=")
	builder.WriteString(fmt.Sprintf("%v", ucr.CoinRewards))
	builder.WriteByte(')')
	return builder.String()
}

// UserCoinRewards is a parsable slice of UserCoinReward.
type UserCoinRewards []*UserCoinReward

func (ucr UserCoinRewards) config(cfg config) {
	for _i := range ucr {
		ucr[_i].config = cfg
	}
}