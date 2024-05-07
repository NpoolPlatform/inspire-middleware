package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
)

func (s *Server) ExistUserCoinRewardConds(ctx context.Context, in *npool.ExistUserCoinRewardCondsRequest) (*npool.ExistUserCoinRewardCondsResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserCoinRewardConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCoinRewardCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistUserCoinRewardConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserCoinRewardConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCoinRewardCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistUserCoinRewardCondsResponse{
		Info: exist,
	}, nil
}
