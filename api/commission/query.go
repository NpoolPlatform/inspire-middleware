package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCommission(ctx context.Context, in *npool.GetCommissionRequest) (*npool.GetCommissionResponse, error) {
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommission",
			"In", in,
			"Err", err,
		)
		return &npool.GetCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCommission(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommission",
			"In", in,
			"Err", err,
		)
		return &npool.GetCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommissionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCommissions(ctx context.Context, in *npool.GetCommissionsRequest) (*npool.GetCommissionsResponse, error) {
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithConds(in.GetConds()),
		commission1.WithOffset(in.GetOffset()),
		commission1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissions",
			"In", in,
			"Err", err,
		)
		return &npool.GetCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCommissions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCommissions",
			"In", in,
			"Err", err,
		)
		return &npool.GetCommissionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommissionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
