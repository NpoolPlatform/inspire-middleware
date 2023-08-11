package registration

import (
	"context"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	InviterID *uuid.UUID
	InviteeID *uuid.UUID
	Conds     *registrationcrud.Conds
	Offset    int32
	Limit     int32
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

func WithInviterID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.InviterID = &_id
		return nil
	}
}

func WithInviteeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.InviteeID = &_id
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &registrationcrud.Conds{}
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
		if conds.InviterID != nil {
			id, err := uuid.Parse(conds.GetInviterID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.InviterID = &cruder.Cond{
				Op:  conds.GetInviterID().GetOp(),
				Val: id,
			}
		}
		if conds.InviteeID != nil {
			id, err := uuid.Parse(conds.GetInviteeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.InviteeID = &cruder.Cond{
				Op:  conds.GetInviteeID().GetOp(),
				Val: id,
			}
		}
		if conds.InviterIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetInviterIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.InviterIDs = &cruder.Cond{
				Op:  conds.GetInviterIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.InviteeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetInviteeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.InviteeIDs = &cruder.Cond{
				Op:  conds.GetInviteeIDs().GetOp(),
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
