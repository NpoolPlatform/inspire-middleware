package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteEventCoin(ctx context.Context, in *npool.DeleteEventCoinRequest) (*npool.DeleteEventCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteEventCoin",
			"In", in,
		)
		return &npool.DeleteEventCoinResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID, false),
		coin1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventCoinResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	err = handler.DeleteEventCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventCoinResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	return &npool.DeleteEventCoinResponse{
		Info: nil,
	}, nil
}
