package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"

	"github.com/google/uuid"
)

// InvitationCode holds the schema definition for the InvitationCode entity.
type InvitationCode struct {
	ent.Schema
}

func (InvitationCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the InvitationCode.
func (InvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("invitation_code").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the InvitationCode.
func (InvitationCode) Edges() []ent.Edge {
	return nil
}
