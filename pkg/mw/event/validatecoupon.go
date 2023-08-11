package event

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"

	"github.com/google/uuid"
)

func (h *Handler) validateCoupons(ctx context.Context) error {
	if len(h.CouponIDs) == 0 {
		return nil
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Coupon.
			Query().
			Where(
				entcoupon.IDIn(h.CouponIDs...),
				entcoupon.DeletedAt(0),
			).
			All(_ctx)
		if err != nil {
			return err
		}
		infoMap := map[uuid.UUID]*ent.Coupon{}
		now := uint32(time.Now().Unix())
		for _, info := range infos {
			if info.StartAt+info.DurationDays*timedef.SecondsPerDay <= now {
				return fmt.Errorf("coupon expired")
			}
			infoMap[info.ID] = info
		}
		for _, id := range h.CouponIDs {
			if _, ok := infoMap[id]; !ok {
				return fmt.Errorf("invalid couponid")
			}
		}
		return nil
	})
}
