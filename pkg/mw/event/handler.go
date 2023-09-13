package event

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	CouponIDs      []uuid.UUID
	Credits        *decimal.Decimal
	CreditsPerUSD  *decimal.Decimal
	MaxConsecutive *uint32
	InviterLayers  *uint32
	UserID         *uuid.UUID
	EventType      *basetypes.UsedFor
	GoodID         *uuid.UUID
	AppGoodID      *uuid.UUID
	Consecutive    *uint32
	Amount         *decimal.Decimal
	Conds          *eventcrud.Conds
	Offset         int32
	Limit          int32
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
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

func WithCouponIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_ids := []uuid.UUID{}
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			_ids = append(_ids, _id)
		}
		h.CouponIDs = _ids
		return nil
	}
}

func WithCredits(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid credits")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Credits = &_amount
		return nil
	}
}

func WithCreditsPerUSD(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid creditsperusd")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.CreditsPerUSD = &_amount
		return nil
	}
}

func WithMaxConsecutive(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxConsecutive = value
		return nil
	}
}

func WithInviterLayers(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.InviterLayers = value
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

func WithEventType(eventType *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventType == nil {
			if must {
				return fmt.Errorf("invalid eventtype")
			}
			return nil
		}

		switch *eventType {
		// Already implemented
		case basetypes.UsedFor_Signup:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliateSignup:
			fallthrough //nolint
		case basetypes.UsedFor_Purchase:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliatePurchase:
		// Not implemented
		case basetypes.UsedFor_Signin:
			fallthrough //nolint
		case basetypes.UsedFor_Update:
			fallthrough //nolint
		case basetypes.UsedFor_Contact:
			fallthrough //nolint
		case basetypes.UsedFor_SetWithdrawAddress:
			fallthrough //nolint
		case basetypes.UsedFor_Withdraw:
			fallthrough //nolint
		case basetypes.UsedFor_CreateInvitationCode:
			fallthrough //nolint
		case basetypes.UsedFor_SetCommission:
			fallthrough //nolint
		case basetypes.UsedFor_SetTransferTargetUser:
			fallthrough //nolint
		case basetypes.UsedFor_WithdrawalRequest:
			fallthrough //nolint
		case basetypes.UsedFor_WithdrawalCompleted:
			fallthrough //nolint
		case basetypes.UsedFor_DepositReceived:
			fallthrough //nolint
		case basetypes.UsedFor_KYCApproved:
			fallthrough //nolint
		case basetypes.UsedFor_KYCRejected:
			return fmt.Errorf("not implemented")
		default:
			return fmt.Errorf("invalid eventtype")
		}

		h.EventType = eventType
		return nil
	}
}

func WithGoodID(goodID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*goodID)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithAppGoodID(goodID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*goodID)
		if err != nil {
			return err
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithConsecutive(consecutive *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Consecutive = consecutive
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &eventcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.EventType != nil {
			h.Conds.EventType = &cruder.Cond{
				Op:  conds.GetEventType().GetOp(),
				Val: basetypes.UsedFor(conds.GetEventType().GetValue()),
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{
				Op:  conds.GetAppGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.IDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{
				Op:  conds.GetIDs().GetOp(),
				Val: ids,
			}
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
