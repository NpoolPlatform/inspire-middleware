package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	couponcoin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCouponCoin(ctx context.Context, in *npool.GetCouponCoinRequest) (*npool.GetCouponCoinResponse, error) {
	handler, err := couponcoin1.NewHandler(
		ctx,
		couponcoin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCouponCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponCoin",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCouponCoins(ctx context.Context, in *npool.GetCouponCoinsRequest) (*npool.GetCouponCoinsResponse, error) {
	handler, err := couponcoin1.NewHandler(
		ctx,
		couponcoin1.WithConds(in.GetConds()),
		couponcoin1.WithOffset(in.GetOffset()),
		couponcoin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponCoins",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCouponCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponCoins",
			"In", in,
			"Err", err,
		)
		return &npool.GetCouponCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
