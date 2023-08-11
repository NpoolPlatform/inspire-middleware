package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneCommissions(ctx context.Context, in *npool.CloneCommissionsRequest) (*npool.CloneCommissionsResponse, error) {
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithAppID(&in.AppID),
		commission1.WithFromGoodID(&in.FromGoodID),
		commission1.WithToGoodID(&in.ToGoodID),
		commission1.WithScalePercent(&in.ScalePercent),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CloneCommissions",
			"In", in,
			"Err", err,
		)
		return &npool.CloneCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CloneCommissions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CloneCommissions",
			"In", in,
			"Err", err,
		)
		return &npool.CloneCommissionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CloneCommissionsResponse{}, nil
}
