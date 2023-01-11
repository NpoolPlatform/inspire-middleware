package discount

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	entdiscount "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/coupondiscount"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateDiscount(
	ctx context.Context,
	id string,
	tx *ent.Tx,
	handler func(ctx context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error,
) {
	coup, err := tx.
		CouponDiscount.
		Query().
		Where(
			entdiscount.ID(uuid.MustParse(id)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	allocated := decimal.NewFromInt(int64(coup.Allocated + 1))
	if allocated.Cmp(coup.Circulation) > 0 {
		return nil, fmt.Errorf("overflow")
	}

	info, err := handler(ctx)
	if err != nil {
		return nil, err
	}

	coup, err = coup.Update().AddAllocated(1).Save(ctx)
	if err != nil {
		return nil, err
	}
	allocated = decimal.NewFromInt(int64(coup.Allocated))
	if allocated.Cmp(coup.Circulation) > 0 {
		return nil, fmt.Errorf("overflow")
	}

	startAt := info.CreatedAt
	if coup.StartAt > startAt {
		startAt = coup.StartAt
	}

	endAt := startAt + coup.DurationDays*timedef.SecondsPerDay

	return &npool.Coupon{
		ID:           info.ID,
		CouponType:   allocatedmgrpb.CouponType_Discount,
		AppID:        info.AppID,
		UserID:       info.UserID,
		Value:        coup.Discount.String(),
		Circulation:  coup.Circulation.String(),
		StartAt:      startAt,
		DurationDays: coup.DurationDays,
		EndAt:        endAt,
		CouponID:     id,
		CouponName:   coup.Name,
		Message:      coup.Message,
		Expired:      false,
		Valid:        uint32(time.Now().Unix()) >= startAt && uint32(time.Now().Unix()) <= endAt,
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
	}, nil
}
