package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEventCoin(ctx context.Context, in *npool.GetEventCoinRequest) (*npool.GetEventCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetEventCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoin",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEventCoins(ctx context.Context, in *npool.GetEventCoinsRequest) (*npool.GetEventCoinsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
		coin1.WithOffset(in.GetOffset()),
		coin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoins",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetEventCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoins",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
