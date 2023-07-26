package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
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

func WithCouponIDs(ids []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		ids := []uuid.UUID{}
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.CouponIDs = ids
		return nil
	}
}

func WithCredits(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
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

func WithCreditsPerUSD(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
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

func WithMaxConsecutive(value *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxConsecutive = value
		return nil
	}
}

func WithInviterLayers(value *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.InviterLayers = value
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
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

//nolint
func WithEventType(eventType *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventType == nil {
			return nil
		}

		switch *eventType {
		/// Already implemented
		case basetypes.UsedFor_Signup:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliateSignup:
			fallthrough //nolint
		case basetypes.UsedFor_Purchase:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliatePurchase:
		/// Not implemented
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

func WithGoodID(goodID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
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

func WithConsecutive(consecutive uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Consecutive = consecutive
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
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
