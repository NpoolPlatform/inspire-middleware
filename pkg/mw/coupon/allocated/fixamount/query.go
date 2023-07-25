package fixamount

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	fixamountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/fixamount"
	fixamountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func expand(info *allocatedmgrpb.Allocated, coup *fixamountmgrpb.FixAmount) *npool.Coupon {
	startAt := coup.StartAt
	if startAt < info.CreatedAt {
		startAt = info.CreatedAt
	}
	endAt := startAt + coup.DurationDays*timedef.SecondsPerDay

	return &npool.Coupon{
		ID:            info.ID,
		CouponType:    allocatedmgrpb.CouponType_FixAmount,
		AppID:         info.AppID,
		UserID:        info.UserID,
		Value:         coup.Denomination,
		Circulation:   coup.Circulation,
		StartAt:       startAt,
		DurationDays:  coup.DurationDays,
		EndAt:         endAt,
		CouponID:      coup.ID,
		CouponName:    coup.Name,
		Message:       coup.Message,
		Expired:       uint32(time.Now().Unix()) > endAt,
		Valid:         uint32(time.Now().Unix()) >= startAt && uint32(time.Now().Unix()) <= endAt,
		Used:          info.Used,
		UsedAt:        info.UsedAt,
		UsedByOrderID: info.UsedByOrderID,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.UpdatedAt,
	}
}

func expandMany(infos []*allocatedmgrpb.Allocated, coups []*fixamountmgrpb.FixAmount) []*npool.Coupon {
	coupMap := map[string]*fixamountmgrpb.FixAmount{}
	for _, coup := range coups {
		coupMap[coup.ID] = coup
	}

	_coups := []*npool.Coupon{}

	for _, info := range infos {
		coup, ok := coupMap[info.CouponID]
		if !ok {
			continue
		}

		_coups = append(_coups, expand(info, coup))
	}

	return _coups
}

func GetFixAmount(
	ctx context.Context,
	id string,
	handler func(context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error,
) {
	coup, err := fixamountmgrcli.GetFixAmount(ctx, id)
	if err != nil {
		return nil, err
	}

	info, err := handler(ctx)
	if err != nil {
		return nil, err
	}

	return expand(info, coup), nil
}

func GetManyFixAmounts(
	ctx context.Context,
	ids []string,
	handler func(context.Context) ([]*allocatedmgrpb.Allocated, error),
) (
	[]*npool.Coupon,
	error,
) {
	coups, _, err := fixamountmgrcli.GetFixAmounts(ctx, &fixamountmgrpb.Conds{
		IDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	}, int32(0), int32(len(ids)))
	if err != nil {
		return nil, err
	}

	infos, err := handler(ctx)
	if err != nil {
		return nil, err
	}

	return expandMany(infos, coups), nil
}
