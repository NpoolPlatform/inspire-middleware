package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAppGoodCommissionConfig(ctx context.Context, in *npool.UpdateAppGoodCommissionConfigRequest) (*npool.UpdateAppGoodCommissionConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAppGoodCommissionConfig",
			"In", in,
		)
		return &npool.UpdateAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, false),
		config1.WithEntID(req.EntID, false),
		config1.WithAmountOrPercent(req.AmountOrPercent, false),
		config1.WithStartAt(req.StartAt, false),
		config1.WithThresholdAmount(req.ThresholdAmount, false),
		config1.WithInvites(req.Invites, false),
		config1.WithDisabled(req.Disabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateAppGoodCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppGoodCommissionConfigResponse{
		Info: info,
	}, nil
}
