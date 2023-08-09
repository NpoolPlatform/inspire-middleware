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
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateCoupon(ctx context.Context) (*npool.Coupon, error) {
	id1 := uuid.New()
	if h.ID == nil {
		h.ID = &id1
	}

	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.CouponID == nil {
		return nil, fmt.Errorf("invalid couponid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}
	now := uint32(time.Now().Unix())

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		coup, err := tx.
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

		if coup.StartAt+coup.DurationDays*timedef.SecondsPerDay < now {
			return fmt.Errorf("coupon expired")
		}
		startAt := now
		if startAt < coup.StartAt {
			startAt = coup.StartAt
		}

		allocated := coup.Allocated
		switch coup.CouponType {
		case types.CouponType_FixAmount.String():
			allocated = allocated.Add(coup.Denomination)
		case types.CouponType_Discount.String():
			allocated = allocated.Add(decimal.NewFromInt(1))
		case types.CouponType_SpecialOffer.String():
			allocated = allocated.Add(coup.Denomination)
		default:
			return fmt.Errorf("invalid coupontype")
		}
		if allocated.Cmp(coup.Circulation) > 0 {
			return fmt.Errorf("insufficient circulation")
		}

		if _, err := allocatedcrud.CreateSet(
			tx.CouponAllocated.Create(),
			&allocatedcrud.Req{
				ID:           h.ID,
				AppID:        h.AppID,
				CouponID:     h.CouponID,
				UserID:       h.UserID,
				StartAt:      &startAt,
				Denomination: &coup.Denomination,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := couponcrud.UpdateSet(
			coup.Update(),
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
