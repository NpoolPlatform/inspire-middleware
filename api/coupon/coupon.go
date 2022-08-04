package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetCoupon(ctx context.Context, in *npool.GetCouponRequest) (*npool.GetCouponResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetCoupon", "ID", in.GetID(), "error", err)
		return &npool.GetCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetType() {
	case npool.CouponType_FixAmount:
	case npool.CouponType_Discount:
	case npool.CouponType_SpecialOffer:
	case npool.CouponType_ThresholdReduction:
	default:
		logger.Sugar().Errorw("GetCoupon", "Type", in.GetType())
		return &npool.GetCouponResponse{}, status.Error(codes.InvalidArgument, "invalid coupon type")
	}

	info, err := coupon1.GetCoupon(ctx, in.GetID(), in.GetType())
	if err != nil {
		logger.Sugar().Errorw("GetCoupon", "error", err)
		return &npool.GetCouponResponse{}, status.Error(codes.Internal, "fail get coupon")
	}

	return &npool.GetCouponResponse{
		Info: info,
	}, nil
}

func (s *Server) GetManyCoupons(ctx context.Context, in *npool.GetManyCouponsRequest) (*npool.GetManyCouponsResponse, error) {
	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetManyCoupons", "ID", id, "error", err)
			return &npool.GetManyCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	switch in.GetType() {
	case npool.CouponType_FixAmount:
	case npool.CouponType_Discount:
	case npool.CouponType_SpecialOffer:
	case npool.CouponType_ThresholdReduction:
	default:
		logger.Sugar().Errorw("GetManyCoupons", "Type", in.GetType())
		return &npool.GetManyCouponsResponse{}, status.Error(codes.InvalidArgument, "invalid coupon type")
	}

	infos, err := coupon1.GetManyCoupons(ctx, in.GetIDs(), in.GetType())
	if err != nil {
		logger.Sugar().Errorw("GetManyCoupons", "error", err)
		return &npool.GetManyCouponsResponse{}, status.Error(codes.Internal, "fail get coupon")
	}

	return &npool.GetManyCouponsResponse{
		Infos: infos,
	}, nil
}
