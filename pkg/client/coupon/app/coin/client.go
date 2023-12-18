package coin

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateCouponCoin(ctx context.Context, in *npool.CouponCoinReq) (*npool.CouponCoin, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCouponCoin(ctx, &npool.CreateCouponCoinRequest{
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
	return info.(*npool.CouponCoin), nil
}

func GetCouponCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.CouponCoin, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponCoins(ctx, &npool.GetCouponCoinsRequest{
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
	return infos.([]*npool.CouponCoin), total, nil
}

func GetCouponCoin(ctx context.Context, id string) (*npool.CouponCoin, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponCoin(ctx, &npool.GetCouponCoinRequest{
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
	return info.(*npool.CouponCoin), nil
}

func DeleteCouponCoin(ctx context.Context, id uint32) (*npool.CouponCoin, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteCouponCoin(ctx, &npool.DeleteCouponCoinRequest{
			Info: &npool.CouponCoinReq{
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
	return info.(*npool.CouponCoin), nil
}

func ExistCouponCoinConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistCouponCoinConds(ctx, &npool.ExistCouponCoinCondsRequest{
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
