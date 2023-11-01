package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/mixin"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CouponAllocated holds the schema definition for the CouponAllocated entity.
type CouponAllocated struct {
	ent.Schema
}

func (CouponAllocated) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CouponAllocated.
func (CouponAllocated) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("coupon_id", uuid.UUID{}),
		field.
			Other("denomination", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("used").
			Optional().
			Default(false),
		field.
			Uint32("used_at").
			Optional().
			Default(0),
		field.
			UUID("used_by_order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			String("coupon_scope").
			Optional().
			Default(inspiretypes.CouponScope_Whitelist.String()),
	}
}

// Edges of the CouponAllocated.
func (CouponAllocated) Edges() []ent.Edge {
	return nil
}
