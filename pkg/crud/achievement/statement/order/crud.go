package orderstatement

import (
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                     *uint32
	EntID                  *uuid.UUID
	AppID                  *uuid.UUID
	UserID                 *uuid.UUID
	GoodID                 *uuid.UUID
	AppGoodID              *uuid.UUID
	OrderID                *uuid.UUID
	OrderUserID            *uuid.UUID
	PaymentID              *uuid.UUID
	CoinTypeID             *uuid.UUID
	PaymentCoinTypeID      *uuid.UUID
	PaymentCoinUSDCurrency *decimal.Decimal
	Units                  *decimal.Decimal
	Amount                 *decimal.Decimal
	USDAmount              *decimal.Decimal
	Commission             *decimal.Decimal
	AppConfigID            *uuid.UUID
	CommissionConfigID     *uuid.UUID
	CommissionConfigType   *types.CommissionConfigType
}

func CreateSet(c *ent.StatementCreate, req *Req) *ent.StatementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.DirectContributorID != nil {
		c.SetDirectContributorID(*req.DirectContributorID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.PaymentCoinTypeID != nil {
		c.SetPaymentCoinTypeID(*req.PaymentCoinTypeID)
	}
	if req.PaymentCoinUSDCurrency != nil {
		c.SetPaymentCoinUsdCurrency(*req.PaymentCoinUSDCurrency)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.USDAmount != nil {
		c.SetUsdAmount(*req.USDAmount)
	}
	if req.Commission != nil {
		c.SetCommission(*req.Commission)
	}
	if req.Units != nil {
		c.SetUnitsV1(*req.Units)
	}
	if req.AppConfigID != nil {
		c.SetAppConfigID(*req.AppConfigID)
	}
	if req.CommissionConfigID != nil {
		c.SetCommissionConfigID(*req.CommissionConfigID)
	}
	if req.CommissionConfigType != nil {
		c.SetCommissionConfigType(req.CommissionConfigType.String())
	}
	return c
}

type Conds struct {
	EntID                *cruder.Cond
	EntIDs               *cruder.Cond
	IDs                  *cruder.Cond
	AppID                *cruder.Cond
	UserID               *cruder.Cond
	DirectContributorID  *cruder.Cond
	GoodID               *cruder.Cond
	AppGoodID            *cruder.Cond
	OrderID              *cruder.Cond
	PaymentID            *cruder.Cond
	CoinTypeID           *cruder.Cond
	PaymentCoinTypeID    *cruder.Cond
	CreatedAt            *cruder.Cond
	UserIDs              *cruder.Cond
	AppConfigID          *cruder.Cond
	CommissionConfigID   *cruder.Cond
	CommissionConfigType *cruder.Cond
}

func SetQueryConds(q *ent.StatementQuery, conds *Conds) (*ent.StatementQuery, error) { //nolint
	q.Where(entstatement.DeletedAt(0))
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
			q.Where(entstatement.EntID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entstatement.AppID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entstatement.UserID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entstatement.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entstatement.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entstatement.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.DirectContributorID != nil {
		id, ok := conds.DirectContributorID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid directcontributorid")
		}
		switch conds.DirectContributorID.Op {
		case cruder.EQ:
			q.Where(entstatement.DirectContributorID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entstatement.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entstatement.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entstatement.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.PaymentID != nil {
		id, ok := conds.PaymentID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid paymentid")
		}
		switch conds.PaymentID.Op {
		case cruder.EQ:
			q.Where(entstatement.PaymentID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entstatement.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.PaymentCoinTypeID != nil {
		id, ok := conds.PaymentCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid paymentcointypeid")
		}
		switch conds.PaymentCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entstatement.PaymentCoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.AppConfigID != nil {
		id, ok := conds.AppConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appconfigid")
		}
		switch conds.AppConfigID.Op {
		case cruder.EQ:
			q.Where(entstatement.AppConfigID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.CommissionConfigID != nil {
		id, ok := conds.CommissionConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid commissionconfigid")
		}
		switch conds.CommissionConfigID.Op {
		case cruder.EQ:
			q.Where(entstatement.CommissionConfigID(id))
		default:
			return nil, fmt.Errorf("invalid statement field")
		}
	}
	if conds.CommissionConfigType != nil {
		commissionConfigType, ok := conds.CommissionConfigType.Val.(types.CommissionConfigType)
		if !ok {
			return nil, fmt.Errorf("invalid commissionconfigtype")
		}
		switch conds.CommissionConfigType.Op {
		case cruder.EQ:
			q.Where(entstatement.CommissionConfigType(commissionConfigType.String()))
		default:
			return nil, fmt.Errorf("invalid commission field")
		}
	}
	return q, nil
}
