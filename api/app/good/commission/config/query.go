package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAppGoodCommissionConfig(ctx context.Context, in *npool.GetAppGoodCommissionConfigRequest) (*npool.GetAppGoodCommissionConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodCommissionConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppGoodCommissionConfigs(ctx context.Context, in *npool.GetAppGoodCommissionConfigsRequest) (*npool.GetAppGoodCommissionConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodCommissionConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCommissionConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodCommissionConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodCommissionConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
