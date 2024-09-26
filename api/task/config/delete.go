package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

func (s *Server) DeleteTaskConfig(ctx context.Context, in *npool.DeleteTaskConfigRequest) (*npool.DeleteTaskConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTaskConfig",
			"In", in,
		)
		return &npool.DeleteTaskConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteTaskConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTaskConfigResponse{}, nil
}
