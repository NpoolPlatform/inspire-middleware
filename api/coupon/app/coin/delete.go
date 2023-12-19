package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	couponcoin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCouponCoin(ctx context.Context, in *npool.DeleteCouponCoinRequest) (*npool.DeleteCouponCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCouponCoin",
			"In", in,
		)
		return &npool.DeleteCouponCoinResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := couponcoin1.NewHandler(
		ctx,
		couponcoin1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteCouponCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteCouponCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCouponCoinResponse{
		Info: info,
	}, nil
}
