package discount

import (
	"context"
	"strconv"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	discountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/discount"
	discountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/discount"
)

func UpdateDiscount(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	var allocated uint32

	if in.Allocated != nil {
		_allocated, err := strconv.ParseUint(in.GetAllocated(), 10, 64) //nolint
		if err != nil {
			return nil, err
		}
		allocated = uint32(_allocated)
	}

	info, err := discountmgrcli.UpdateDiscount(ctx, &discountmgrpb.DiscountReq{
		ID:               in.ID,
		AppID:            in.AppID,
		Discount:         in.Value,
		Circulation:      in.Circulation,
		ReleasedByUserID: in.ReleasedByUserID,
		StartAt:          in.StartAt,
		DurationDays:     in.DurationDays,
		Message:          in.Message,
		Name:             in.Name,
		Allocated:        &allocated,
	})
	if err != nil {
		return nil, err
	}
	return GetDiscount(ctx, info.ID)
}
