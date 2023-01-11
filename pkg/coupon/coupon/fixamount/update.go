package fixamount

import (
	"context"
	"strconv"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	fixamountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/fixamount"
	fixamountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"
)

func UpdateFixAmount(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	var allocated uint32

	if in.Allocated != nil {
		_allocated, err := strconv.ParseUint(in.GetAllocated(), 10, 64) //nolint
		if err != nil {
			return nil, err
		}
		allocated = uint32(_allocated)
	}

	info, err := fixamountmgrcli.UpdateFixAmount(ctx, &fixamountmgrpb.FixAmountReq{
		ID:               in.ID,
		AppID:            in.AppID,
		Denomination:     in.Value,
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
	return GetFixAmount(ctx, info.ID)
}
