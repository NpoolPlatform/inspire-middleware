package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
)

func (s *Server) GetUserCoinReward(ctx context.Context, in *npool.GetUserCoinRewardRequest) (*npool.GetUserCoinRewardResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetUserCoinReward(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCoinRewardResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserCoinRewards(ctx context.Context, in *npool.GetUserCoinRewardsRequest) (*npool.GetUserCoinRewardsResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithConds(in.GetConds()),
		reward1.WithOffset(in.GetOffset()),
		reward1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCoinRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCoinRewardsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetUserCoinRewards(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCoinRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCoinRewardsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCoinRewardsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
