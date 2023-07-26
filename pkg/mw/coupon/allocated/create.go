package allocated

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
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

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := allocatedcrud.CreateSet(
			cli.CouponAllocated.Create(),
			&allocatedcrud.Req{
				ID:       h.ID,
				AppID:    h.AppID,
				CouponID: h.CouponID,
				UserID:   h.UserID,
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
