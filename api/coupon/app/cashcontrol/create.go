package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cashcontrol1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCashControl(ctx context.Context, in *npool.CreateCashControlRequest) (*npool.CreateCashControlResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCashControl",
			"In", in,
		)
		return &npool.CreateCashControlResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithEntID(req.EntID, false),
		cashcontrol1.WithAppID(req.AppID, true),
		cashcontrol1.WithCouponID(req.CouponID, true),
		cashcontrol1.WithControlType(req.ControlType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCashControl(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCashControlResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateCashControlResponse{
		Info: info,
	}, nil
}
