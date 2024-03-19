package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppCommissionConfig holds the schema definition for the AppCommissionConfig entity.
type AppCommissionConfig struct {
	ent.Schema
}

func (AppCommissionConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppCommissionConfig.
func (AppCommissionConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("threshold_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("amount_or_percent", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			Uint32("invites").
			Optional().
			Default(0),
		field.
			String("settle_type").
			Optional().
			Default(types.SettleType_DefaultSettleType.String()),
	}
}

// Edges of the AppCommissionConfig.
func (AppCommissionConfig) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppCommissionConfig.
func (AppCommissionConfig) Indexes() []ent.Index {
	return nil
}
