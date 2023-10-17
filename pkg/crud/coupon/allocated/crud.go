package allocated

import (
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcouponallocated "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	CouponID      *uuid.UUID
	Used          *bool
	UsedByOrderID *uuid.UUID
	Denomination  *decimal.Decimal
	StartAt       *uint32
	DeletedAt     *uint32
}

func CreateSet(c *ent.CouponAllocatedCreate, req *Req) *ent.CouponAllocatedCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.Denomination != nil {
		c.SetDenomination(*req.Denomination)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	c.SetUsed(false)
	c.SetUsedAt(0)
	return c
}

func UpdateSet(u *ent.CouponAllocatedUpdateOne, req *Req) *ent.CouponAllocatedUpdateOne {
	if req.Used != nil && *req.Used && req.UsedByOrderID != nil {
		u.SetUsed(*req.Used)
		u.SetUsedAt(uint32(time.Now().Unix()))
		u.SetUsedByOrderID(*req.UsedByOrderID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	CouponType     *cruder.Cond
	CouponID       *cruder.Cond
	CouponIDs      *cruder.Cond
	Used           *cruder.Cond
	UsedByOrderID  *cruder.Cond
	UsedByOrderIDs *cruder.Cond
}

func SetQueryConds(q *ent.CouponAllocatedQuery, conds *Conds) (*ent.CouponAllocatedQuery, error) { //nolint
	q.Where(entcouponallocated.DeletedAt(0))
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
			q.Where(entcouponallocated.ID(id))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entcouponallocated.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid allocated ids")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.AppID(id))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.UserID(id))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.CouponID(id))
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
			q.Where(entcouponallocated.CouponIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid couponids field")
		}
	}
	if conds.Used != nil {
		used, ok := conds.Used.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid used")
		}
		switch conds.Used.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.Used(used))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	if conds.UsedByOrderID != nil {
		id, ok := conds.UsedByOrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid usedbyorderid")
		}
		switch conds.UsedByOrderID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.UsedByOrderID(id))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	if conds.UsedByOrderIDs != nil {
		ids, ok := conds.UsedByOrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid usedbyorderid")
		}
		switch conds.UsedByOrderIDs.Op {
		case cruder.IN:
			q.Where(entcouponallocated.UsedByOrderIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid allocated field")
		}
	}
	return q, nil
}
