package user

import (
	"context"

	user1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

func (s *Server) UpdateTaskUser(ctx context.Context, in *npool.UpdateTaskUserRequest) (*npool.UpdateTaskUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTaskUser",
			"In", in,
		)
		return &npool.UpdateTaskUserResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID, true),
		user1.WithTaskState(req.TaskState, false),
		user1.WithRewardInfo(req.RewardInfo, false),
		user1.WithRewardState(req.RewardState, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateTaskUser(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTaskUserResponse{}, nil
}
