package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/reward"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
)

func (s *Server) GetUserReward(ctx context.Context, in *npool.GetUserRewardRequest) (*npool.GetUserRewardResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetUserReward(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserRewardResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserRewards(ctx context.Context, in *npool.GetUserRewardsRequest) (*npool.GetUserRewardsResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithConds(in.GetConds()),
		reward1.WithOffset(in.GetOffset()),
		reward1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserRewardsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetUserRewards(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserRewardsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserRewardsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
