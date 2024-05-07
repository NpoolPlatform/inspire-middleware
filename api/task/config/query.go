package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

func (s *Server) GetTaskConfig(ctx context.Context, in *npool.GetTaskConfigRequest) (*npool.GetTaskConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetTaskConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTaskConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTaskConfigs(ctx context.Context, in *npool.GetTaskConfigsRequest) (*npool.GetTaskConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskConfigsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTaskConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTaskConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTaskConfigsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTaskConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
