package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCommission(ctx context.Context, in *npool.CreateCommissionRequest) (*npool.CreateCommissionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCommission",
			"In", in,
		)
		return &npool.CreateCommissionResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := commission1.NewHandler(
		ctx,
		commission1.WithID(req.ID),
		commission1.WithAppID(req.AppID),
		commission1.WithUserID(req.UserID),
		commission1.WithGoodID(req.GoodID),
		commission1.WithSettleType(req.SettleType),
		commission1.WithSettleMode(req.SettleMode),
		commission1.WithSettleAmount(req.SettleAmount),
		commission1.WithSettleInterval(req.SettleInterval),
		commission1.WithAmountOrPercent(req.AmountOrPercent),
		commission1.WithStartAt(req.StartAt),
		commission1.WithThreshold(req.Threshold),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCommission",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCommission(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCommission",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCommissionResponse{
		Info: info,
	}, nil
}
