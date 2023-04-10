package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	AppID       string
	UserID      string
	EventType   basetypes.UsedFor
	GoodID      *string
	Consecutive uint32
	Amount      string
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

func WithAppID(appID string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(appID); err != nil {
			return err
		}
		h.AppID = appID
		return nil
	}
}

func WithUserID(userID string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(userID); err != nil {
			return err
		}
		h.UserID = userID
		return nil
	}
}

//nolint
func WithEventType(eventType basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch eventType {
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
		if _, err := uuid.Parse(*goodID); err != nil {
			return err
		}
		h.GoodID = goodID
		return nil
	}
}

func WithConsecutive(consecutive uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Consecutive = consecutive
		return nil
	}
}

func WithAmount(amount string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amount, err := decimal.NewFromString(amount)
		if err != nil {
			_amount = decimal.NewFromInt(0)
		}
		h.Amount = _amount.String()
		return nil
	}
}
