package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoupon(ctx context.Context, in *npool.GetCouponRequest) (*npool.GetCouponResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(&in.EntID, true),
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

func (s *Server) GetCoupons(ctx context.Context, in *npool.GetCouponsRequest) (*npool.GetCouponsResponse, error) {
	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithConds(in.GetConds()),
		allocated1.WithOffset(in.GetOffset()),
		allocated1.WithLimit(in.GetLimit()),
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
