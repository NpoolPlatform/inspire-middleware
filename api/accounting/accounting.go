package accounting

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	accounting1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/accounting"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Accounting(ctx context.Context, in *npool.AccountingRequest) (*npool.AccountingResponse, error) {
	handler, err := accounting1.NewHandler(
		ctx,
		accounting1.WithAppID(in.GetAppID()),
		accounting1.WithUserID(in.GetUserID()),
		accounting1.WithGoodID(in.GetGoodID()),
		accounting1.WithOrderID(in.GetOrderID()),
		accounting1.WithPaymentID(in.GetPaymentID()),
		accounting1.WithCoinTypeID(in.GetCoinTypeID()),
		accounting1.WithPaymentCoinTypeID(in.GetPaymentCoinTypeID()),
		accounting1.WithPaymentCoinUSDCurrency(in.GetPaymentCoinUSDCurrency()),
		accounting1.WithUnits(in.GetUnits()),
		accounting1.WithSettleType(in.GetSettleType()),
		accounting1.WithSettleMode(in.GetSettleMode()),
		accounting1.WithPaymentAmount(in.GetPaymentAmount()),
		accounting1.WithGoodValue(in.GetGoodValue()),
		accounting1.WithHasCommission(in.GetHasCommission()),
		accounting1.WithOrderCreatedAt(in.GetOrderCreatedAt()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Accounting",
			"In", in,
			"Err", err,
		)
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.Accounting(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Accounting",
			"In", in,
			"Err", err,
		)
		return &npool.AccountingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AccountingResponse{
		Infos: infos,
	}, nil
}
