package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	config1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppGoodCommissionConfig(ctx context.Context, in *npool.CreateAppGoodCommissionConfigRequest) (*npool.CreateAppGoodCommissionConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppGoodCommissionConfig",
			"In", in,
		)
		return &npool.CreateAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithGoodID(req.GoodID, true),
		config1.WithAppGoodID(req.AppGoodID, true),
		config1.WithThresholdAmount(req.ThresholdAmount, true),
		config1.WithAmountOrPercent(req.AmountOrPercent, true),
		config1.WithStartAt(req.StartAt, true),
		config1.WithInvites(req.Invites, true),
		config1.WithSettleType(req.SettleType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppGoodCommissionConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCommissionConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppGoodCommissionConfig",
			"In", in,
			"Err", err,
		)
		return &npool.CreateAppGoodCommissionConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppGoodCommissionConfigResponse{
		Info: info,
	}, nil
}
