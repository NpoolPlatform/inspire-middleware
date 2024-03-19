package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneAppGoodCommissionConfigs(ctx context.Context, in *npool.CloneAppGoodCommissionConfigsRequest) (*npool.CloneAppGoodCommissionConfigsResponse, error) {
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithAppID(&in.AppID, true),
		commission1.WithFromGoodID(&in.FromGoodID, true),
		commission1.WithToGoodID(&in.ToGoodID, true),
		commission1.WithFromAppGoodID(&in.FromAppGoodID, true),
		commission1.WithToAppGoodID(&in.ToAppGoodID, true),
		commission1.WithScalePercent(&in.ScalePercent, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CloneCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.CloneAppGoodCommissionConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CloneCommissionConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CloneCommissionConfigs",
			"In", in,
			"Err", err,
		)
		return &npool.CloneAppGoodCommissionConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CloneAppGoodCommissionConfigsResponse{}, nil
}
