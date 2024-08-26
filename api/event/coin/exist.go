package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	coin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"
)

func (s *Server) ExistEventCoinConds(ctx context.Context, in *npool.ExistEventCoinCondsRequest) (*npool.ExistEventCoinCondsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCoinCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistEventCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCoinCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistEventCoinCondsResponse{
		Info: exist,
	}, nil
}
