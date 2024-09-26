package calculate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	calculate1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReconcileCalculate(ctx context.Context, in *npool.ReconcileCalculateRequest) (*npool.ReconcileCalculateResponse, error) {
	handler, err := calculate1.NewHandler(
		ctx,
		calculate1.WithOrderID(in.GetOrderID()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ReconcileCalculate",
			"In", in,
			"Err", err,
		)
		return &npool.ReconcileCalculateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.ReconcileCalculate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ReconcileCalculate",
			"In", in,
			"Err", err,
		)
		return &npool.ReconcileCalculateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ReconcileCalculateResponse{
		Infos: infos,
	}, nil
}
