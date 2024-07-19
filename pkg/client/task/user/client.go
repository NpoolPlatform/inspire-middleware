package user

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
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

func CreateTaskUser(ctx context.Context, req *npool.TaskUserReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTaskUser(_ctx, &npool.CreateTaskUserRequest{
			Info: req,
		})
	})
	return err
}

func GetTaskUser(ctx context.Context, id string) (*npool.TaskUser, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskUser(ctx, &npool.GetTaskUserRequest{
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
	return info.(*npool.TaskUser), nil
}

func GetTaskUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.TaskUser, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskUsers(ctx, &npool.GetTaskUsersRequest{
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
	return _infos.([]*npool.TaskUser), total, nil
}

func ExistTaskUserConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistTaskUserConds(ctx, &npool.ExistTaskUserCondsRequest{
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

func GetTaskUserOnly(ctx context.Context, conds *npool.Conds) (*npool.TaskUser, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTaskUsers(ctx, &npool.GetTaskUsersRequest{
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
	if len(infos.([]*npool.TaskUser)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TaskUser)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.TaskUser)[0], nil
}

func DeleteTaskUser(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTaskUser(_ctx, &npool.DeleteTaskUserRequest{
			Info: &npool.TaskUserReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateTaskUser(ctx context.Context, req *npool.TaskUserReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTaskUser(_ctx, &npool.UpdateTaskUserRequest{
			Info: req,
		})
	})
	return err
}
