package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cashcontrol1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCashControl(ctx context.Context, in *npool.GetCashControlRequest) (*npool.GetCashControlResponse, error) {
	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.GetCashControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCashControl(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCashControl",
			"In", in,
			"Err", err,
		)
		return &npool.GetCashControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCashControlResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCashControls(ctx context.Context, in *npool.GetCashControlsRequest) (*npool.GetCashControlsResponse, error) {
	handler, err := cashcontrol1.NewHandler(
		ctx,
		cashcontrol1.WithConds(in.GetConds()),
		cashcontrol1.WithOffset(in.GetOffset()),
		cashcontrol1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCashControls",
			"In", in,
			"Err", err,
		)
		return &npool.GetCashControlsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCashControls(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCashControls",
			"In", in,
			"Err", err,
		)
		return &npool.GetCashControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCashControlsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
