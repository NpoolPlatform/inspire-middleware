package fixamount

import (
	"context"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
)

func GetFixAmount(
	ctx context.Context,
	id string,
	handler func(context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error,
) {
	return nil, nil
}

func GetManyFixAmounts(
	ctx context.Context,
	ids []string,
	handler func(context.Context) ([]*allocatedmgrpb.Allocated, error),
) (
	[]*npool.Coupon,
	error,
) {
	return nil, nil
}

func GetFixAmounts(
	ctx context.Context,
	conds *allocatedmgrpb.Conds,
	offset, limit int32,
	handler func(context.Context) ([]*allocatedmgrpb.Allocated, error),
) (
	[]*npool.Coupon,
	uint32,
	error) {
	return nil, 0, nil
}

func GetFixAmountOnly(
	ctx context.Context,
	conds *allocatedmgrpb.Conds,
	handler func(context.Context) (*allocatedmgrpb.Allocated, error),
) (
	*npool.Coupon,
	error) {
	return nil, nil
}
