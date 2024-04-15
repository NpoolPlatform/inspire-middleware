package coupon

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateCoupon(ctx, &npool.CreateCouponRequest{
			Info: in,
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

func GetCoupons(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	var total uint32

	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoupons(ctx, &npool.GetCouponsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Coupon), total, nil
}

func UpdateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.UpdateCoupon(ctx, &npool.UpdateCouponRequest{
			Info: in,
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

func GetCoupon(ctx context.Context, id string) (*npool.Coupon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoupon(ctx, &npool.GetCouponRequest{
			EntID: id,
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

func ExistCoupon(ctx context.Context, id string) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoupon(ctx, &npool.GetCouponRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	if info == nil {
		return false, nil
	}
	return true, nil
}

func DeleteCoupon(ctx context.Context, id uint32) (*npool.Coupon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteCoupon(ctx, &npool.DeleteCouponRequest{
			Info: &npool.CouponReq{
				ID: &id,
			},
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
