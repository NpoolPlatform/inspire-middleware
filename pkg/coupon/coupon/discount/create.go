package discount

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	discountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/discount"
	discountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"
)

func CreateDiscount(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := discountmgrcli.CreateDiscount(ctx, &discountmgrpb.DiscountReq{
		ID:               in.ID,
		AppID:            in.AppID,
		Discount:         in.Value,
		Circulation:      in.Circulation,
		ReleasedByUserID: in.ReleasedByUserID,
		StartAt:          in.StartAt,
		DurationDays:     in.DurationDays,
		Message:          in.Message,
		Name:             in.Name,
	})
	if err != nil {
		return nil, err
	}
	return GetDiscount(ctx, info.ID)
}
