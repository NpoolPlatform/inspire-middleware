package achievement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAchievements(ctx context.Context, in *npool.GetAchievementsRequest) (*npool.GetAchievementsResponse, error) {
	handler, err := achievement1.NewHandler(
		ctx,
		achievement1.WithConds(in.GetConds()),
		achievement1.WithOffset(in.GetOffset()),
		achievement1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchievements",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchievementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAchievements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchievements",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchievementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAchievementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
