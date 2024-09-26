package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func (s *Server) ExistEventConds(ctx context.Context, in *npool.ExistEventCondsRequest) (*npool.ExistEventCondsResponse, error) {
	handler, err := event1.NewHandler(
		ctx,
		event1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistEventConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistEventCondsResponse{
		Info: exist,
	}, nil
}
