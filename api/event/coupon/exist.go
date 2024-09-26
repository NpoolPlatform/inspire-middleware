package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"
)

func (s *Server) ExistEventCouponConds(ctx context.Context, in *npool.ExistEventCouponCondsRequest) (*npool.ExistEventCouponCondsResponse, error) {
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventCouponConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCouponCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistEventCouponConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistEventCouponConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistEventCouponCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistEventCouponCondsResponse{
		Info: exist,
	}, nil
}
