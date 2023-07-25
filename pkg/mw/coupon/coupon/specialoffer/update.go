package specialoffer

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	specialoffermgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/specialoffer"
	specialoffermgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"
)

func UpdateSpecialOffer(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := specialoffermgrcli.UpdateSpecialOffer(ctx, &specialoffermgrpb.SpecialOfferReq{
		ID:               in.ID,
		AppID:            in.AppID,
		UserID:           in.UserID,
		Amount:           in.Value,
		ReleasedByUserID: in.ReleasedByUserID,
		StartAt:          in.StartAt,
		DurationDays:     in.DurationDays,
		Message:          in.Message,
	})
	if err != nil {
		return nil, err
	}
	return GetSpecialOffer(ctx, info.ID)
}
