package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyCouponScope(ctx context.Context, in *npool.VerifyCouponScopeRequest) (*npool.VerifyCouponScopeResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithAppID(&in.AppID, true),
		scope1.WithGoodID(&in.GoodID, true),
		scope1.WithAppGoodID(&in.AppGoodID, true),
		scope1.WithCouponScope(&in.CouponScope, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyCouponScope",
			"In", in,
			"Err", err,
		)
		return &npool.VerifyCouponScopeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistAppGoodScopeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyCouponScope",
			"In", in,
			"Err", err,
		)
		return &npool.VerifyCouponScopeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.VerifyCouponScopeResponse{
		Info: info,
	}, nil
}
