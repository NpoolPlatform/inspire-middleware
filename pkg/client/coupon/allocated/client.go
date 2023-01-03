package allocated

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

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

func CreateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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

func UpdateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCoupon(ctx, &npool.GetCouponRequest{
			ID: id,
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

func GetManyCoupons(ctx context.Context, ids []string) ([]*npool.Coupon, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetManyCoupons(ctx, &npool.GetManyCouponsRequest{
			IDs: ids,
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

func GetCoupons(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Coupon, uint32, error) {
	var total uint32

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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

func GetCouponOnly(ctx context.Context, conds *mgrpb.Conds) (*npool.Coupon, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponOnly(ctx, &npool.GetCouponOnlyRequest{
			Conds: conds,
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
