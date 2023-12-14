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
	if h.EntID == nil {
		h.EntID = &id
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
	}

	if h.CouponConstraint != nil {
		switch *h.CouponConstraint {
		case types.CouponConstraint_Normal:
			threshold := decimal.RequireFromString("0")
			h.Threshold = &threshold
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
				EntID:                         h.EntID,
				CouponType:                    h.CouponType,
				AppID:                         h.AppID,
				Denomination:                  h.Denomination,
				Circulation:                   h.Circulation,
				IssuedBy:                      h.IssuedBy,
				StartAt:                       h.StartAt,
				EndAt:                         h.EndAt,
				DurationDays:                  h.DurationDays,
				Message:                       h.Message,
				Name:                          h.Name,
				CouponConstraint:              h.CouponConstraint,
				CouponScope:                   h.CouponScope,
				Threshold:                     h.Threshold,
				Allocated:                     h.Allocated,
				Random:                        h.Random,
				CashableProbabilityPerMillion: h.CashableProbabilityPerMillion,
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
