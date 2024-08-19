package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateEvent(ctx context.Context, in *npool.UpdateEventRequest) (*npool.UpdateEventResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateEvent",
			"In", in,
		)
		return &npool.UpdateEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := event1.NewHandler(
		ctx,
		event1.WithID(req.ID, false),
		event1.WithEntID(req.EntID, false),
		event1.WithCredits(req.Credits, false),
		event1.WithCreditsPerUSD(req.CreditsPerUSD, false),
		event1.WithMaxConsecutive(req.MaxConsecutive, false),
		event1.WithInviterLayers(req.InviterLayers, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEvent",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateEvent(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEvent",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.UpdateEventResponse{
		Info: nil,
	}, nil
}
