package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAchievementUser(ctx context.Context, in *npool.DeleteAchievementUserRequest) (*npool.DeleteAchievementUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAchievementUser",
			"In", in,
		)
		return &npool.DeleteAchievementUserResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := achievement1.NewHandler(
		ctx,
		achievement1.WithID(req.ID, false),
		achievement1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievementUser",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.DeleteAchievementUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchievementUser",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchievementUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAchievementUserResponse{}, nil
}
