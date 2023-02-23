package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/event"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	coupon "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	allocated "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated"
	allocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/shopspring/decimal"
)

func RewardEvent(
	ctx context.Context,
	appID, userID string,
	eventType basetypes.UsedFor,
	goodID *string,
	consecutive uint32,
	amount decimal.Decimal,
) (decimal.Decimal, error) {
	conds := &mgrpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(eventType)},
	}
	if goodID != nil && eventType == basetypes.UsedFor_Purchase {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *goodID}
	}

	info, err := mgrcli.GetEventOnly(ctx, conds)
	if err != nil {
		return decimal.Decimal{}, err
	}
	if info == nil {
		return decimal.Decimal{}, fmt.Errorf("event is invalid")
	}

	if consecutive > info.MaxConsecutive {
		return decimal.Decimal{}, nil
	}

	credits, err := decimal.NewFromString(info.Credits)
	if err != nil {
		return decimal.Decimal{}, err
	}

	_credits, err := decimal.NewFromString(info.CreditsPerUSD)
	if err != nil {
		return decimal.Decimal{}, err
	}

	credits = credits.Add(_credits.Mul(amount))

	coups := []*couponmwpb.Coupon{}

	for _, coup := range info.Coupons {
		_coupon, err := coupon.GetCoupon(ctx, coup.ID, coup.CouponType)
		if err != nil {
			return decimal.Decimal{}, err
		}
		if _coupon == nil {
			return decimal.Decimal{}, fmt.Errorf("coupon is invalid")
		}

		coups = append(coups, _coupon)
	}

	for _, coup := range coups {
		_, err := allocated.CreateCoupon(ctx, &allocatedmwpb.CouponReq{
			AppID:      &appID,
			UserID:     &userID,
			CouponID:   &coup.ID,
			CouponType: &coup.CouponType,
		})
		if err != nil {
			return decimal.Decimal{}, err
		}
	}

	return credits, nil
}
