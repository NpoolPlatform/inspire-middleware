package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAppConfig(ctx context.Context, in *npool.GetAppConfigRequest) (*npool.GetAppConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppConfigs(ctx context.Context, in *npool.GetAppConfigsRequest) (*npool.GetAppConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAppConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
