package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cashcontrol1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCashControl(ctx context.Context, in *npool.UpdateCashControlRequest) (*npool.UpdateCashControlResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCashControl",
			"In", in,
		)
		return &npool.UpdateCashControlResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithID(req.ID, true),
		cashcontrol1.WithValue(req.Value, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateCashControl(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateCashControlResponse{
		Info: info,
	}, nil
}
