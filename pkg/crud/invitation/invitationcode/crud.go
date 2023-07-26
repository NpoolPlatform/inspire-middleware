package invitationcode

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	entinvitationcode "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/invitationcode"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	InvitationCode *string
	Disabled       *bool
	DeletedAt      *uint32
}

func CreateSet(c *ent.InvitationCodeCreate, req *Req) *ent.InvitationCodeCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.InvitationCode != nil {
		c.SetInvitationCode(*req.InvitationCode)
	}
	c.SetDisabled(false)
	return c
}

func UpdateSet(u *ent.InvitationCodeUpdateOne, req *Req) *ent.InvitationCodeUpdateOne {
	if req.Disabled != nil {
		u = u.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	InvitationCode *cruder.Cond
	Disabled       *cruder.Cond
	UserIDs        *cruder.Cond
}

func SetQueryConds(q *ent.InvitationCodeQuery, conds *Conds) (*ent.InvitationCodeQuery, error) {
	q.Where(entinvitationcode.DeletedAt(0))
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
			q.Where(entinvitationcode.ID(id))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entinvitationcode.AppID(id))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entinvitationcode.UserID(id))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	if conds.InvitationCode != nil {
		code, ok := conds.InvitationCode.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid invitationcode")
		}
		switch conds.InvitationCode.Op {
		case cruder.EQ:
			q.Where(entinvitationcode.InvitationCodeEQ(code))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entinvitationcode.DisabledEQ(disabled))
		case cruder.NEQ:
			q.Where(entinvitationcode.DisabledNEQ(disabled))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entinvitationcode.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid invitationcode field")
		}
	}
	return q, nil
}
