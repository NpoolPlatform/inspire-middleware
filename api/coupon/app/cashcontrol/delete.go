package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cashcontrol1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCashControl(ctx context.Context, in *npool.DeleteCashControlRequest) (*npool.DeleteCashControlResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCashControl",
			"In", in,
		)
		return &npool.DeleteCashControlResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithID(req.ID, false),
		cashcontrol1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteCashControl(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCashControlResponse{
		Info: info,
	}, nil
}
