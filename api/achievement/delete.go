package achievement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

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
	handler, err := achievement1.NewHandler(
		ctx,
		achievement1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteAchievement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAchievementResponse{
		Info: info,
	}, nil
}
