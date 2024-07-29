package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	user1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

func (s *Server) GetTaskUser(ctx context.Context, in *npool.GetTaskUserRequest) (*npool.GetTaskUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetTaskUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTaskUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTaskUsers(ctx context.Context, in *npool.GetTaskUsersRequest) (*npool.GetTaskUsersResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
		user1.WithOffset(in.GetOffset()),
		user1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTaskUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTaskUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
