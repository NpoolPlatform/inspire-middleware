package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
)

func (s *Server) DeleteCoinConfig(ctx context.Context, in *npool.DeleteCoinConfigRequest) (*npool.DeleteCoinConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCoinConfig",
			"In", in,
		)
		return &npool.DeleteCoinConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteCoinConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCoinConfigResponse{}, nil
}
