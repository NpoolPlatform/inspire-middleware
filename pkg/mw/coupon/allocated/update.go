package allocated

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	allocatedmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/allocated"
	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
)

func UpdateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	_, err := allocatedmgrcli.UpdateAllocated(ctx, &allocatedmgrpb.AllocatedReq{
		ID:            in.ID,
		Used:          in.Used,
		UsedByOrderID: in.UsedByOrderID,
	})
	if err != nil {
		return nil, err
	}

	return GetCoupon(ctx, in.GetID())
}
