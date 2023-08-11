package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteEvent(ctx context.Context, in *npool.DeleteEventRequest) (*npool.DeleteEventResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteEvent",
			"In", in,
		)
		return &npool.DeleteEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := event1.NewHandler(
		ctx,
		event1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEvent",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	info, err := handler.DeleteEvent(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEvent",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	return &npool.DeleteEventResponse{
		Info: info,
	}, nil
}
