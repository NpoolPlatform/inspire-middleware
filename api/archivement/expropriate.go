package archivement

import (
	"context"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Expropriate(ctx context.Context, in *npool.ExpropriateRequest) (*npool.ExpropriateResponse, error) {
	if _, err := uuid.Parse(in.GetOrderID()); err != nil {
		logger.Sugar().Errorw("Expropriate", "OrderID", in.GetOrderID(), "error", err)
		return &npool.ExpropriateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := archivement1.Expropriate(ctx, in.GetOrderID()); err != nil {
		return &npool.ExpropriateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExpropriateResponse{}, nil
}
