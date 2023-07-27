package registration

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entregistration "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/registration"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	InviterID *uuid.UUID
	InviteeID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.RegistrationCreate, req *Req) *ent.RegistrationCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.InviterID != nil {
		c.SetInviterID(*req.InviterID)
	}
	if req.InviteeID != nil {
		c.SetInviteeID(*req.InviteeID)
	}

	return c
}

func UpdateSet(u *ent.RegistrationUpdateOne, req *Req) *ent.RegistrationUpdateOne {
	if req.InviterID != nil {
		u = u.SetInviterID(*req.InviterID)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	InviterID  *cruder.Cond
	InviteeID  *cruder.Cond
	InviterIDs *cruder.Cond
	InviteeIDs *cruder.Cond
}

func SetQueryConds(q *ent.RegistrationQuery, conds *Conds) (*ent.RegistrationQuery, error) {
	q.Where(entregistration.DeletedAt(0))
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
			q.Where(entregistration.ID(id))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entregistration.AppID(id))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviterID != nil {
		id, ok := conds.InviterID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid inviterid")
		}
		switch conds.InviterID.Op {
		case cruder.EQ:
			q.Where(entregistration.InviterID(id))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviteeID != nil {
		id, ok := conds.InviteeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid inviteeid")
		}
		switch conds.InviteeID.Op {
		case cruder.EQ:
			q.Where(entregistration.InviteeID(id))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviterIDs != nil {
		ids, ok := conds.InviterIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid inviterids")
		}
		switch conds.InviterIDs.Op {
		case cruder.IN:
			q.Where(entregistration.InviterIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	if conds.InviteeIDs != nil {
		ids, ok := conds.InviteeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid inviteeids")
		}
		switch conds.InviteeIDs.Op {
		case cruder.IN:
			q.Where(entregistration.InviteeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid registration field")
		}
	}
	return q, nil
}
