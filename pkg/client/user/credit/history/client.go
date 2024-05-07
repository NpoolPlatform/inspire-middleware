package history

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
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

func CreateUserCreditHistory(ctx context.Context, req *npool.UserCreditHistoryReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateUserCreditHistory(_ctx, &npool.CreateUserCreditHistoryRequest{
			Info: req,
		})
	})
	return err
}

func GetUserCreditHistory(ctx context.Context, id string) (*npool.UserCreditHistory, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCreditHistory(ctx, &npool.GetUserCreditHistoryRequest{
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
	return info.(*npool.UserCreditHistory), nil
}

func GetUserCreditHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.UserCreditHistory, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCreditHistories(ctx, &npool.GetUserCreditHistoriesRequest{
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
	return _infos.([]*npool.UserCreditHistory), total, nil
}

func ExistUserCreditHistoryConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistUserCreditHistoryConds(ctx, &npool.ExistUserCreditHistoryCondsRequest{
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

func GetUserCreditHistoryOnly(ctx context.Context, conds *npool.Conds) (*npool.UserCreditHistory, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCreditHistories(ctx, &npool.GetUserCreditHistoriesRequest{
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
	if len(infos.([]*npool.UserCreditHistory)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.UserCreditHistory)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.UserCreditHistory)[0], nil
}

func DeleteUserCreditHistory(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteUserCreditHistory(_ctx, &npool.DeleteUserCreditHistoryRequest{
			Info: &npool.UserCreditHistoryReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateUserCreditHistory(ctx context.Context, req *npool.UserCreditHistoryReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateUserCreditHistory(_ctx, &npool.UpdateUserCreditHistoryRequest{
			Info: req,
		})
	})
	return err
}
