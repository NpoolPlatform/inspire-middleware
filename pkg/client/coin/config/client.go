package config

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
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

func CreateCoinConfig(ctx context.Context, req *npool.CoinConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateCoinConfig(_ctx, &npool.CreateCoinConfigRequest{
			Info: req,
		})
	})
	return err
}

func GetCoinConfig(ctx context.Context, id string) (*npool.CoinConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinConfig(ctx, &npool.GetCoinConfigRequest{
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
	return info.(*npool.CoinConfig), nil
}

func GetCoinConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.CoinConfig, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinConfigs(ctx, &npool.GetCoinConfigsRequest{
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
	return _infos.([]*npool.CoinConfig), total, nil
}

func ExistCoinConfigConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistCoinConfigConds(ctx, &npool.ExistCoinConfigCondsRequest{
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

func GetCoinConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.CoinConfig, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinConfigs(ctx, &npool.GetCoinConfigsRequest{
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
	if len(infos.([]*npool.CoinConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.CoinConfig)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.CoinConfig)[0], nil
}

func DeleteCoinConfig(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteCoinConfig(_ctx, &npool.DeleteCoinConfigRequest{
			Info: &npool.CoinConfigReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateCoinConfig(ctx context.Context, req *npool.CoinConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateCoinConfig(_ctx, &npool.UpdateCoinConfigRequest{
			Info: req,
		})
	})
	return err
}
