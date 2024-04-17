package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCommissionConds(
	ctx context.Context,
	in *npool.ExistCommissionCondsRequest,
) (*npool.ExistCommissionCondsResponse, error) {
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCommissionConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCommissionCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCommissions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCommissionConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistCommissionCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCommissionCondsResponse{
		Info: info,
	}, nil
}
