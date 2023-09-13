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
		commission1.WithID(req.ID, false),
		commission1.WithAppID(req.AppID, true),
		commission1.WithUserID(req.UserID, true),
		commission1.WithGoodID(req.GoodID, true),
		commission1.WithAppGoodID(req.AppGoodID, true),
		commission1.WithSettleType(req.SettleType, true),
		commission1.WithSettleMode(req.SettleMode, true),
		commission1.WithSettleAmountType(req.SettleAmountType, true),
		commission1.WithSettleInterval(req.SettleInterval, true),
		commission1.WithAmountOrPercent(req.AmountOrPercent, true),
		commission1.WithStartAt(req.StartAt, true),
		commission1.WithThreshold(req.Threshold, false),
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
