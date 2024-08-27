package reward

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
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

func GetUserReward(ctx context.Context, id string) (*npool.UserReward, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserReward(ctx, &npool.GetUserRewardRequest{
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
	return info.(*npool.UserReward), nil
}

func GetUserRewards(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.UserReward, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserRewards(ctx, &npool.GetUserRewardsRequest{
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
	return _infos.([]*npool.UserReward), total, nil
}

func ExistUserRewardConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistUserRewardConds(ctx, &npool.ExistUserRewardCondsRequest{
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

func GetUserRewardOnly(ctx context.Context, conds *npool.Conds) (*npool.UserReward, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserRewards(ctx, &npool.GetUserRewardsRequest{
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
	if len(infos.([]*npool.UserReward)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.UserReward)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.UserReward)[0], nil
}

func DeleteUserReward(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteUserReward(_ctx, &npool.DeleteUserRewardRequest{
			Info: &npool.UserRewardReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func AddUserReward(ctx context.Context, req *npool.UserRewardReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.AddUserReward(_ctx, &npool.AddUserRewardRequest{
			Info: req,
		})
	})
	return err
}

func SubUserReward(ctx context.Context, req *npool.UserRewardReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.SubUserReward(_ctx, &npool.SubUserRewardRequest{
			Info: req,
		})
	})
	return err
}
