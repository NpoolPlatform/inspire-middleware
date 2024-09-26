package allocated

import (
	"context"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/credit/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/credit/allocated"
)

func (s *Server) CreateCreditAllocated(ctx context.Context, in *npool.CreateCreditAllocatedRequest) (*npool.CreateCreditAllocatedResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCreditAllocated",
			"In", in,
		)
		return &npool.CreateCreditAllocatedResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(req.EntID, false),
		allocated1.WithAppID(req.AppID, true),
		allocated1.WithUserID(req.UserID, true),
		allocated1.WithValue(req.Value, true),
		allocated1.WithExtra(req.Extra, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateCreditAllocated(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateCreditAllocated",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCreditAllocatedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCreditAllocatedResponse{}, nil
}
