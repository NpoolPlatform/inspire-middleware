package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/credit/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/credit/allocated"
)

func (s *Server) ExistCreditAllocatedConds(ctx context.Context, in *npool.ExistCreditAllocatedCondsRequest) (*npool.ExistCreditAllocatedCondsResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCreditAllocatedConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCreditAllocatedCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCreditAllocatedConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCreditAllocatedConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCreditAllocatedCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCreditAllocatedCondsResponse{
		Info: exist,
	}, nil
}
