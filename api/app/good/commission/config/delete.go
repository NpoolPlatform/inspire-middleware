package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAppGoodCommissionConfig(ctx context.Context, in *npool.DeleteAppGoodCommissionConfigRequest) (*npool.DeleteAppGoodCommissionConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
		)
		return &npool.DeleteAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppGoodCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppGoodCommissionConfigResponse{
		Info: info,
	}, nil
}
