package calculate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	calculate1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Calculate(ctx context.Context, in *npool.CalculateRequest) (*npool.CalculateResponse, error) {
	handler, err := calculate1.NewHandler(
		ctx,
		calculate1.WithAppID(in.GetAppID()),
		calculate1.WithUserID(in.GetUserID()),
		calculate1.WithGoodID(in.GetGoodID()),
		calculate1.WithAppGoodID(in.GetAppGoodID()),
		calculate1.WithOrderID(in.GetOrderID()),
		calculate1.WithGoodCoinTypeID(in.GetGoodCoinTypeID()),
		calculate1.WithUnits(in.GetUnits()),
		calculate1.WithPaymentAmountUSD(in.GetPaymentAmountUSD()),
		calculate1.WithSettleType(in.GetSettleType()),
		calculate1.WithGoodValueUSD(in.GetGoodValueUSD()),
		calculate1.WithHasCommission(in.GetHasCommission()),
		calculate1.WithOrderCreatedAt(in.GetOrderCreatedAt()),
		calculate1.WithPaymentCoinTypeID(in.GetPaymentCoinTypeID()),
		calculate1.WithPaymentCoinUSDCurrency(in.GetPaymentCoinUSDCurrency()),
		calculate1.WithPaymentAmount(in.GetPaymentAmount()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Calculate",
			"In", in,
			"Err", err,
		)
		return &npool.CalculateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.Calculate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Calculate",
			"In", in,
			"Err", err,
		)
		return &npool.CalculateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CalculateResponse{
		Infos: infos,
	}, nil
}
