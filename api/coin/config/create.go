package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
)

func (s *Server) CreateCoinConfig(ctx context.Context, in *npool.CreateCoinConfigRequest) (*npool.CreateCoinConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoinConfig",
			"In", in,
		)
		return &npool.CreateCoinConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithCoinTypeID(req.CoinTypeID, true),
		config1.WithMaxValue(req.MaxValue, true),
		config1.WithAllocated(req.Allocated, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateCoinConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateCoinConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCoinConfigResponse{}, nil
}
