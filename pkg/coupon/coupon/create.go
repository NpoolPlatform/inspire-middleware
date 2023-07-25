//nolint:dupl
package coupon

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	discount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/discount"
	fixamount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/fixamount"
	specialoffer "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon/specialoffer"
)

func CreateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	switch in.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
		return fixamount.CreateFixAmount(ctx, in)
	case allocatedmgrpb.CouponType_Discount:
		return discount.CreateDiscount(ctx, in)
	case allocatedmgrpb.CouponType_SpecialOffer:
		return specialoffer.CreateSpecialOffer(ctx, in)
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
