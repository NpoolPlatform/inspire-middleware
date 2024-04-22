package config

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

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

func CreateCommissionConfig(ctx context.Context, in *npool.AppGoodCommissionConfigReq) (*npool.AppGoodCommissionConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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

	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppGoodCommissionConfigs(ctx, &npool.GetAppGoodCommissionConfigsRequest{
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
	if len(infos.([]*npool.AppGoodCommissionConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.AppGoodCommissionConfig)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.AppGoodCommissionConfig)[0], nil
}

func DeleteCommissionConfig(ctx context.Context, id *uint32, entID *string) (*npool.AppGoodCommissionConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteAppGoodCommissionConfig(ctx, &npool.DeleteAppGoodCommissionConfigRequest{
			Info: &npool.AppGoodCommissionConfigReq{
				ID:    id,
				EntID: entID,
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
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
