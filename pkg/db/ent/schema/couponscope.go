package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
)

// CouponScope holds the schema definition for the CouponScope entity.
type CouponScope struct {
	ent.Schema
}

func (CouponScope) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CouponScope.
func (CouponScope) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("coupon_scope").
			Optional().
			Default(inspiretypes.CouponScope_Whitelist.String()),
	}
}

// Edges of the CouponScope.
func (CouponScope) Edges() []ent.Edge {
	return nil
}
