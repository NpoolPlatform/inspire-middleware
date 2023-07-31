package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement/statement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateStatement(ctx context.Context, in *npool.CreateStatementRequest) (*npool.CreateStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID),
		statement1.WithAppID(req.AppID),
		statement1.WithUserID(req.UserID),
		statement1.WithDirectContributorID(req.DirectContributorID),
		statement1.WithGoodID(req.GoodID),
		statement1.WithOrderID(req.OrderID),
		statement1.WithSelfOrder(req.SelfOrder),
		statement1.WithPaymentID(req.PaymentID),
		statement1.WithCoinTypeID(req.CoinTypeID),
		statement1.WithPaymentCoinTypeID(req.PaymentCoinTypeID),
		statement1.WithPaymentCoinUSDCurrency(req.PaymentCoinUSDCurrency),
		statement1.WithUnits(req.Units),
		statement1.WithAmount(req.Amount),
		statement1.WithUSDAmount(req.USDAmount),
		statement1.WithCommission(req.Commission),
	)

	info, err := handler.CreateStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStatementResponse{
		Info: info,
	}, nil
}
