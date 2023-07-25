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
	ID               *uuid.UUID
	CouponType       *types.CouponType
	AppID            *uuid.UUID
	UserID           *uuid.UUID
	GoodID           *uuid.UUID
	Denomination     *decimal.Decimal
	Circulation      *decimal.Decimal
	IssuedBy         *uuid.UUID
	StartAt          *uint32
	DurationDays     *uint32
	Message          *string
	Name             *string
	CouponConstraint *types.CouponConstraint
	Threshold        *decimal.Decimal
	Allocated        *decimal.Decimal
	DeletedAt        *uint32
}

func CreateSet(c *ent.CouponCreate, req *Req) *ent.CouponCreate {
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
	if req.IssuedBy != nil {
		c.SetIssuedBy(*req.IssuedBy)
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
	if req.Threshold != nil {
		c.SetThreshold(*req.Threshold)
	}
	if req.CouponConstraint != nil {
		c.SetCouponConstraint(req.CouponConstraint.String())
	}
	if req.Allocated != nil {
		c.SetAllocated(*req.Allocated)
	}
	return c
}

func UpdateSet(u *ent.CouponUpdateOne, req *Req) *ent.CouponUpdateOne {
	if req.Denomination != nil {
		u.SetDenomination(*req.Denomination)
	}
	if req.Circulation != nil {
		u.SetCirculation(*req.Circulation)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.DurationDays != nil {
		u.SetDurationDays(*req.DurationDays)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.Name != nil {
		u.SetName(*req.Message)
	}
	if req.Allocated != nil {
		u.SetAllocated(*req.Allocated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	CouponType *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CouponQuery, conds *Conds) (*ent.CouponQuery, error) {
	q.Where(entcoupon.DeletedAt(0))
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
			q.Where(entcoupon.ID(id))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entcoupon.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	if conds.CouponType != nil {
		_type, ok := conds.CouponType.Val.(types.CouponType)
		if !ok {
			return nil, fmt.Errorf("invalid coupontype")
		}
		switch conds.CouponType.Op {
		case cruder.EQ:
			q.Where(entcoupon.CouponType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoupon.AppID(id))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcoupon.UserID(id))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcoupon.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid coupon field")
		}
	}
	return q, nil
}
