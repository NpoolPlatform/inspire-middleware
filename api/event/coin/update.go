package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateEventCoin(ctx context.Context, in *npool.UpdateEventCoinRequest) (*npool.UpdateEventCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateEventCoin",
			"In", in,
		)
		return &npool.UpdateEventCoinResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID, true),
		coin1.WithCoinValue(req.CoinValue, false),
		coin1.WithCoinPreUSD(req.CoinPreUSD, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateEventCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateEventCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateEventCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.UpdateEventCoinResponse{
		Info: nil,
	}, nil
}
