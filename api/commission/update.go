package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCommission(ctx context.Context, in *npool.UpdateCommissionRequest) (*npool.UpdateCommissionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCommission",
			"In", in,
		)
		return &npool.UpdateCommissionResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithID(req.ID, true),
		commission1.WithAmountOrPercent(req.AmountOrPercent, false),
		commission1.WithStartAt(req.StartAt, false),
		commission1.WithThreshold(req.Threshold, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCommission",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateCommission(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCommission",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCommissionResponse{}, nil
}
