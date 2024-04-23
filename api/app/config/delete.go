package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAppConfig(ctx context.Context, in *npool.DeleteAppConfigRequest) (*npool.DeleteAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.DeleteAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppConfigResponse{}, nil
}
