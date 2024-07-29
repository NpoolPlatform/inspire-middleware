package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/credit/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/credit/allocated"
)

func (s *Server) GetCreditAllocated(ctx context.Context, in *npool.GetCreditAllocatedRequest) (*npool.GetCreditAllocatedResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.GetCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetCreditAllocated(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.GetCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCreditAllocatedResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCreditAllocateds(ctx context.Context, in *npool.GetCreditAllocatedsRequest) (*npool.GetCreditAllocatedsResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithConds(in.GetConds()),
		allocated1.WithOffset(in.GetOffset()),
		allocated1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCreditAllocateds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCreditAllocatedsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetCreditAllocateds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCreditAllocateds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCreditAllocatedsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCreditAllocatedsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
