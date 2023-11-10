package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAppGoodScope(ctx context.Context, in *npool.DeleteAppGoodScopeRequest) (*npool.DeleteAppGoodScopeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAppGoodScope",
			"In", in,
		)
		return &npool.DeleteAppGoodScopeResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppGoodScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteAppGoodScope(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAppGoodScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteAppGoodScopeResponse{
		Info: info,
	}, nil
}
