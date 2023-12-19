package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	couponcoin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCouponCoin(ctx context.Context, in *npool.CreateCouponCoinRequest) (*npool.CreateCouponCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCouponCoin",
			"In", in,
		)
		return &npool.CreateCouponCoinResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := couponcoin1.NewHandler(
		ctx,
		couponcoin1.WithEntID(req.EntID, false),
		couponcoin1.WithAppID(req.AppID, true),
		couponcoin1.WithCoinTypeID(req.CoinTypeID, true),
		couponcoin1.WithCouponID(req.CouponID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCouponCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateCouponCoinResponse{
		Info: info,
	}, nil
}
