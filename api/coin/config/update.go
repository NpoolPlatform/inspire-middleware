package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
)

func (s *Server) UpdateCoinConfig(ctx context.Context, in *npool.UpdateCoinConfigRequest) (*npool.UpdateCoinConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCoinConfig",
			"In", in,
		)
		return &npool.UpdateCoinConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
		config1.WithMaxValue(req.MaxValue, false),
		config1.WithAllocated(req.Allocated, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateCoinConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCoinConfigResponse{}, nil
}
