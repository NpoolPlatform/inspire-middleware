package fixamount

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	fixamountmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/fixamount"
	fixamountmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"
)

func CreateFixAmount(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := fixamountmgrcli.CreateFixAmount(ctx, &fixamountmgrpb.FixAmountReq{
		ID:               in.ID,
		AppID:            in.AppID,
		Denomination:     in.Value,
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
	return GetFixAmount(ctx, info.ID)
}
