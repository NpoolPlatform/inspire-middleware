package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistScopeConds(ctx context.Context, in *npool.ExistScopeCondsRequest) (*npool.ExistScopeCondsResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistScopeConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistScopeCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistScopeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistScopeConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistScopeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistScopeCondsResponse{
		Info: info,
	}, nil
}
