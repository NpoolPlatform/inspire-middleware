package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

func (s *Server) CreateTaskConfig(ctx context.Context, in *npool.CreateTaskConfigRequest) (*npool.CreateTaskConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTaskConfig",
			"In", in,
		)
		return &npool.CreateTaskConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithEventID(req.EventID, true),
		config1.WithTaskType(req.TaskType, true),
		config1.WithName(req.Name, true),
		config1.WithTaskDesc(req.TaskDesc, true),
		config1.WithStepGuide(req.StepGuide, true),
		config1.WithRecommendMessage(req.RecommendMessage, true),
		config1.WithIndex(req.Index, true),
		config1.WithLastTaskID(req.LastTaskID, false),
		config1.WithMaxRewardCount(req.MaxRewardCount, true),
		config1.WithCooldownSecord(req.CooldownSecord, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateTaskConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTaskConfigResponse{}, nil
}
