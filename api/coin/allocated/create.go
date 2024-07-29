package allocated

import (
	"context"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
)

func (s *Server) CreateCoinAllocated(ctx context.Context, in *npool.CreateCoinAllocatedRequest) (*npool.CreateCoinAllocatedResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoinAllocated",
			"In", in,
		)
		return &npool.CreateCoinAllocatedResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(req.EntID, false),
		allocated1.WithAppID(req.AppID, true),
		allocated1.WithCoinConfigID(req.CoinConfigID, true),
		allocated1.WithCoinTypeID(req.CoinTypeID, true),
		allocated1.WithUserID(req.UserID, true),
		allocated1.WithValue(req.Value, true),
		allocated1.WithExtra(req.Extra, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateCoinAllocated(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCoinAllocatedResponse{}, nil
}
