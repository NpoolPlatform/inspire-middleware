package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAppGoodScopeConds(ctx context.Context, in *npool.ExistAppGoodScopeCondsRequest) (*npool.ExistAppGoodScopeCondsResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppGoodScopeConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppGoodScopeCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistAppGoodScopeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppGoodScopeConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppGoodScopeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppGoodScopeCondsResponse{
		Info: info,
	}, nil
}
