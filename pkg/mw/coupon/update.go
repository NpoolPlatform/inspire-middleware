package coupon

import (
	"context"

	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
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
			return err
		}

		if _, err := couponcrud.UpdateSet(
			info.Update(),
			&couponcrud.Req{
				Denomination:                  h.Denomination,
				Circulation:                   h.Circulation,
				StartAt:                       h.StartAt,
				EndAt:                         h.EndAt,
				DurationDays:                  h.DurationDays,
				Message:                       h.Message,
				Name:                          h.Name,
				Random:                        h.Random,
				Threshold:                     h.Threshold,
				CouponScope:                   h.CouponScope,
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
