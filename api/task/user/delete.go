package user

import (
	"context"

	user1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

func (s *Server) DeleteTaskUser(ctx context.Context, in *npool.DeleteTaskUserRequest) (*npool.DeleteTaskUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTaskUser",
			"In", in,
		)
		return &npool.DeleteTaskUserResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID, false),
		user1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteTaskUser(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteTaskUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTaskUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTaskUserResponse{}, nil
}
