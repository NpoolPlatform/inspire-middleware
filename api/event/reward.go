package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RewardEvent(ctx context.Context, in *npool.RewardEventRequest) (*npool.RewardEventResponse, error) {
	handler, err := event1.NewHandler(
		ctx,
		event1.WithAppID(&in.AppID),
		event1.WithUserID(&in.UserID),
		event1.WithEventType(&in.EventType),
		event1.WithGoodID(in.GoodID),
		event1.WithConsecutive(&in.Consecutive),
		event1.WithAmount(&in.Amount),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"RewardEvent",
			"In", in,
			"Error", err,
		)
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.RewardEvent(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"RewardEvent",
			"In", in,
			"Error", err,
		)
		return &npool.RewardEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.RewardEventResponse{
		Infos: infos,
	}, nil
}
