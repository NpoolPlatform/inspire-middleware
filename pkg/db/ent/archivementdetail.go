// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/archivementdetail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ArchivementDetail is the model entity for the ArchivementDetail schema.
type ArchivementDetail struct {
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
	// DirectContributorID holds the value of the "direct_contributor_id" field.
	DirectContributorID uuid.UUID `json:"direct_contributor_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
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
func (*ArchivementDetail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case archivementdetail.FieldPaymentCoinUsdCurrency, archivementdetail.FieldUnitsV1, archivementdetail.FieldAmount, archivementdetail.FieldUsdAmount, archivementdetail.FieldCommission:
			values[i] = new(decimal.Decimal)
		case archivementdetail.FieldSelfOrder:
			values[i] = new(sql.NullBool)
		case archivementdetail.FieldCreatedAt, archivementdetail.FieldUpdatedAt, archivementdetail.FieldDeletedAt, archivementdetail.FieldUnits:
			values[i] = new(sql.NullInt64)
		case archivementdetail.FieldID, archivementdetail.FieldAppID, archivementdetail.FieldUserID, archivementdetail.FieldDirectContributorID, archivementdetail.FieldGoodID, archivementdetail.FieldOrderID, archivementdetail.FieldPaymentID, archivementdetail.FieldCoinTypeID, archivementdetail.FieldPaymentCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ArchivementDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArchivementDetail fields.
func (ad *ArchivementDetail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case archivementdetail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ad.ID = *value
			}
		case archivementdetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ad.CreatedAt = uint32(value.Int64)
			}
		case archivementdetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ad.UpdatedAt = uint32(value.Int64)
			}
		case archivementdetail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ad.DeletedAt = uint32(value.Int64)
			}
		case archivementdetail.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ad.AppID = *value
			}
		case archivementdetail.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ad.UserID = *value
			}
		case archivementdetail.FieldDirectContributorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field direct_contributor_id", values[i])
			} else if value != nil {
				ad.DirectContributorID = *value
			}
		case archivementdetail.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ad.GoodID = *value
			}
		case archivementdetail.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				ad.OrderID = *value
			}
		case archivementdetail.FieldSelfOrder:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field self_order", values[i])
			} else if value.Valid {
				ad.SelfOrder = value.Bool
			}
		case archivementdetail.FieldPaymentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value != nil {
				ad.PaymentID = *value
			}
		case archivementdetail.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ad.CoinTypeID = *value
			}
		case archivementdetail.FieldPaymentCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_type_id", values[i])
			} else if value != nil {
				ad.PaymentCoinTypeID = *value
			}
		case archivementdetail.FieldPaymentCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_usd_currency", values[i])
			} else if value != nil {
				ad.PaymentCoinUsdCurrency = *value
			}
		case archivementdetail.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				ad.Units = uint32(value.Int64)
			}
		case archivementdetail.FieldUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field units_v1", values[i])
			} else if value != nil {
				ad.UnitsV1 = *value
			}
		case archivementdetail.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				ad.Amount = *value
			}
		case archivementdetail.FieldUsdAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field usd_amount", values[i])
			} else if value != nil {
				ad.UsdAmount = *value
			}
		case archivementdetail.FieldCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field commission", values[i])
			} else if value != nil {
				ad.Commission = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ArchivementDetail.
// Note that you need to call ArchivementDetail.Unwrap() before calling this method if this ArchivementDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (ad *ArchivementDetail) Update() *ArchivementDetailUpdateOne {
	return (&ArchivementDetailClient{config: ad.config}).UpdateOne(ad)
}

// Unwrap unwraps the ArchivementDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ad *ArchivementDetail) Unwrap() *ArchivementDetail {
	_tx, ok := ad.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArchivementDetail is not a transactional entity")
	}
	ad.config.driver = _tx.drv
	return ad
}

// String implements the fmt.Stringer.
func (ad *ArchivementDetail) String() string {
	var builder strings.Builder
	builder.WriteString("ArchivementDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ad.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ad.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ad.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ad.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.UserID))
	builder.WriteString(", ")
	builder.WriteString("direct_contributor_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.DirectContributorID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.GoodID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.OrderID))
	builder.WriteString(", ")
	builder.WriteString("self_order=")
	builder.WriteString(fmt.Sprintf("%v", ad.SelfOrder))
	builder.WriteString(", ")
	builder.WriteString("payment_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.PaymentID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("payment_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.PaymentCoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("payment_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", ad.PaymentCoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", ad.Units))
	builder.WriteString(", ")
	builder.WriteString("units_v1=")
	builder.WriteString(fmt.Sprintf("%v", ad.UnitsV1))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", ad.Amount))
	builder.WriteString(", ")
	builder.WriteString("usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", ad.UsdAmount))
	builder.WriteString(", ")
	builder.WriteString("commission=")
	builder.WriteString(fmt.Sprintf("%v", ad.Commission))
	builder.WriteByte(')')
	return builder.String()
}

// ArchivementDetails is a parsable slice of ArchivementDetail.
type ArchivementDetails []*ArchivementDetail

func (ad ArchivementDetails) config(cfg config) {
	for _i := range ad {
		ad[_i].config = cfg
	}
}
