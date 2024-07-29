package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEventCoupon(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateEventCoupon",
			"In", in,
		)
		return &npool.CreateEventCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithEntID(req.EntID, false),
		coupon1.WithAppID(req.AppID, true),
		coupon1.WithEventID(req.EventID, true),
		coupon1.WithCouponID(req.CouponID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateEventCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateEventCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateEventCouponResponse{
		Info: nil,
	}, nil
}
