package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

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
		statement1.WithEntID(req.EntID, false),
		statement1.WithAppID(req.AppID, true),
		statement1.WithUserID(req.UserID, true),
		statement1.WithDirectContributorID(req.DirectContributorID, true),
		statement1.WithGoodID(req.GoodID, true),
		statement1.WithAppGoodID(req.AppGoodID, true),
		statement1.WithOrderID(req.OrderID, true),
		statement1.WithSelfOrder(req.SelfOrder, false),
		statement1.WithPaymentID(req.PaymentID, true),
		statement1.WithCoinTypeID(req.CoinTypeID, true),
		statement1.WithPaymentCoinTypeID(req.PaymentCoinTypeID, true),
		statement1.WithPaymentCoinUSDCurrency(req.PaymentCoinUSDCurrency, true),
		statement1.WithUnits(req.Units, true),
		statement1.WithAmount(req.Amount, true),
		statement1.WithUSDAmount(req.USDAmount, true),
		statement1.WithCommission(req.Commission, true),
		statement1.WithAppConfigID(req.AppConfigID, true),
		statement1.WithCommissionConfigID(req.CommissionConfigID, true),
		statement1.WithCommissionConfigType(req.CommissionConfigType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

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

func (s *Server) CreateStatements(ctx context.Context, in *npool.CreateStatementsRequest) (*npool.CreateStatementsResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"CreateStatements",
			"In", in,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(reqs, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStatementsResponse{
		Infos: infos,
	}, nil
}
