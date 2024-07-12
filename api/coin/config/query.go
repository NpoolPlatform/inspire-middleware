package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
)

func (s *Server) GetCoinConfig(ctx context.Context, in *npool.GetCoinConfigRequest) (*npool.GetCoinConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetCoinConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCoinConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinConfigs(ctx context.Context, in *npool.GetCoinConfigsRequest) (*npool.GetCoinConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinConfigsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetCoinConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinConfigsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCoinConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
