package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoupon(ctx context.Context, in *npool.UpdateCouponRequest) (*npool.UpdateCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCoupon",
			"In", in,
		)
		return &npool.UpdateCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithID(req.ID, true),
		coupon1.WithDenomination(req.Denomination, false),
		coupon1.WithCirculation(req.Circulation, false),
		coupon1.WithStartAt(req.StartAt, false),
		coupon1.WithDurationDays(req.DurationDays, false),
		coupon1.WithMessage(req.Message, false),
		coupon1.WithName(req.Name, false),
		coupon1.WithRandom(req.Random, false),
		coupon1.WithCouponScope(req.CouponScope, false),
		coupon1.WithThreshold(req.Threshold, false),
		coupon1.WithAllocated(req.Allocated, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCouponResponse{
		Info: info,
	}, nil
}
