package commission

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uuid.UUID
	AppID           *uuid.UUID
	UserID          *uuid.UUID
	GoodID          *uuid.UUID
	AmountOrPercent *decimal.Decimal
	EndAt           *uint32
	StartAt         *uint32
	SettleType      *types.SettleType
	SettleMode      *types.SettleMode
	SettleInterval  *types.SettleInterval
	Threshold       *decimal.Decimal
	DeletedAt       *uint32
}

func CreateSet(c *ent.CommissionCreate, req *Req) *ent.CommissionCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
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
	if req.AmountOrPercent != nil {
		c.SetAmountOrPercent(*req.AmountOrPercent)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	c.SetEndAt(0)
	if req.SettleType != nil {
		c.SetSettleType(req.SettleType.String())
	}
	if req.SettleMode != nil {
		c.SetSettleMode(req.SettleMode.String())
	}
	if req.SettleInterval != nil {
		c.SetSettleInterval(req.SettleInterval.String())
	}
	if req.Threshold != nil {
		c.SetThreshold(*req.Threshold)
	}
	return c
}

func UpdateSet(u *ent.CommissionUpdateOne, req *Req) *ent.CommissionUpdateOne {
	if req.AmountOrPercent != nil {
		u = u.SetAmountOrPercent(*req.AmountOrPercent)
	}
	if req.StartAt != nil {
		u = u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u = u.SetEndAt(*req.EndAt)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	SettleType *cruder.Cond
	EndAt      *cruder.Cond
	UserIDs    *cruder.Cond
	GoodIDs    *cruder.Cond
	StartAt    *cruder.Cond
}

func SetQueryConds(q *ent.CommissionQuery, conds *Conds) (*ent.CommissionQuery, error) { //nolint
	q.Where(entcommission.DeletedAt(0))
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
			q.Where(entcommission.ID(id))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcommission.AppID(id))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcommission.UserID(id))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcommission.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.LT:
			q.Where(entcommission.EndAtLT(at))
		case cruder.GT:
			q.Where(entcommission.EndAtGT(at))
		case cruder.EQ:
			q.Where(entcommission.EndAt(at))
		case cruder.NEQ:
			q.Where(entcommission.EndAtNEQ(at))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LT:
			q.Where(entcommission.StartAtLT(at))
		case cruder.GT:
			q.Where(entcommission.StartAtGT(at))
		case cruder.EQ:
			q.Where(entcommission.StartAt(at))
		case cruder.NEQ:
			q.Where(entcommission.StartAtNEQ(at))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entcommission.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entcommission.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	return q, nil
}
