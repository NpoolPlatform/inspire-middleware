package discount

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	discountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/discount"
	discountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"
)

func GetDiscounts(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	infos, total, err := discountmgrcli.GetDiscounts(ctx, &discountmgrpb.Conds{
		ID:    conds.ID,
		AppID: conds.AppID,
		IDs:   conds.IDs,
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return discounts2Coupons(infos), total, nil
}

func GetDiscount(ctx context.Context, id string) (*npool.Coupon, error) {
	info, err := discountmgrcli.GetDiscount(ctx, id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return discount2Coupon(info), nil
}

func discount2Coupon(info *discountmgrpb.Discount) *npool.Coupon {
	allocated := fmt.Sprintf("%v", info.Allocated)

	return &npool.Coupon{
		ID:               info.ID,
		CouponType:       allocatedmgrpb.CouponType_Discount,
		AppID:            info.AppID,
		Value:            info.Discount,
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

func discounts2Coupons(infos []*discountmgrpb.Discount) []*npool.Coupon {
	var coupons []*npool.Coupon
	for _, info := range infos {
		coupons = append(coupons, discount2Coupon(info))
	}
	return coupons
}
