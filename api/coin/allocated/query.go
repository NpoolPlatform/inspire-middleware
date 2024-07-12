package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
)

func (s *Server) GetCoinAllocated(ctx context.Context, in *npool.GetCoinAllocatedRequest) (*npool.GetCoinAllocatedResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetCoinAllocated(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCoinAllocatedResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinAllocateds(ctx context.Context, in *npool.GetCoinAllocatedsRequest) (*npool.GetCoinAllocatedsResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithConds(in.GetConds()),
		allocated1.WithOffset(in.GetOffset()),
		allocated1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinAllocateds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinAllocatedsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetCoinAllocateds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinAllocateds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinAllocatedsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCoinAllocatedsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
