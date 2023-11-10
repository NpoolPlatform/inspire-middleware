package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppGoodScope(ctx context.Context, in *npool.CreateAppGoodScopeRequest) (*npool.CreateAppGoodScopeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppGoodScope",
			"In", in,
		)
		return &npool.CreateAppGoodScopeResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithEntID(req.EntID, false),
		scope1.WithAppID(req.AppID, true),
		scope1.WithAppGoodID(req.AppGoodID, true),
		scope1.WithCouponID(req.CouponID, true),
		scope1.WithCouponScope(req.CouponScope, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppGoodScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAppGoodScope(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppGoodScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateAppGoodScopeResponse{
		Info: info,
	}, nil
}
