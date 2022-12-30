package coupon

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	fixamount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/fixamount"
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
