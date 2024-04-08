package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppCommissionConfig(ctx context.Context, in *npool.CreateAppCommissionConfigRequest) (*npool.CreateAppCommissionConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppCommissionConfig",
			"In", in,
		)
		return &npool.CreateAppCommissionConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithThresholdAmount(req.ThresholdAmount, true),
		config1.WithAmountOrPercent(req.AmountOrPercent, true),
		config1.WithStartAt(req.StartAt, true),
		config1.WithInvites(req.Invites, true),
		config1.WithSettleType(req.SettleType, true),
		config1.WithDisabled(req.Disabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppCommissionConfigResponse{
		Info: info,
	}, nil
}
