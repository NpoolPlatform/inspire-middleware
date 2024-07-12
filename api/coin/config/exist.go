package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
)

func (s *Server) ExistCoinConfigConds(ctx context.Context, in *npool.ExistCoinConfigCondsRequest) (*npool.ExistCoinConfigCondsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCoinConfigConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCoinConfigCondsResponse{
		Info: exist,
	}, nil
}
