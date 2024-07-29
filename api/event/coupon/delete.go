package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteEventCoupon(ctx context.Context, in *npool.DeleteEventCouponRequest) (*npool.DeleteEventCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteEventCoupon",
			"In", in,
		)
		return &npool.DeleteEventCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithID(req.ID, false),
		coupon1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	err = handler.DeleteEventCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteEventCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	return &npool.DeleteEventCouponResponse{
		Info: nil,
	}, nil
}
