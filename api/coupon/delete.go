package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoupon(ctx context.Context, in *npool.DeleteCouponRequest) (*npool.DeleteCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCoupon",
			"In", in,
		)
		return &npool.DeleteCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCouponResponse{
		Info: info,
	}, nil
}
