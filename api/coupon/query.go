package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoupons(ctx context.Context, in *npool.GetCouponsRequest) (*npool.GetCouponsResponse, error) {
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithConds(in.GetConds()),
		coupon1.WithOffset(in.GetOffset()),
		coupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoupons",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoupons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoupons",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCoupon(ctx context.Context, in *npool.GetCouponRequest) (*npool.GetCouponResponse, error) {
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponResponse{
		Info: info,
	}, nil
}
