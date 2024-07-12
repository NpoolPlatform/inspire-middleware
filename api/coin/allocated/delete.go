package allocated

import (
	"context"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
)

func (s *Server) DeleteCoinAllocated(ctx context.Context, in *npool.DeleteCoinAllocatedRequest) (*npool.DeleteCoinAllocatedResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCoinAllocated",
			"In", in,
		)
		return &npool.DeleteCoinAllocatedResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithID(req.ID, false),
		allocated1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteCoinAllocated(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCoinAllocatedResponse{}, nil
}
