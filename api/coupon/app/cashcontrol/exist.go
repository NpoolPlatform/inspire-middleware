package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cashcontrol1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCashControlConds(ctx context.Context, in *npool.ExistCashControlCondsRequest) (*npool.ExistCashControlCondsResponse, error) {
	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCashControlConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCashControlCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCashControlConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCashControlConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCashControlCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCashControlCondsResponse{
		Info: info,
	}, nil
}
