package role

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uuid.UUID
	CouponType   *types.CouponType
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	GoodID       *string
	Denomination *decimal.Decimal
	Circulation  *decimal.Decimal
	ReleasedBy   *uuid.UUID
	StartAt      *uint32
	DurationDays *uint32
	Message      *string
	Name         *string
	Constraint   *types.CouponConstraint
	Threshold    *decimal.Decimal
	Allocated    *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppRoleCreate, req *Req) *ent.AppRoleCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CouponType != nil {
		c.SetCouponType(req.CouponType.String())
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Denomination != nil {
		c.SetDenomination(*req.Denomination)
	}
	if req.Circulation != nil {
		c.SetCirculation(*req.Circulation)
	}
	if req.ReleasedBy != nil {
		c.SetReleasedBy(*req.ReleasedBy)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.DurationDays != nil {
		c.SetDurationDays(*req.DurationDays)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.Name != nil {
		c.SetName(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.AppRoleUpdateOne, req *Req) *ent.AppRoleUpdateOne {
	if req.Role != nil {
		u.SetRole(*req.Role)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.Default != nil {
		u.SetDefault(*req.Default)
	}
	if req.Genesis != nil {
		u.SetGenesis(*req.Genesis)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	AppID     *cruder.Cond
	AppIDs    *cruder.Cond
	CreatedBy *cruder.Cond
	Role      *cruder.Cond
	Default   *cruder.Cond
	Genesis   *cruder.Cond
	Roles     *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppRoleQuery, conds *Conds) (*ent.AppRoleQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entapprole.ID(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapprole.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entapprole.AppID(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entapprole.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.CreatedBy != nil {
		id, ok := conds.CreatedBy.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.CreatedBy.Op {
		case cruder.EQ:
			q.Where(entapprole.CreatedBy(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Role != nil {
		role, ok := conds.Role.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid role")
		}
		switch conds.Role.Op {
		case cruder.EQ:
			q.Where(entapprole.Role(role))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Default != nil {
		defautl, ok := conds.Default.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid default")
		}
		switch conds.Default.Op {
		case cruder.EQ:
			q.Where(entapprole.Default(defautl))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Genesis != nil {
		genesis, ok := conds.Genesis.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid default")
		}
		switch conds.Genesis.Op {
		case cruder.EQ:
			q.Where(entapprole.Genesis(genesis))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Roles != nil {
		roles, ok := conds.Roles.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid roles")
		}
		switch conds.Roles.Op {
		case cruder.IN:
			q.Where(entapprole.RoleIn(roles...))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	return q, nil
}
