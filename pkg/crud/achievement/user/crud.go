package user

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievementuser "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievementuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                *uuid.UUID
	AppID                *uuid.UUID
	UserID               *uuid.UUID
	TotalCommission      *decimal.Decimal
	SelfCommission       *decimal.Decimal
	DirectInvites        *uint32
	IndirectInvites      *uint32
	DirectConsumeAmount  *decimal.Decimal
	InviteeConsumeAmount *decimal.Decimal
}

func CreateSet(c *ent.AchievementUserCreate, req *Req) *ent.AchievementUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.TotalCommission != nil {
		c.SetTotalCommission(*req.TotalCommission)
	}
	if req.SelfCommission != nil {
		c.SetSelfCommission(*req.SelfCommission)
	}
	if req.DirectInvites != nil {
		c.SetDirectInvites(*req.DirectInvites)
	}
	if req.IndirectInvites != nil {
		c.SetIndirectInvites(*req.IndirectInvites)
	}
	if req.DirectConsumeAmount != nil {
		c.SetDirectConsumeAmount(*req.DirectConsumeAmount)
	}
	if req.InviteeConsumeAmount != nil {
		c.SetInviteeConsumeAmount(*req.InviteeConsumeAmount)
	}

	return c
}

func UpdateSet(u *ent.AchievementUserUpdateOne, req *Req) *ent.AchievementUserUpdateOne {
	if req.TotalCommission != nil {
		u = u.SetTotalCommission(*req.TotalCommission)
	}
	if req.SelfCommission != nil {
		u = u.SetSelfCommission(*req.SelfCommission)
	}
	if req.DirectInvites != nil {
		u = u.SetDirectInvites(*req.DirectInvites)
	}
	if req.DirectInvites != nil {
		u = u.SetDirectInvites(*req.DirectInvites)
	}
	if req.DirectConsumeAmount != nil {
		u = u.SetDirectConsumeAmount(*req.DirectConsumeAmount)
	}
	if req.InviteeConsumeAmount != nil {
		u = u.SetInviteeConsumeAmount(*req.InviteeConsumeAmount)
	}
	return u
}

type Conds struct {
	EntID   *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	UserIDs *cruder.Cond
}

func SetQueryConds(q *ent.AchievementUserQuery, conds *Conds) (*ent.AchievementUserQuery, error) {
	q.Where(entachievementuser.DeletedAt(0))
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
			q.Where(entachievementuser.EntID(id))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entachievementuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entachievementuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entachievementuser.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	return q, nil
}
