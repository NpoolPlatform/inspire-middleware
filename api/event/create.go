package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEvent(ctx context.Context, in *npool.CreateEventRequest) (*npool.CreateEventResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateEvent",
			"In", in,
		)
		return &npool.CreateEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := event1.NewHandler(
		ctx,
		event1.WithID(req.ID),
		event1.WithAppID(req.AppID),
		event1.WithEventType(req.EventType),
		event1.WithCouponIDs(req.CouponIDs),
		event1.WithCredits(req.Credits),
		event1.WithCreditsPerUSD(req.CreditsPerUSD),
		event1.WithMaxConsecutive(req.MaxConsecutive),
		event1.WithGoodID(req.GoodID),
		event1.WithInviterLayers(req.InviterLayers),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEvent",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	info, err := handler.CreateEvent(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEvent",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	return &npool.CreateEventResponse{
		Info: info,
	}, nil
}
