package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAppGoodScope(ctx context.Context, in *npool.GetAppGoodScopeRequest) (*npool.GetAppGoodScopeResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodScopeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAppGoodScope(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodScope",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodScopeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodScopeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppGoodScopes(ctx context.Context, in *npool.GetAppGoodScopesRequest) (*npool.GetAppGoodScopesResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithConds(in.GetConds()),
		scope1.WithOffset(in.GetOffset()),
		scope1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodScopes",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodScopesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAppGoodScopes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppGoodScopes",
			"In", in,
			"Err", err,
		)
		return &npool.GetAppGoodScopesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppGoodScopesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
