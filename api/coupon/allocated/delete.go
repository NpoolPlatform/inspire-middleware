package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

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
		return &npool.DeleteCouponResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithID(req.ID, false),
		allocated1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCouponResponse{
		Info: info,
	}, nil
}
