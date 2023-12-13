package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/shopspring/decimal"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/google/uuid"
)

// Coupon holds the schema definition for the Coupon entity.
type Coupon struct {
	ent.Schema
}

func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("denomination", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("circulation", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("random").
			Optional().
			Default(false),
		field.
			UUID("issued_by", uuid.UUID{}),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			Uint32("duration_days").
			Optional().
			Default(365), //nolint
		field.
			String("message").
			Optional().
			Default(""),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			Other("allocated", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("coupon_type").
			Optional().
			Default(types.CouponType_DefaultCouponType.String()),
		field.
			Other("threshold", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("coupon_constraint").
			Optional().
			Default(types.CouponConstraint_Normal.String()),
		field.
			String("coupon_scope").
			Optional().
			Default(types.CouponScope_Whitelist.String()),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return nil
}
