package specialoffer

import (
	"context"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	specialoffermgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/coupon/specialoffer"
	specialoffermgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"
)

func GetSpecialOffers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	infos, total, err := specialoffermgrcli.GetSpecialOffers(ctx, &specialoffermgrpb.Conds{
		ID:    conds.ID,
		AppID: conds.AppID,
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return specialOffers2Coupons(infos), total, nil
}

func GetSpecialOffer(ctx context.Context, id string) (*npool.Coupon, error) {
	info, err := specialoffermgrcli.GetSpecialOffer(ctx, id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return specialOffer2Coupon(info), nil
}

func specialOffer2Coupon(info *specialoffermgrpb.SpecialOffer) *npool.Coupon {
	return &npool.Coupon{
		ID:               info.ID,
		CouponType:       allocatedmgrpb.CouponType_SpecialOffer,
		AppID:            info.AppID,
		Value:            info.Amount,
		ReleasedByUserID: info.ReleasedByUserID,
		StartAt:          info.StartAt,
		DurationDays:     info.DurationDays,
		Message:          info.Message,
		UserID:           &info.UserID,
		CreatedAt:        info.CreatedAt,
		UpdatedAt:        info.UpdatedAt,
	}
}

func specialOffers2Coupons(infos []*specialoffermgrpb.SpecialOffer) []*npool.Coupon {
	var coupons []*npool.Coupon
	for _, info := range infos {
		coupons = append(coupons, specialOffer2Coupon(info))
	}
	return coupons
}
