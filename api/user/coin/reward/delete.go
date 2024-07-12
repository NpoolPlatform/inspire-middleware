//nolint:dupl
package reward

import (
	"context"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/coin/reward"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
)

func (s *Server) DeleteUserCoinReward(ctx context.Context, in *npool.DeleteUserCoinRewardRequest) (*npool.DeleteUserCoinRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteUserCoinReward",
			"In", in,
		)
		return &npool.DeleteUserCoinRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithID(req.ID, false),
		reward1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteUserCoinReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteUserCoinReward",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCoinRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserCoinRewardResponse{}, nil
}
