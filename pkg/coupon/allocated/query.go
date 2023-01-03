package allocated

import (
	"context"
	"fmt"

	allocatedmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/allocated"
	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	fixamount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated/fixamount"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func GetCoupon(ctx context.Context, id string) (*npool.Coupon, error) {
	info, err := allocatedmgrcli.GetAllocated(ctx, id)
	if err != nil {
		return nil, err
	}

	switch info.CouponType {
	case allocatedmgrpb.CouponType_FixAmount:
		return fixamount.GetFixAmount(
			ctx,
			info.CouponID,
			func(_ctx context.Context) (*allocatedmgrpb.Allocated, error) {
				return info, nil
			})
	case allocatedmgrpb.CouponType_Discount:
	case allocatedmgrpb.CouponType_SpecialOffer:
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
	case allocatedmgrpb.CouponType_ThresholdDiscount:
	case allocatedmgrpb.CouponType_GoodFixAmount:
	case allocatedmgrpb.CouponType_GoodDiscount:
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
	default:
		return nil, fmt.Errorf("unknown coupon type")
	}
	return nil, fmt.Errorf("not supported")
}

func getTypeCouponIDs(
	infos []*allocatedmgrpb.Allocated,
	couponType allocatedmgrpb.CouponType,
) []string {
	ids := []string{}
	for _, info := range infos {
		if info.CouponType == couponType {
			ids = append(ids, info.CouponID)
		}
	}
	return ids
}

func getTypeAllocateds(
	infos []*allocatedmgrpb.Allocated,
	couponType allocatedmgrpb.CouponType,
) []*allocatedmgrpb.Allocated {
	_infos := []*allocatedmgrpb.Allocated{}
	for _, info := range infos {
		if info.CouponType == couponType {
			_infos = append(_infos, info)
		}
	}
	return _infos
}

func expandManyCoupons(ctx context.Context, infos []*allocatedmgrpb.Allocated) ([]*npool.Coupon, error) {
	var coups []*npool.Coupon

	fmIDs := getTypeCouponIDs(infos, allocatedmgrpb.CouponType_FixAmount)
	if len(fmIDs) > 0 {
		_coups, err := fixamount.GetManyFixAmounts(
			ctx,
			fmIDs,
			func(_ctx context.Context) ([]*allocatedmgrpb.Allocated, error) {
				return getTypeAllocateds(infos, allocatedmgrpb.CouponType_FixAmount), nil
			})
		if err != nil {
			return nil, err
		}

		coups = append(coups, _coups...)
	}

	// allocatedmgrpb.CouponType_Discount:
	// allocatedmgrpb.CouponType_SpecialOffer:
	// allocatedmgrpb.CouponType_ThresholdFixAmount:
	// allocatedmgrpb.CouponType_ThresholdDiscount:
	// allocatedmgrpb.CouponType_GoodFixAmount:
	// allocatedmgrpb.CouponType_GoodDiscount:
	// allocatedmgrpb.CouponType_GoodThresholdFixAmount:
	// allocatedmgrpb.CouponType_GoodThresholdDiscount:

	return coups, nil
}

func GetManyCoupons(ctx context.Context, ids []string) ([]*npool.Coupon, error) {
	infos, _, err := allocatedmgrcli.GetAllocateds(ctx, &allocatedmgrpb.Conds{
		IDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	}, int32(0), int32(len(ids)))
	if err != nil {
		return nil, err
	}

	return expandManyCoupons(ctx, infos)
}

func GetCoupons(ctx context.Context, conds *allocatedmgrpb.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	infos, total, err := allocatedmgrcli.GetAllocateds(ctx, conds, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	coups, err := expandManyCoupons(ctx, infos)
	if err != nil {
		return nil, 0, err
	}

	return coups, total, nil
}

func GetCouponOnly(ctx context.Context, conds *allocatedmgrpb.Conds) (*npool.Coupon, error) {
	info, err := allocatedmgrcli.GetAllocatedOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	switch allocatedmgrpb.CouponType(conds.GetCouponType().GetValue()) {
	case allocatedmgrpb.CouponType_FixAmount:
		return fixamount.GetFixAmount(
			ctx,
			info.CouponID,
			func(_ctx context.Context) (*allocatedmgrpb.Allocated, error) {
				return info, nil
			})
	case allocatedmgrpb.CouponType_Discount:
	case allocatedmgrpb.CouponType_SpecialOffer:
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
	case allocatedmgrpb.CouponType_ThresholdDiscount:
	case allocatedmgrpb.CouponType_GoodFixAmount:
	case allocatedmgrpb.CouponType_GoodDiscount:
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
	default:
		return nil, fmt.Errorf("unknown coupon type")
	}
	return nil, fmt.Errorf("not supported")
}
