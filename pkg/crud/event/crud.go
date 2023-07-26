package event

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entevent "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/event"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	EventType      *basetypes.UsedFor
	CouponIDs      []uuid.UUID
	Credits        *decimal.Decimal
	CreditsPerUSD  *decimal.Decimal
	MaxConsecutive *uint32
	GoodID         *uuid.UUID
	InviterLayers  *uint32
	DeletedAt      *uint32
}

func CreateSet(c *ent.EventCreate, req *Req) *ent.EventCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	if len(req.CouponIDs) > 0 {
		c.SetCouponIds(req.CouponIDs)
	}
	if req.Credits != nil {
		c.SetCredits(*req.Credits)
	}
	if req.CreditsPerUSD != nil {
		c.SetCreditsPerUsd(*req.CreditsPerUSD)
	}
	if req.MaxConsecutive != nil {
		c.SetMaxConsecutive(*req.MaxConsecutive)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.InviterLayers != nil {
		c.SetInviterLayers(*req.InviterLayers)
	}
	return c
}

func UpdateSet(u *ent.EventUpdateOne, req *Req) *ent.EventUpdateOne {
	if len(req.CouponIDs) > 0 {
		u.SetCouponIds(req.CouponIDs)
	}
	if req.Credits != nil {
		u.SetCredits(*req.Credits)
	}
	if req.CreditsPerUSD != nil {
		u.SetCreditsPerUsd(*req.CreditsPerUSD)
	}
	if req.MaxConsecutive != nil {
		u.SetMaxConsecutive(*req.MaxConsecutive)
	}
	if req.InviterLayers != nil {
		u.SetInviterLayers(*req.InviterLayers)
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
	EventType *cruder.Cond
	GoodID    *cruder.Cond
}

func SetQueryConds(q *ent.EventQuery, conds *Conds) (*ent.EventQuery, error) {
	q.Where(entevent.DeletedAt(0))
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
			q.Where(entevent.ID(id))
		default:
			return nil, fmt.Errorf("invalid event field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.ID.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.ID.Op {
		case cruder.IN:
			q.Where(entevent.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid event field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entevent.AppID(id))
		default:
			return nil, fmt.Errorf("invalid event field")
		}
	}
	if conds.EventType != nil {
		_type, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid eventtype")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entevent.EventType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid event field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entevent.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid event field")
		}
	}
	return q, nil
}
