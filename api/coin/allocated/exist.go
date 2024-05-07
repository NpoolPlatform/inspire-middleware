package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
)

func (s *Server) ExistCoinAllocatedConds(ctx context.Context, in *npool.ExistCoinAllocatedCondsRequest) (*npool.ExistCoinAllocatedCondsResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinAllocatedConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinAllocatedCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCoinAllocatedConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinAllocatedConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinAllocatedCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCoinAllocatedCondsResponse{
		Info: exist,
	}, nil
}
