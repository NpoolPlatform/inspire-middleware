package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	couponcoin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/coin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCouponCoinConds(ctx context.Context, in *npool.ExistCouponCoinCondsRequest) (*npool.ExistCouponCoinCondsResponse, error) {
	handler, err := couponcoin1.NewHandler(
		ctx,
		couponcoin1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCouponCoinConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCouponCoinCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCouponCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCouponCoinConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCouponCoinCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCouponCoinCondsResponse{
		Info: info,
	}, nil
}
