package role

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcouponscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	GoodID      *uuid.UUID
	CouponID    *uuid.UUID
	CouponScope *types.CouponScope
	DeletedAt   *uint32
}

func CreateSet(c *ent.CouponScopeCreate, req *Req) *ent.CouponScopeCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.CouponScope != nil {
		c.SetCouponScope(req.CouponScope.String())
	}
	return c
}

func UpdateSet(u *ent.CouponScopeUpdateOne, req *Req) *ent.CouponScopeUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	GoodID      *cruder.Cond
	CouponID    *cruder.Cond
	CouponIDs   *cruder.Cond
	CouponScope *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CouponScopeQuery, conds *Conds) (*ent.CouponScopeQuery, error) {
	q.Where(entcouponscope.DeletedAt(0))
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
			q.Where(entcouponscope.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcouponscope.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid goodid field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcouponscope.CouponID(id))
		default:
			return nil, fmt.Errorf("invalid couponid field")
		}
	}
	if conds.CouponScope != nil {
		scope, ok := conds.CouponScope.Val.(types.CouponScope)
		if !ok {
			return nil, fmt.Errorf("invalid couponscope")
		}
		switch conds.CouponScope.Op {
		case cruder.EQ:
			q.Where(entcouponscope.CouponScope(scope.String()))
		case cruder.NEQ:
			q.Where(entcouponscope.CouponScopeNEQ(scope.String()))
		default:
			return nil, fmt.Errorf("invalid couponscope field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid couponids")
		}
		switch conds.CouponIDs.Op {
		case cruder.IN:
			q.Where(entcouponscope.CouponIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid couponids field")
		}
	}
	return q, nil
}
