package specialoffer

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	entspecialoffer "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponspecialoffer"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
)

func CreateSpecialOffer(
	ctx context.Context,
	id string,
	tx *ent.Tx,
	handler func(ctx context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error,
) {
	coup, err := tx.
		CouponSpecialOffer.
		Query().
		Where(
			entspecialoffer.ID(uuid.MustParse(id)),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	info, err := handler(ctx)
	if err != nil {
		return nil, err
	}

	startAt := info.CreatedAt
	if coup.StartAt > startAt {
		startAt = coup.StartAt
	}

	endAt := startAt + coup.DurationDays*timedef.SecondsPerDay

	return &npool.Coupon{
		ID:           info.ID,
		CouponType:   allocatedmgrpb.CouponType_SpecialOffer,
		AppID:        info.AppID,
		UserID:       info.UserID,
		Value:        coup.Amount.String(),
		StartAt:      startAt,
		DurationDays: coup.DurationDays,
		EndAt:        endAt,
		CouponID:     id,
		Message:      coup.Message,
		Expired:      false,
		Valid:        uint32(time.Now().Unix()) >= startAt && uint32(time.Now().Unix()) <= endAt,
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
	}, nil
}
