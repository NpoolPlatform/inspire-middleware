package coupon

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"
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

func CreateEventCoupon(ctx context.Context, req *npool.EventCouponReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateEventCoupon(_ctx, &npool.CreateEventCouponRequest{
			Info: req,
		})
	})
	return err
}

func GetEventCoupon(ctx context.Context, id string) (*npool.EventCoupon, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoupon(ctx, &npool.GetEventCouponRequest{
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
	return info.(*npool.EventCoupon), nil
}

func GetEventCoupons(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.EventCoupon, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoupons(ctx, &npool.GetEventCouponsRequest{
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
	return _infos.([]*npool.EventCoupon), total, nil
}

func ExistEventCouponConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistEventCouponConds(ctx, &npool.ExistEventCouponCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func GetEventCouponOnly(ctx context.Context, conds *npool.Conds) (*npool.EventCoupon, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoupons(ctx, &npool.GetEventCouponsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.EventCoupon)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.EventCoupon)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.EventCoupon)[0], nil
}

func DeleteEventCoupon(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteEventCoupon(_ctx, &npool.DeleteEventCouponRequest{
			Info: &npool.EventCouponReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateEventCoupon(ctx context.Context, req *npool.EventCouponReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateEventCoupon(_ctx, &npool.UpdateEventCouponRequest{
			Info: req,
		})
	})
	return err
}
