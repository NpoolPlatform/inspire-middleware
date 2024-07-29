package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/reward"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
)

func (s *Server) ExistUserRewardConds(ctx context.Context, in *npool.ExistUserRewardCondsRequest) (*npool.ExistUserRewardCondsResponse, error) {
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserRewardConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserRewardCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistUserRewardConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserRewardConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserRewardCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistUserRewardCondsResponse{
		Info: exist,
	}, nil
}
