package goodcoinachievement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	goodcoinachievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/good/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAchievement(ctx context.Context, in *npool.DeleteAchievementRequest) (*npool.DeleteAchievementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAchievement",
			"In", in,
		)
		return &npool.DeleteAchievementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := goodcoinachievement1.NewHandler(
		ctx,
		goodcoinachievement1.WithID(req.ID, false),
		goodcoinachievement1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.DeleteAchievement(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAchievementResponse{}, nil
}
