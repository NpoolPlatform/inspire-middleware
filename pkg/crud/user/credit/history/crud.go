package history

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entusercredithistory "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/usercredithistory"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	TaskID    *uuid.UUID
	EventID   *uuid.UUID
	Credits   *decimal.Decimal
	DeletedAt *uint32
}

func CreateSet(c *ent.UserCreditHistoryCreate, req *Req) *ent.UserCreditHistoryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.TaskID != nil {
		c.SetTaskID(*req.TaskID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.Credits != nil {
		c.SetCredits(*req.Credits)
	}
	return c
}

func UpdateSet(u *ent.UserCreditHistoryUpdateOne, req *Req) *ent.UserCreditHistoryUpdateOne {
	if req.Credits != nil {
		u.SetCredits(*req.Credits)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	TaskID  *cruder.Cond
	EventID *cruder.Cond
	ID      *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.UserCreditHistoryQuery, conds *Conds) (*ent.UserCreditHistoryQuery, error) {
	q.Where(entusercredithistory.DeletedAt(0))
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
			q.Where(entusercredithistory.EntID(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entusercredithistory.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entusercredithistory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entusercredithistory.UserID(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.TaskID != nil {
		id, ok := conds.TaskID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid taskid")
		}
		switch conds.TaskID.Op {
		case cruder.EQ:
			q.Where(entusercredithistory.TaskID(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.EventID != nil {
		id, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(entusercredithistory.EventID(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entusercredithistory.ID(id))
		case cruder.NEQ:
			q.Where(entusercredithistory.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid usercredithistory field")
		}
	}
	return q, nil
}
