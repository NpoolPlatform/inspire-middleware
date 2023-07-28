package general

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	entarchivementgeneral "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementgeneral"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uuid.UUID
	AppID           *uuid.UUID
	UserID          *uuid.UUID
	GoodID          *uuid.UUID
	CoinTypeID      *uuid.UUID
	TotalAmount     *decimal.Decimal
	SelfAmount      *decimal.Decimal
	TotalUnits      *decimal.Decimal
	SelfUnits       *decimal.Decimal
	TotalCommission *decimal.Decimal
	SelfCommission  *decimal.Decimal
}

func CreateSet(c *ent.ArchivementGeneralCreate, req *Req) *ent.ArchivementGeneralCreate {
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
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	c.SetTotalAmount(decimal.NewFromInt(0))
	c.SetSelfAmount(decimal.NewFromInt(0))
	c.SetTotalUnitsV1(decimal.NewFromInt(0))
	c.SetSelfUnitsV1(decimal.NewFromInt(0))
	c.SetTotalCommission(decimal.NewFromInt(0))
	c.SetSelfCommission(decimal.NewFromInt(0))

	return c
}

func UpdateSet(u *ent.ArchivementGeneralUpdateOne, req *Req) *ent.ArchivementGeneralUpdateOne {
	if req.TotalAmount != nil {
		u = u.SetTotalAmount(*req.TotalAmount)
	}
	if req.SelfAmount != nil {
		u = u.SetSelfAmount(*req.SelfAmount)
	}
	if req.TotalUnits != nil {
		u = u.SetTotalUnitsV1(*req.TotalUnits)
	}
	if req.SelfUnits != nil {
		u = u.SetSelfUnitsV1(*req.SelfUnits)
	}
	if req.TotalCommission != nil {
		u = u.SetTotalCommission(*req.TotalCommission)
	}
	if req.SelfCommission != nil {
		u = u.SetSelfCommission(*req.SelfCommission)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	CoinTypeID *cruder.Cond
	UserIDs    *cruder.Cond
}

func SetQueryConds(q *ent.ArchivementGeneralQuery, conds *Conds) (*ent.ArchivementGeneralQuery, error) { //nolint
	q.Where(entarchivementgeneral.DeletedAt(0))
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
			q.Where(entarchivementgeneral.ID(id))
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
			q.Where(entarchivementgeneral.AppID(id))
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
			q.Where(entarchivementgeneral.UserID(id))
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
			q.Where(entarchivementgeneral.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entarchivementgeneral.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entarchivementgeneral.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	return q, nil
}
