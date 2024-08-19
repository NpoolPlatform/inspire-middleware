package config

import (
	"context"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

func (s *Server) UpdateTaskConfig(ctx context.Context, in *npool.UpdateTaskConfigRequest) (*npool.UpdateTaskConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTaskConfig",
			"In", in,
		)
		return &npool.UpdateTaskConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
		config1.WithEventID(req.EventID, false),
		config1.WithTaskType(req.TaskType, false),
		config1.WithName(req.Name, false),
		config1.WithTaskDesc(req.TaskDesc, false),
		config1.WithStepGuide(req.StepGuide, false),
		config1.WithRecommendMessage(req.RecommendMessage, false),
		config1.WithIndex(req.Index, false),
		config1.WithLastTaskID(req.LastTaskID, false),
		config1.WithMaxRewardCount(req.MaxRewardCount, false),
		config1.WithCooldownSecord(req.CooldownSecord, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateTaskConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateTaskConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTaskConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTaskConfigResponse{}, nil
}
