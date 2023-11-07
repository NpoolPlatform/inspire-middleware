package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEvent(ctx context.Context, in *npool.GetEventRequest) (*npool.GetEventResponse, error) {
	handler, err := event1.NewHandler(
		ctx,
		event1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEvent",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetEvent(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEvent",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEvents(ctx context.Context, in *npool.GetEventsRequest) (*npool.GetEventsResponse, error) {
	handler, err := event1.NewHandler(
		ctx,
		event1.WithConds(in.GetConds()),
		event1.WithOffset(in.GetOffset()),
		event1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEvents",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetEvents(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEvents",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
