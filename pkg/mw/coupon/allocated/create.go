package allocated

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	inspiretypes "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateCoupon(ctx context.Context) (*npool.Coupon, error) {
	id1 := uuid.New()
	if h.ID == nil {
		h.ID = &id1
	}

	now := uint32(time.Now().Unix())
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		coupon, err := tx.
			Coupon.
			Query().
			Where(
				entcoupon.ID(*h.CouponID),
				entcoupon.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if coupon.StartAt+coupon.DurationDays*timedef.SecondsPerDay < now {
			return fmt.Errorf("coupon expired")
		}
		startAt := now
		if startAt < coupon.StartAt {
			startAt = coupon.StartAt
		}

		allocated := coupon.Allocated
		switch coupon.CouponType {
		case inspiretypes.CouponType_FixAmount.String():
			allocated = allocated.Add(coupon.Denomination)
		case inspiretypes.CouponType_Discount.String():
			allocated = allocated.Add(decimal.NewFromInt(1))
		case inspiretypes.CouponType_SpecialOffer.String():
			allocated = allocated.Add(coupon.Denomination)
		default:
			return fmt.Errorf("invalid coupontype")
		}
		if allocated.Cmp(coupon.Circulation) > 0 {
			return fmt.Errorf("insufficient circulation")
		}

		couponScope := inspiretypes.CouponScope(inspiretypes.CouponScope_value[coupon.CouponScope])
		if _, err := allocatedcrud.CreateSet(
			tx.CouponAllocated.Create(),
			&allocatedcrud.Req{
				ID:           h.ID,
				AppID:        h.AppID,
				CouponID:     h.CouponID,
				UserID:       h.UserID,
				StartAt:      &startAt,
				Denomination: &coupon.Denomination,
				CouponScope:  &couponScope,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := couponcrud.UpdateSet(
			coupon.Update(),
			&couponcrud.Req{
				Allocated: &allocated,
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
