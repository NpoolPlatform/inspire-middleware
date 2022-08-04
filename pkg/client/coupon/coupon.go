package coupon

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/coupon"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetCoupon(ctx context.Context, id string, couponType npool.CouponType) (*npool.Coupon, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCoupon(ctx, &npool.GetCouponRequest{
			ID:   id,
			Type: couponType,
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Coupon), nil
}

func GetManyCoupons(ctx context.Context, ids []string, couponType npool.CouponType) ([]*npool.Coupon, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetManyCoupons(ctx, &npool.GetManyCouponsRequest{
			IDs:  ids,
			Type: couponType,
		})
		if err != nil {
			return nil, err
		}

		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Coupon), nil
}
