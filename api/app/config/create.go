package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppConfig(ctx context.Context, in *npool.CreateAppConfigRequest) (*npool.CreateAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithEntID(req.EntID, false),
		commission1.WithAppID(req.AppID, true),
		commission1.WithSettleMode(req.SettleMode, true),
		commission1.WithSettleAmountType(req.SettleAmountType, true),
		commission1.WithSettleInterval(req.SettleInterval, true),
		commission1.WithCommissionType(req.CommissionType, true),
		commission1.WithSettleBenefit(req.SettleBenefit, true),
		commission1.WithStartAt(req.StartAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppConfigResponse{
		Info: info,
	}, nil
}
