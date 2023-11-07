// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArchivementGeneralsColumns holds the columns for the "archivement_generals" table.
	ArchivementGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "total_units_v1", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_units_v1", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "total_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "total_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "self_commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ArchivementGeneralsTable holds the schema information for the "archivement_generals" table.
	ArchivementGeneralsTable = &schema.Table{
		Name:       "archivement_generals",
		Columns:    ArchivementGeneralsColumns,
		PrimaryKey: []*schema.Column{ArchivementGeneralsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "achievement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ArchivementGeneralsColumns[4]},
			},
		},
	}
	// CommissionsColumns holds the columns for the "commissions" table.
	CommissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount_or_percent", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 1699346837},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "settle_type", Type: field.TypeString, Nullable: true, Default: "DefaultSettleType"},
		{Name: "settle_mode", Type: field.TypeString, Nullable: true, Default: "DefaultSettleMode"},
		{Name: "settle_interval", Type: field.TypeString, Nullable: true, Default: "DefaultSettleInterval"},
		{Name: "settle_amount_type", Type: field.TypeString, Nullable: true, Default: "SettleByPercent"},
		{Name: "threshold", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "order_limit", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// CommissionsTable holds the schema information for the "commissions" table.
	CommissionsTable = &schema.Table{
		Name:       "commissions",
		Columns:    CommissionsColumns,
		PrimaryKey: []*schema.Column{CommissionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "commission_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CommissionsColumns[4]},
			},
		},
	}
	// CouponsColumns holds the columns for the "coupons" table.
	CouponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "denomination", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "circulation", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "random", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "issued_by", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 1699346837},
		{Name: "duration_days", Type: field.TypeUint32, Nullable: true, Default: 365},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "allocated", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "coupon_type", Type: field.TypeString, Nullable: true, Default: "DefaultCouponType"},
		{Name: "threshold", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "coupon_constraint", Type: field.TypeString, Nullable: true, Default: "Normal"},
	}
	// CouponsTable holds the schema information for the "coupons" table.
	CouponsTable = &schema.Table{
		Name:       "coupons",
		Columns:    CouponsColumns,
		PrimaryKey: []*schema.Column{CouponsColumns[0]},
	}
	// CouponAllocatedsColumns holds the columns for the "coupon_allocateds" table.
	CouponAllocatedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "coupon_id", Type: field.TypeUUID},
		{Name: "denomination", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "used", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "used_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "used_by_order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 1699346837},
	}
	// CouponAllocatedsTable holds the schema information for the "coupon_allocateds" table.
	CouponAllocatedsTable = &schema.Table{
		Name:       "coupon_allocateds",
		Columns:    CouponAllocatedsColumns,
		PrimaryKey: []*schema.Column{CouponAllocatedsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "event_type", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "coupon_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "credits", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "credits_per_usd", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "max_consecutive", Type: field.TypeUint32, Nullable: true, Default: 1},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "inviter_layers", Type: field.TypeUint32, Nullable: true, Default: 1},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "event_ent_id",
				Unique:  true,
				Columns: []*schema.Column{EventsColumns[4]},
			},
		},
	}
	// InvitationCodesColumns holds the columns for the "invitation_codes" table.
	InvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "invitation_code", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// InvitationCodesTable holds the schema information for the "invitation_codes" table.
	InvitationCodesTable = &schema.Table{
		Name:       "invitation_codes",
		Columns:    InvitationCodesColumns,
		PrimaryKey: []*schema.Column{InvitationCodesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "invitationcode_ent_id",
				Unique:  true,
				Columns: []*schema.Column{InvitationCodesColumns[4]},
			},
		},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[5], PubsubMessagesColumns[6]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[5], PubsubMessagesColumns[7]},
			},
		},
	}
	// RegistrationsColumns holds the columns for the "registrations" table.
	RegistrationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "inviter_id", Type: field.TypeUUID, Nullable: true},
		{Name: "invitee_id", Type: field.TypeUUID, Nullable: true},
	}
	// RegistrationsTable holds the schema information for the "registrations" table.
	RegistrationsTable = &schema.Table{
		Name:       "registrations",
		Columns:    RegistrationsColumns,
		PrimaryKey: []*schema.Column{RegistrationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "registration_ent_id",
				Unique:  true,
				Columns: []*schema.Column{RegistrationsColumns[4]},
			},
		},
	}
	// ArchivementDetailsColumns holds the columns for the "archivement_details" table.
	ArchivementDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "direct_contributor_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "self_order", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "payment_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "units", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "units_v1", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "usd_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "commission", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ArchivementDetailsTable holds the schema information for the "archivement_details" table.
	ArchivementDetailsTable = &schema.Table{
		Name:       "archivement_details",
		Columns:    ArchivementDetailsColumns,
		PrimaryKey: []*schema.Column{ArchivementDetailsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "statement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ArchivementDetailsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArchivementGeneralsTable,
		CommissionsTable,
		CouponsTable,
		CouponAllocatedsTable,
		EventsTable,
		InvitationCodesTable,
		PubsubMessagesTable,
		RegistrationsTable,
		ArchivementDetailsTable,
	}
)

func init() {
	ArchivementGeneralsTable.Annotation = &entsql.Annotation{
		Table: "archivement_generals",
	}
	ArchivementDetailsTable.Annotation = &entsql.Annotation{
		Table: "archivement_details",
	}
}
