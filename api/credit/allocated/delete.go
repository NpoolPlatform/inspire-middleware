package allocated

import (
	"context"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/credit/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/credit/allocated"
)

func (s *Server) DeleteCreditAllocated(ctx context.Context, in *npool.DeleteCreditAllocatedRequest) (*npool.DeleteCreditAllocatedResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCreditAllocated",
			"In", in,
		)
		return &npool.DeleteCreditAllocatedResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithID(req.ID, false),
		allocated1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteCreditAllocated(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCreditAllocatedResponse{}, nil
}
