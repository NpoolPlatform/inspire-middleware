package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAppCommissionConfigConds(ctx context.Context, in *npool.ExistAppCommissionConfigCondsRequest) (*npool.ExistAppCommissionConfigCondsResponse, error) {
	handler, err := commissionconfig1.NewHandler(
		ctx,
		commissionconfig1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppCommissionConfigConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppCommissionConfigCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCommissionConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppCommissionConfigConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistAppCommissionConfigCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCommissionConfigCondsResponse{
		Info: info,
	}, nil
}
