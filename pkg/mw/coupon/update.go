package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	"github.com/shopspring/decimal"
)

func (h *Handler) UpdateCoupon(ctx context.Context) (*npool.Coupon, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Coupon.
			Query().
			Where(
				entcoupon.ID(*h.ID),
				entcoupon.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		err1 := wlog.Errorf("endat less than startat")
		if h.StartAt != nil && h.EndAt != nil {
			if *h.EndAt <= *h.StartAt {
				return wlog.WrapError(err1)
			}
		}
		if h.StartAt != nil {
			if info.EndAt <= *h.StartAt {
				return wlog.WrapError(err1)
			}
		}
		if h.EndAt != nil {
			if *h.EndAt <= info.StartAt {
				return wlog.WrapError(err1)
			}
		}

		if info.CouponType == inspiretypes.CouponType_Discount.String() {
			if h.CashableProbability != nil && h.CashableProbability.Cmp(decimal.NewFromInt(0)) > 0 {
				return wlog.Errorf("discount can not set probability")
			}
		}

		if _, err := couponcrud.UpdateSet(
			info.Update(),
			&couponcrud.Req{
				Denomination:        h.Denomination,
				Circulation:         h.Circulation,
				StartAt:             h.StartAt,
				EndAt:               h.EndAt,
				DurationDays:        h.DurationDays,
				Message:             h.Message,
				Name:                h.Name,
				Random:              h.Random,
				Threshold:           h.Threshold,
				CouponScope:         h.CouponScope,
				CouponConstraint:    h.CouponConstraint,
				CashableProbability: h.CashableProbability,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetCoupon(ctx)
}
