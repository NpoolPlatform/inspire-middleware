package scope

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	scope1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/scope"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteScope(ctx context.Context, in *npool.DeleteScopeRequest) (*npool.DeleteScopeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteScope",
			"In", in,
		)
		return &npool.DeleteScopeResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := scope1.NewHandler(
		ctx,
		scope1.WithID(req.ID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScope",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteScope(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScope",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteScopeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteScopeResponse{
		Info: info,
	}, nil
}
