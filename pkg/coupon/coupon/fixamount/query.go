package fixamount

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	fixamountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/fixamount"
	fixamountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"
)

func GetFixAmounts(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	infos, total, err := fixamountmgrcli.GetFixAmounts(ctx, &fixamountmgrpb.Conds{
		ID:    conds.ID,
		AppID: conds.AppID,
		IDs:   conds.IDs,
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return fixAmounts2Coupons(infos), total, nil
}

func GetFixAmount(ctx context.Context, id string) (*npool.Coupon, error) {
	info, err := fixamountmgrcli.GetFixAmount(ctx, id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return fixAmount2Coupon(info), nil
}

func fixAmount2Coupon(info *fixamountmgrpb.FixAmount) *npool.Coupon {
	allocated := fmt.Sprintf("%v", info.Allocated)

	return &npool.Coupon{
		ID:               info.ID,
		CouponType:       allocatedmgrpb.CouponType_FixAmount,
		AppID:            info.AppID,
		Value:            info.Denomination,
		Circulation:      info.Circulation,
		ReleasedByUserID: info.ReleasedByUserID,
		StartAt:          info.StartAt,
		DurationDays:     info.DurationDays,
		Message:          info.Message,
		Name:             info.Name,
		Allocated:        allocated,
		CreatedAt:        info.CreatedAt,
		UpdatedAt:        info.UpdatedAt,
	}
}

func fixAmounts2Coupons(infos []*fixamountmgrpb.FixAmount) []*npool.Coupon {
	var coupons []*npool.Coupon
	for _, info := range infos {
		coupons = append(coupons, fixAmount2Coupon(info))
	}
	return coupons
}
