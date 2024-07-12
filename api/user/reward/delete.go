package reward

import (
	"context"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/reward"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
)

func (s *Server) DeleteUserReward(ctx context.Context, in *npool.DeleteUserRewardRequest) (*npool.DeleteUserRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteUserReward",
			"In", in,
		)
		return &npool.DeleteUserRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithID(req.ID, false),
		reward1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteUserReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserRewardResponse{}, nil
}
