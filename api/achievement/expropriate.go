package achievement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExpropriateAchievement(ctx context.Context, in *npool.ExpropriateAchievementRequest) (*npool.ExpropriateAchievementResponse, error) {
	handler, err := achievement1.NewHandler(
		ctx,
		achievement1.WithOrderID(&in.OrderID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExpropriateAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.ExpropriateAchievementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.ExpropriateAchievement(ctx); err != nil {
		logger.Sugar().Errorw(
			"ExpropriateAchievement",
			"In", in,
			"Err", err,
		)
		return &npool.ExpropriateAchievementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExpropriateAchievementResponse{}, nil
}
