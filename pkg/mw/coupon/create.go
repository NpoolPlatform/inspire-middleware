package coupon

import (
	"context"
	"fmt"

	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateCoupon(ctx context.Context) (*npool.Coupon, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	if h.CouponType == nil {
		return nil, fmt.Errorf("invalid coupontype")
	}
	if h.Denomination == nil {
		return nil, fmt.Errorf("invalid denomination")
	}
	if h.Circulation == nil {
		return nil, fmt.Errorf("invalid circulation")
	}
	if h.Name == nil {
		return nil, fmt.Errorf("invalid name")
	}

	switch *h.CouponType {
	case types.CouponType_FixAmount:
		if h.Denomination.Cmp(*h.Circulation) > 0 {
			return nil, fmt.Errorf("denomination > circulation")
		}
	case types.CouponType_Discount:
		if h.Denomination.Cmp(decimal.NewFromInt(100)) > 0 { //nolint
			return nil, fmt.Errorf("100 discounat not allowed")
		}
	case types.CouponType_SpecialOffer:
		if h.UserID == nil {
			return nil, fmt.Errorf("userid is must")
		}
		if h.Denomination.Cmp(*h.Circulation) != 0 {
			return nil, fmt.Errorf("denomination != circulation")
		}
	}

	if h.CouponConstraint != nil {
		switch *h.CouponConstraint {
		case types.CouponConstraint_GoodThreshold:
			fallthrough //nolint
		case types.CouponConstraint_GoodOnly:
			if h.GoodID == nil {
				return nil, fmt.Errorf("goodid is must")
			}
			fallthrough
		case types.CouponConstraint_PaymentThreshold:
			if h.Threshold == nil {
				return nil, fmt.Errorf("threshold is must")
			}
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := couponcrud.CreateSet(
			cli.Coupon.Create(),
			&couponcrud.Req{
				ID:               h.ID,
				CouponType:       h.CouponType,
				AppID:            h.AppID,
				UserID:           h.UserID,
				GoodID:           h.GoodID,
				Denomination:     h.Denomination,
				Circulation:      h.Circulation,
				IssuedBy:         h.IssuedBy,
				StartAt:          h.StartAt,
				DurationDays:     h.DurationDays,
				Message:          h.Message,
				Name:             h.Name,
				CouponConstraint: h.CouponConstraint,
				Threshold:        h.Threshold,
				Allocated:        h.Allocated,
				Random:           h.Random,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoupon(ctx)
}
