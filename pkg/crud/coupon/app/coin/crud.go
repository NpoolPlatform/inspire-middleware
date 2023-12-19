package coin

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcouponcoin "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	CouponID   *uuid.UUID
	CoinTypeID *uuid.UUID
	DeletedAt  *uint32
}

func CreateSet(c *ent.CouponCoinCreate, req *Req) *ent.CouponCoinCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	return c
}

func UpdateSet(u *ent.CouponCoinUpdateOne, req *Req) *ent.CouponCoinUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
	CouponID   *cruder.Cond
	CouponIDs  *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CouponCoinQuery, conds *Conds) (*ent.CouponCoinQuery, error) {
	q.Where(entcouponcoin.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcouponcoin.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcouponcoin.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcouponcoin.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodid field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcouponcoin.CouponID(id))
		default:
			return nil, fmt.Errorf("invalid couponid field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid couponids")
		}
		switch conds.CouponIDs.Op {
		case cruder.IN:
			q.Where(entcouponcoin.CouponIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid couponids field")
		}
	}
	return q, nil
}
