package statement

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	statementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	statementcrud.Req
	Reqs   []*statementcrud.Req
	Conds  *statementcrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}
func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithDirectContributorID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid directcontributorid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.DirectContributorID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderID = &_id
		return nil
	}
}

func WithSelfOrder(selfOrder *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SelfOrder = selfOrder
		return nil
	}
}

func WithPaymentID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithPaymentCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid paymentcointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PaymentCoinTypeID = &_id
		return nil
	}
}

func WithPaymentCoinUSDCurrency(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid paymentcoinusdcurrency")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.PaymentCoinUSDCurrency = &_amount
		return nil
	}
}

func WithUnits(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid units")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Units = &_amount
		return nil
	}
}

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Amount = &_amount
		return nil
	}
}

func WithUSDAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid usdamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.USDAmount = &_amount
		return nil
	}
}

func WithCommission(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid commission")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Commission = &_amount
		return nil
	}
}

func WithAppConfigID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appconfigid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppConfigID = &_id
		return nil
	}
}

func WithCommissionConfigID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid commissionconfigid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CommissionConfigID = &_id
		return nil
	}
}

func WithCommissionConfigType(value *types.CommissionConfigType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid commissionconfigtype")
			}
			return nil
		}
		switch *value {
		case types.CommissionConfigType_AppCommissionConfig:
		case types.CommissionConfigType_AppGoodCommissionConfig:
		case types.CommissionConfigType_LegacyCommissionConfig:
		default:
			return fmt.Errorf("invalid commissionconfigtype")
		}
		h.CommissionConfigType = value
		return nil
	}
}

func WithReqs(reqs []*npool.StatementReq, must bool) func(context.Context, *Handler) error { //nolint
	return func(ctx context.Context, h *Handler) error {
		appMap := map[string]struct{}{}
		orderMap := map[string]struct{}{}
		_reqs := []*statementcrud.Req{}

		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.UserID == nil {
					return fmt.Errorf("invalid userid")
				}
				if req.GoodID == nil {
					return fmt.Errorf("invalid goodid")
				}
				if req.AppGoodID == nil {
					return fmt.Errorf("invalid appgoodid")
				}
				if req.OrderID == nil {
					return fmt.Errorf("invalid orderid")
				}
				if req.PaymentID == nil {
					return fmt.Errorf("invalid paymentid")
				}
				if req.CoinTypeID == nil {
					return fmt.Errorf("invalid paymentid")
				}
				if req.PaymentCoinTypeID == nil {
					return fmt.Errorf("invalid paymentcointypeid")
				}
				if req.PaymentCoinUSDCurrency == nil {
					return fmt.Errorf("invalid paymentcoinusdcurrency")
				}
				if req.Units == nil {
					return fmt.Errorf("invalid units")
				}
				if req.Amount == nil {
					return fmt.Errorf("invalid amount")
				}
				if req.USDAmount == nil {
					return fmt.Errorf("invalid usdamount")
				}
				if req.Commission == nil {
					return fmt.Errorf("invalid commission")
				}
				if req.AppConfigID == nil {
					return fmt.Errorf("invalid appconfigid")
				}
				if req.CommissionConfigID == nil {
					return fmt.Errorf("invalid commissionconfigid")
				}
				if req.CommissionConfigType == nil {
					return fmt.Errorf("invalid commissionconfigtype")
				}
			}
			if !must {
				if req.ID == nil {
					return fmt.Errorf("invalid id")
				}
			}
			_req := &statementcrud.Req{
				SelfOrder: req.SelfOrder,
			}

			if req.ID != nil {
				_req.ID = req.ID
			}

			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}

			if req.AppID != nil {
				id1, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &id1
			}

			if req.UserID != nil {
				id2, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &id2
			}

			if req.DirectContributorID != nil {
				id3, err := uuid.Parse(*req.DirectContributorID)
				if err != nil {
					return err
				}
				_req.DirectContributorID = &id3
			}

			if req.GoodID != nil {
				id4, err := uuid.Parse(*req.GoodID)
				if err != nil {
					return err
				}
				_req.GoodID = &id4
			}

			if req.OrderID != nil {
				id5, err := uuid.Parse(*req.OrderID)
				if err != nil {
					return err
				}
				_req.OrderID = &id5
			}

			if req.PaymentID != nil {
				id6, err := uuid.Parse(*req.PaymentID)
				if err != nil {
					return err
				}
				_req.PaymentID = &id6
			}

			if req.CoinTypeID != nil {
				id7, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &id7
			}

			if req.PaymentCoinTypeID != nil {
				id8, err := uuid.Parse(*req.PaymentCoinTypeID)
				if err != nil {
					return err
				}
				_req.PaymentCoinTypeID = &id8
			}

			if req.AppGoodID != nil {
				id9, err := uuid.Parse(*req.AppGoodID)
				if err != nil {
					return err
				}
				_req.AppGoodID = &id9
			}

			if req.PaymentCoinUSDCurrency != nil {
				amount1, err := decimal.NewFromString(*req.PaymentCoinUSDCurrency)
				if err != nil {
					return err
				}
				_req.PaymentCoinUSDCurrency = &amount1
			}

			if req.Units != nil {
				amount2, err := decimal.NewFromString(*req.Units)
				if err != nil {
					return err
				}
				_req.Units = &amount2
			}

			if req.Amount != nil {
				amount3, err := decimal.NewFromString(*req.Amount)
				if err != nil {
					return err
				}
				_req.Amount = &amount3
			}

			if req.USDAmount != nil {
				amount4, err := decimal.NewFromString(*req.USDAmount)
				if err != nil {
					return err
				}
				_req.USDAmount = &amount4
			}

			if req.Commission != nil {
				amount5, err := decimal.NewFromString(*req.Commission)
				if err != nil {
					return err
				}
				_req.Commission = &amount5
			}

			if req.AppConfigID != nil {
				id10, err := uuid.Parse(*req.AppConfigID)
				if err != nil {
					return err
				}
				_req.AppConfigID = &id10
			}

			if req.CommissionConfigID != nil {
				id11, err := uuid.Parse(*req.CommissionConfigID)
				if err != nil {
					return err
				}
				_req.CommissionConfigID = &id11
			}

			if req.CommissionConfigType != nil {
				switch *req.CommissionConfigType {
				case types.CommissionConfigType_AppCommissionConfig:
				case types.CommissionConfigType_AppGoodCommissionConfig:
				case types.CommissionConfigType_LegacyCommissionConfig:
				default:
					return fmt.Errorf("invalid commissionconfigtype")
				}
				_req.CommissionConfigType = req.CommissionConfigType
			}

			if req.AppID != nil {
				appMap[*req.AppID] = struct{}{}
			}
			if req.OrderID != nil {
				orderMap[*req.OrderID] = struct{}{}
			}
			_reqs = append(_reqs, _req)
		}

		if len(appMap) > 1 {
			return fmt.Errorf("too many apps")
		}
		if len(orderMap) > 1 {
			return fmt.Errorf("too many orders")
		}

		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &statementcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.UserIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.UserIDs.GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.UserIDs = &cruder.Cond{Op: conds.GetUserIDs().GetOp(), Val: ids}
		}
		if conds.DirectContributorID != nil {
			id, err := uuid.Parse(conds.GetDirectContributorID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.DirectContributorID = &cruder.Cond{Op: conds.GetDirectContributorID().GetOp(), Val: id}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: id}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{Op: conds.GetAppGoodID().GetOp(), Val: id}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
		}
		if conds.SelfOrder != nil {
			h.Conds.SelfOrder = &cruder.Cond{Op: conds.GetSelfOrder().GetOp(), Val: conds.GetSelfOrder().GetValue()}
		}
		if conds.PaymentID != nil {
			id, err := uuid.Parse(conds.GetPaymentID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PaymentID = &cruder.Cond{Op: conds.GetPaymentID().GetOp(), Val: id}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.PaymentCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPaymentCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PaymentCoinTypeID = &cruder.Cond{Op: conds.GetPaymentCoinTypeID().GetOp(), Val: id}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
