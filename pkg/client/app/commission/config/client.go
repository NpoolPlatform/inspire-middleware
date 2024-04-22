package config

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
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

func CreateCommissionConfig(ctx context.Context, in *npool.AppCommissionConfigReq) (*npool.AppCommissionConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateAppCommissionConfig(ctx, &npool.CreateAppCommissionConfigRequest{
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
	return info.(*npool.AppCommissionConfig), nil
}

func UpdateCommissionConfig(ctx context.Context, in *npool.AppCommissionConfigReq) (*npool.AppCommissionConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.UpdateAppCommissionConfig(ctx, &npool.UpdateAppCommissionConfigRequest{
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
	return info.(*npool.AppCommissionConfig), nil
}

func GetCommissionConfig(ctx context.Context, id string) (*npool.AppCommissionConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppCommissionConfig(ctx, &npool.GetAppCommissionConfigRequest{
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
	return info.(*npool.AppCommissionConfig), nil
}

func GetCommissionConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.AppCommissionConfig, uint32, error) {
	var total uint32

	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppCommissionConfigs(ctx, &npool.GetAppCommissionConfigsRequest{
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
	return infos.([]*npool.AppCommissionConfig), total, nil
}

func GetCommissionConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.AppCommissionConfig, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppCommissionConfigs(ctx, &npool.GetAppCommissionConfigsRequest{
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
	if len(infos.([]*npool.AppCommissionConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.AppCommissionConfig)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.AppCommissionConfig)[0], nil
}

func ExistCommissionConfigConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistAppCommissionConfigConds(ctx, &npool.ExistAppCommissionConfigCondsRequest{
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

func DeleteCommissionConfig(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteAppCommissionConfig(_ctx, &npool.DeleteAppCommissionConfigRequest{
			Info: &npool.AppCommissionConfigReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
