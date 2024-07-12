package user

import (
	"context"

	user1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

func (s *Server) CreateTaskUser(ctx context.Context, in *npool.CreateTaskUserRequest) (*npool.CreateTaskUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTaskUser",
			"In", in,
		)
		return &npool.CreateTaskUserResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(req.EntID, false),
		user1.WithAppID(req.AppID, true),
		user1.WithUserID(req.UserID, true),
		user1.WithTaskID(req.TaskID, true),
		user1.WithEventID(req.EventID, true),
		user1.WithTaskState(req.TaskState, true),
		user1.WithRewardInfo(req.RewardInfo, true),
		user1.WithRewardState(req.RewardState, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateTaskUser(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTaskUserResponse{}, nil
}
