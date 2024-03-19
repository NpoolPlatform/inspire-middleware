package config

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

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

func CreateCommissionConfig(ctx context.Context, in *npool.AppGoodCommissionConfigReq) (*npool.AppGoodCommissionConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAppGoodCommissionConfig(ctx, &npool.CreateAppGoodCommissionConfigRequest{
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
	return info.(*npool.AppGoodCommissionConfig), nil
}

func UpdateCommissionConfig(ctx context.Context, in *npool.AppGoodCommissionConfigReq) (*npool.AppGoodCommissionConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppGoodCommissionConfig(ctx, &npool.UpdateAppGoodCommissionConfigRequest{
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
	return info.(*npool.AppGoodCommissionConfig), nil
}

func GetCommissionConfig(ctx context.Context, id string) (*npool.AppGoodCommissionConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppGoodCommissionConfig(ctx, &npool.GetAppGoodCommissionConfigRequest{
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
	return info.(*npool.AppGoodCommissionConfig), nil
}

func GetCommissionConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.AppGoodCommissionConfig, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppGoodCommissionConfigs(ctx, &npool.GetAppGoodCommissionConfigsRequest{
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
	return infos.([]*npool.AppGoodCommissionConfig), total, nil
}

func GetCommissionConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.AppGoodCommissionConfig, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppGoodCommissionConfigs(ctx, &npool.GetAppGoodCommissionConfigsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.AppGoodCommissionConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.AppGoodCommissionConfig)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.AppGoodCommissionConfig)[0], nil
}

func CloneCommissionConfigs(ctx context.Context, req *npool.CloneAppGoodCommissionConfigsRequest) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.CloneAppGoodCommissionConfigs(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommissionConfig(ctx context.Context, id uint32) (*npool.AppGoodCommissionConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppGoodCommissionConfig(ctx, &npool.DeleteAppGoodCommissionConfigRequest{
			Info: &npool.AppGoodCommissionConfigReq{
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
	return info.(*npool.AppGoodCommissionConfig), nil
}

func ExistCommissionConfigConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAppGoodCommissionConfigConds(ctx, &npool.ExistAppGoodCommissionConfigCondsRequest{
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
