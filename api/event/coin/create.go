package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEventCoin(ctx context.Context, in *npool.CreateEventCoinRequest) (*npool.CreateEventCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateEventCoin",
			"In", in,
		)
		return &npool.CreateEventCoinResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithEntID(req.EntID, false),
		coin1.WithAppID(req.AppID, true),
		coin1.WithEventID(req.EventID, true),
		coin1.WithCoinConfigID(req.CoinConfigID, true),
		coin1.WithCoinValue(req.CoinValue, true),
		coin1.WithCoinPreUSD(req.CoinPreUSD, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateEventCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateEventCoinResponse{
		Info: nil,
	}, nil
}
