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
		coupon1.WithID(req.ID),
		coupon1.WithCouponType(req.CouponType),
		coupon1.WithAppID(req.AppID),
		coupon1.WithDenomination(req.Denomination),
		coupon1.WithCirculation(req.Circulation),
		coupon1.WithIssuedBy(req.IssuedBy),
		coupon1.WithStartAt(req.StartAt),
		coupon1.WithDurationDays(req.DurationDays),
		coupon1.WithMessage(req.Message),
		coupon1.WithName(req.Name),
		coupon1.WithUserID(req.UserID),
		coupon1.WithGoodID(req.GoodID),
		coupon1.WithThreshold(req.Threshold),
		coupon1.WithCouponConstraint(req.CouponConstraint),
		coupon1.WithRandom(req.Random),
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
