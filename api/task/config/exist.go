package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
)

func (s *Server) ExistTaskConfigConds(ctx context.Context, in *npool.ExistTaskConfigCondsRequest) (*npool.ExistTaskConfigCondsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTaskConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTaskConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTaskConfigConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTaskConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTaskConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTaskConfigCondsResponse{
		Info: exist,
	}, nil
}
