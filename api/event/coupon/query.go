package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEventCoupon(ctx context.Context, in *npool.GetEventCouponRequest) (*npool.GetEventCouponResponse, error) {
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetEventCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventCouponResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEventCoupons(ctx context.Context, in *npool.GetEventCouponsRequest) (*npool.GetEventCouponsResponse, error) {
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithConds(in.GetConds()),
		coupon1.WithOffset(in.GetOffset()),
		coupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoupons",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetEventCoupons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetEventCoupons",
			"In", in,
			"Err", err,
		)
		return &npool.GetEventCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
