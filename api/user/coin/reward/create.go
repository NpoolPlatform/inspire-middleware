package reward

import (
	"context"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/coin/reward"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
)

func (s *Server) CreateUserCoinReward(ctx context.Context, in *npool.CreateUserCoinRewardRequest) (*npool.CreateUserCoinRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateUserCoinReward",
			"In", in,
		)
		return &npool.CreateUserCoinRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithEntID(req.EntID, false),
		reward1.WithAppID(req.AppID, true),
		reward1.WithUserID(req.UserID, true),
		reward1.WithCoinTypeID(req.CoinTypeID, true),
		reward1.WithCoinRewards(req.CoinRewards, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateUserCoinReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserCoinRewardResponse{}, nil
}
