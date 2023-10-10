package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateScope(ctx context.Context, in *npool.CreateScopeRequest) (*npool.CreateScopeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateScope",
			"In", in,
		)
		return &npool.CreateScopeResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithID(req.ID, false),
		scope1.WithAppID(req.AppID, true),
		scope1.WithAppGoodID(req.AppGoodID, false),
		scope1.WithCouponID(req.CouponID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScope",
			"In", in,
			"Err", err,
		)
		return &npool.CreateScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateScope(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScope",
			"In", in,
			"Err", err,
		)
		return &npool.CreateScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateScopeResponse{
		Info: info,
	}, nil
}
