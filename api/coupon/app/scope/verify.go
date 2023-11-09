package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyCouponScopes(ctx context.Context, in *npool.VerifyCouponScopesRequest) (*npool.VerifyCouponScopesResponse, error) {
	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyCouponScopes",
			"In", in,
			"Err", err,
		)
		return &npool.VerifyCouponScopesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.VerifyCouponScopes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyCouponScopes",
			"In", in,
			"Err", err,
		)
		return &npool.VerifyCouponScopesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.VerifyCouponScopesResponse{}, nil
}
