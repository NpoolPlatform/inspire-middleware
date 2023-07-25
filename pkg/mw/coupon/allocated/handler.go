package allocated

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uuid.UUID
	CouponType    *types.CouponType
	AppID         *uuid.UUID
	CouponID      *uuid.UUID
	UserID        *uuid.UUID
	Used          *bool
	UsedByOrderID *uuid.UUID
	Conds         *allocatedcrud.Conds
	Offset        int32
	Limit         int32
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

func WithCouponType(couponType *types.CouponType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if couponType == nil {
			return nil
		}
		switch *couponType {
		case types.CouponType_FixAmount:
		case types.CouponType_Discount:
		case types.CouponType_SpecialOffer:
		default:
			return fmt.Errorf("invalid coupontype")
		}
		h.CouponType = couponType
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

func WithCouponID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		// TODO: check coupon exist
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CouponID = &_id
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

func WithUsed(value *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Used = value
		return nil
	}
}

func WithUsedByOrderID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UsedByOrderID = &_id
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &allocatedcrud.Conds{}
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
		if conds.CouponType != nil {
			h.Conds.CouponType = &cruder.Cond{
				Op:  conds.GetCouponType().GetOp(),
				Val: types.CouponType(conds.GetCouponType().GetValue()),
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
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.CouponID != nil {
			id, err := uuid.Parse(conds.GetCouponID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CouponID = &cruder.Cond{
				Op:  conds.GetCouponID().GetOp(),
				Val: id,
			}
		}
		if conds.Used != nil {
			h.Conds.Used = &cruder.Cond{
				Op:  conds.GetUsed().GetOp(),
				Val: conds.GetUsed().GetValue(),
			}
		}
		if conds.UsedByOrderID != nil {
			id, err := uuid.Parse(conds.GetUsedByOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UsedByOrderID = &cruder.Cond{
				Op:  conds.GetUsedByOrderID().GetOp(),
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
		if conds.UsedByOrderIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetUsedByOrderIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.UsedByOrderIDs = &cruder.Cond{
				Op:  conds.GetUsedByOrderIDs().GetOp(),
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
