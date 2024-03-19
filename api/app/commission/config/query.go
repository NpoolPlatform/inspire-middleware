package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCommissionConfig(ctx context.Context, in *npool.GetAppCommissionConfigRequest) (*npool.GetAppCommissionConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCommissionConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCommissionConfigs(ctx context.Context, in *npool.GetAppCommissionConfigsRequest) (*npool.GetAppCommissionConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppCommissionConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCommissionConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppCommissionConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCommissionConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
