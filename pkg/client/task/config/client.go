package config

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
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

func CreateTaskConfig(ctx context.Context, req *npool.TaskConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTaskConfig(_ctx, &npool.CreateTaskConfigRequest{
			Info: req,
		})
	})
	return err
}

func GetTaskConfig(ctx context.Context, id string) (*npool.TaskConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskConfig(ctx, &npool.GetTaskConfigRequest{
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
	return info.(*npool.TaskConfig), nil
}

func GetTaskConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.TaskConfig, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskConfigs(ctx, &npool.GetTaskConfigsRequest{
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
	return _infos.([]*npool.TaskConfig), total, nil
}

func ExistTaskConfigConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistTaskConfigConds(ctx, &npool.ExistTaskConfigCondsRequest{
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

func GetTaskConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.TaskConfig, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskConfigs(ctx, &npool.GetTaskConfigsRequest{
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
	if len(infos.([]*npool.TaskConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TaskConfig)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.TaskConfig)[0], nil
}

func DeleteTaskConfig(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTaskConfig(_ctx, &npool.DeleteTaskConfigRequest{
			Info: &npool.TaskConfigReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateTaskConfig(ctx context.Context, req *npool.TaskConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTaskConfig(_ctx, &npool.UpdateTaskConfigRequest{
			Info: req,
		})
	})
	return err
}
