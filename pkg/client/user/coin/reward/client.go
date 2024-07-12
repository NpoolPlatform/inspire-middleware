package reward

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
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

func CreateUserCoinReward(ctx context.Context, req *npool.UserCoinRewardReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateUserCoinReward(_ctx, &npool.CreateUserCoinRewardRequest{
			Info: req,
		})
	})
	return err
}

func GetUserCoinReward(ctx context.Context, id string) (*npool.UserCoinReward, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCoinReward(ctx, &npool.GetUserCoinRewardRequest{
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
	return info.(*npool.UserCoinReward), nil
}

func GetUserCoinRewards(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.UserCoinReward, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCoinRewards(ctx, &npool.GetUserCoinRewardsRequest{
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
	return _infos.([]*npool.UserCoinReward), total, nil
}

func ExistUserCoinRewardConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistUserCoinRewardConds(ctx, &npool.ExistUserCoinRewardCondsRequest{
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

func GetUserCoinRewardOnly(ctx context.Context, conds *npool.Conds) (*npool.UserCoinReward, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetUserCoinRewards(ctx, &npool.GetUserCoinRewardsRequest{
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
	if len(infos.([]*npool.UserCoinReward)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.UserCoinReward)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.UserCoinReward)[0], nil
}

func DeleteUserCoinReward(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteUserCoinReward(_ctx, &npool.DeleteUserCoinRewardRequest{
			Info: &npool.UserCoinRewardReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateUserCoinReward(ctx context.Context, req *npool.UserCoinRewardReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateUserCoinReward(_ctx, &npool.UpdateUserCoinRewardRequest{
			Info: req,
		})
	})
	return err
}
