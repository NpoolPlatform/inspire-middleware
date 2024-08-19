package reward

import (
	"context"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/coin/reward"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
)

func (s *Server) UpdateUserCoinReward(ctx context.Context, in *npool.UpdateUserCoinRewardRequest) (*npool.UpdateUserCoinRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateUserCoinReward",
			"In", in,
		)
		return &npool.UpdateUserCoinRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithID(req.ID, false),
		reward1.WithEntID(req.EntID, false),
		reward1.WithCoinRewards(req.CoinRewards, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateUserCoinReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateUserCoinRewardResponse{}, nil
}
