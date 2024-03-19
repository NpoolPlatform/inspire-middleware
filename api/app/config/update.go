package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAppConfig(ctx context.Context, in *npool.UpdateAppConfigRequest) (*npool.UpdateAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, true),
		config1.WithStartAt(req.StartAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppConfigResponse{
		Info: info,
	}, nil
}
