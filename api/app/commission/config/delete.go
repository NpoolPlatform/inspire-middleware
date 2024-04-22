package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAppCommissionConfig(ctx context.Context, in *npool.DeleteAppCommissionConfigRequest) (*npool.DeleteAppCommissionConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
		)
		return &npool.DeleteAppCommissionConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.DeleteCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppCommissionConfigResponse{
		Info: nil,
	}, nil
}
