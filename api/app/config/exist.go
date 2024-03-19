package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAppConfigConds(ctx context.Context, in *npool.ExistAppConfigCondsRequest) (*npool.ExistAppConfigCondsResponse, error) {
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfigConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppConfigCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCommissionConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfigConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppConfigCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppConfigCondsResponse{
		Info: info,
	}, nil
}
