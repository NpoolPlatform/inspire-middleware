package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

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

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithID(req.ID, true),
		allocated1.WithUsed(req.Used, false),
		allocated1.WithUsedByOrderID(req.UsedByOrderID, false),
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
