//nolint:dupl
package config

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appconfig"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID               *uint32
	EntID            *uuid.UUID
	AppID            *uuid.UUID
	SettleMode       *types.SettleMode
	SettleAmountType *types.SettleAmountType
	SettleInterval   *types.SettleInterval
	CommissionType   *types.CommissionType
	SettleBenefit    *bool
	EndAt            *uint32
	StartAt          *uint32
	MaxLevelCount    *uint32
	DeletedAt        *uint32
}

func CreateSet(c *ent.AppConfigCreate, req *Req) *ent.AppConfigCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	c.SetEndAt(0)
	if req.CommissionType != nil {
		c.SetCommissionType(req.CommissionType.String())
	}
	if req.SettleAmountType != nil {
		c.SetSettleAmountType(req.SettleAmountType.String())
	}
	if req.SettleMode != nil {
		c.SetSettleMode(req.SettleMode.String())
	}
	if req.SettleInterval != nil {
		c.SetSettleInterval(req.SettleInterval.String())
	}
	if req.SettleBenefit != nil {
		c.SetSettleBenefit(*req.SettleBenefit)
	}
	if req.MaxLevelCount != nil {
		c.SetMaxLevelCount(*req.MaxLevelCount)
	}
	return c
}

func UpdateSet(u *ent.AppConfigUpdateOne, req *Req) *ent.AppConfigUpdateOne {
	if req.StartAt != nil {
		u = u.SetStartAt(*req.StartAt)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID            *cruder.Cond
	AppID            *cruder.Cond
	StartAt          *cruder.Cond
	EndAt            *cruder.Cond
	SettleMode       *cruder.Cond
	SettleAmountType *cruder.Cond
	SettleInterval   *cruder.Cond
	CommissionType   *cruder.Cond
	SettleBenefit    *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.AppConfigQuery, conds *Conds) (*ent.AppConfigQuery, error) { //nolint
	q.Where(entappconfig.DeletedAt(0))
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
			q.Where(entappconfig.EntID(id))
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
			q.Where(entappconfig.AppID(id))
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
			q.Where(entappconfig.EndAtLT(at))
		case cruder.LTE:
			q.Where(entappconfig.EndAtLTE(at))
		case cruder.GT:
			q.Where(entappconfig.EndAtGT(at))
		case cruder.GTE:
			q.Where(entappconfig.EndAtGTE(at))
		case cruder.EQ:
			q.Where(entappconfig.EndAt(at))
		case cruder.NEQ:
			q.Where(entappconfig.EndAtNEQ(at))
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
			q.Where(entappconfig.StartAtLT(at))
		case cruder.LTE:
			q.Where(entappconfig.StartAtLTE(at))
		case cruder.GT:
			q.Where(entappconfig.StartAtGT(at))
		case cruder.GTE:
			q.Where(entappconfig.StartAtGTE(at))
		case cruder.EQ:
			q.Where(entappconfig.StartAt(at))
		case cruder.NEQ:
			q.Where(entappconfig.StartAtNEQ(at))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.SettleMode != nil {
		settleType, ok := conds.SettleMode.Val.(types.SettleMode)
		if !ok {
			return nil, fmt.Errorf("invalid settlemode")
		}
		switch conds.SettleMode.Op {
		case cruder.EQ:
			q.Where(entappconfig.SettleMode(settleType.String()))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.SettleAmountType != nil {
		settleAmountType, ok := conds.SettleAmountType.Val.(types.SettleAmountType)
		if !ok {
			return nil, fmt.Errorf("invalid settleamounttype")
		}
		switch conds.SettleAmountType.Op {
		case cruder.EQ:
			q.Where(entappconfig.SettleAmountType(settleAmountType.String()))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.SettleInterval != nil {
		settleInterval, ok := conds.SettleInterval.Val.(types.SettleInterval)
		if !ok {
			return nil, fmt.Errorf("invalid settleinterval")
		}
		switch conds.SettleInterval.Op {
		case cruder.EQ:
			q.Where(entappconfig.SettleInterval(settleInterval.String()))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.CommissionType != nil {
		commissionType, ok := conds.CommissionType.Val.(types.CommissionType)
		if !ok {
			return nil, fmt.Errorf("invalid commissiontype")
		}
		switch conds.CommissionType.Op {
		case cruder.EQ:
			q.Where(entappconfig.CommissionType(commissionType.String()))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.SettleBenefit != nil {
		settleBenefit, ok := conds.SettleBenefit.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid settletype")
		}
		switch conds.SettleBenefit.Op {
		case cruder.EQ:
			q.Where(entappconfig.SettleBenefit(settleBenefit))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappconfig.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	return q, nil
}
