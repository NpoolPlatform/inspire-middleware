package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	common1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user/common"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAchievementUsers(ctx context.Context, in *npool.GetAchievementUsersRequest) (*npool.GetAchievementUsersResponse, error) {
	handler, err := achievement1.NewHandler(
		ctx,
		common1.WithConds(in.GetConds()),
		common1.WithOffset(in.GetOffset()),
		common1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchievementUsers",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchievementUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchievementUsers",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchievementUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAchievementUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
