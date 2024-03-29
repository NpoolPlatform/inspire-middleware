package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoupon(ctx context.Context, in *npool.CreateCouponRequest) (*npool.CreateCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithEntID(req.EntID, false),
		coupon1.WithAppID(req.AppID, true),
		coupon1.WithName(req.Name, true),
		coupon1.WithMessage(req.Message, true),
		coupon1.WithCouponType(req.CouponType, true),
		coupon1.WithDenomination(req.Denomination, true),
		coupon1.WithCirculation(req.Circulation, true),
		coupon1.WithIssuedBy(req.IssuedBy, true),
		coupon1.WithStartAt(req.StartAt, true),
		coupon1.WithEndAt(req.EndAt, true),
		coupon1.WithDurationDays(req.DurationDays, true),
		coupon1.WithThreshold(req.Threshold, false),
		coupon1.WithCouponConstraint(req.CouponConstraint, false),
		coupon1.WithCouponScope(req.CouponScope, false),
		coupon1.WithRandom(req.Random, false),
		coupon1.WithCashableProbability(req.CashableProbability, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCouponResponse{
		Info: info,
	}, nil
}
