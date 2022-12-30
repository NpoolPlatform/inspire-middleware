package fixamount

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	entfixamount "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateFixAmount(
	ctx context.Context,
	id string,
	tx *ent.Tx,
	handler func(ctx context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error,
) {
	coup, err := tx.
		CouponFixAmount.
		Query().
		Where(
			entfixamount.ID(uuid.MustParse(id)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	allocated := coup.Denomination.Mul(decimal.NewFromInt(int64(coup.Allocated + 1)))
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
	allocated = coup.Denomination.Mul(decimal.NewFromInt(int64(coup.Allocated)))
	if allocated.Cmp(coup.Circulation) > 0 {
		return nil, fmt.Errorf("overflow")
	}

	now := uint32(time.Now().Unix())
	startAt := now

	if coup.StartAt > now {
		startAt = coup.StartAt
	}

	endAt := startAt + coup.DurationDays*timedef.SecondsPerDay

	return &npool.Coupon{
		ID:            info.ID,
		CouponTypeStr: allocatedmgrpb.CouponType_FixAmount.String(),
		CouponType:    allocatedmgrpb.CouponType_FixAmount,
		AppID:         info.AppID,
		UserID:        info.UserID,
		Value:         coup.Denomination.String(),
		Circulation:   coup.Circulation.String(),
		StartAt:       startAt,
		DurationDays:  coup.DurationDays,
		EndAt:         endAt,
		CouponID:      id,
		CouponName:    coup.Name,
		Message:       coup.Message,
		Expired:       false,
		Valid:         uint32(time.Now().Unix()) >= startAt && uint32(time.Now().Unix()) <= endAt,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.UpdatedAt,
	}, nil
}
