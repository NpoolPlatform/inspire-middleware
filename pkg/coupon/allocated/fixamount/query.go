package fixamount

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
)

func GetFixAmount(ctx context.Context, id string) (*npool.Coupon, error) {
	return nil, nil
}

func GetFixAmounts(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	return nil, 0, nil
}
