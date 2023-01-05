package coupon

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	discount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/discount"
	fixamount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/fixamount"
	specialoffer "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/specialoffer"
)

func GetCoupon(ctx context.Context, id string, couponType allocatedmgrpb.CouponType) (*npool.Coupon, error) {
	switch couponType {
	case allocatedmgrpb.CouponType_FixAmount:
		return fixamount.GetFixAmount(ctx, id)
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

func GetCoupons(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	switch allocatedmgrpb.CouponType(conds.GetCouponType().GetValue()) {
	case allocatedmgrpb.CouponType_FixAmount:
		return fixamount.GetFixAmounts(ctx, conds, offset, limit)
	case allocatedmgrpb.CouponType_Discount:
		return discount.GetDiscounts(ctx, conds, offset, limit)
	case allocatedmgrpb.CouponType_SpecialOffer:
		return specialoffer.GetSpecialOffers(ctx, conds, offset, limit)
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
	case allocatedmgrpb.CouponType_ThresholdDiscount:
	case allocatedmgrpb.CouponType_GoodFixAmount:
	case allocatedmgrpb.CouponType_GoodDiscount:
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
	default:
		return nil, 0, fmt.Errorf("unknown coupon type")
	}
	return nil, 0, fmt.Errorf("not supported")
}
