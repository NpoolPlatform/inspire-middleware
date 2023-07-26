package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

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
		return &npool.CreateCouponResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithID(req.ID),
		allocated1.WithAppID(req.AppID),
		allocated1.WithCouponID(req.CouponID),
		allocated1.WithUserID(req.UserID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCouponResponse{
		Info: info,
	}, nil
}
