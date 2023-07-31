package achivement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExpropriateAchivement(ctx context.Context, in *npool.ExpropriateAchivementRequest) (*npool.ExpropriateAchivementResponse, error) {
	handler, err := achivement1.NewHandler(
		ctx,
		achivement1.WithOrderID(&in.OrderID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExpropriateAchivement",
			"In", in,
			"Err", err,
		)
		return &npool.ExpropriateAchivementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.ExpropriateAchivement(ctx); err != nil {
		logger.Sugar().Errorw(
			"ExpropriateAchivement",
			"In", in,
			"Err", err,
		)
		return &npool.ExpropriateAchivementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExpropriateAchivementResponse{}, nil
}
