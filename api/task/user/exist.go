package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	user1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
)

func (s *Server) ExistTaskUserConds(ctx context.Context, in *npool.ExistTaskUserCondsRequest) (*npool.ExistTaskUserCondsResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTaskUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTaskUserCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTaskUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTaskUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTaskUserCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTaskUserCondsResponse{
		Info: exist,
	}, nil
}
