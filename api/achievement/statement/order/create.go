//nolint:dupl
package orderstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orderstatement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"

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
	handler, err := orderstatement1.NewHandler(
		ctx,
		orderstatement1.WithEntID(req.EntID, false),
		orderstatement1.WithAppID(req.AppID, true),
		orderstatement1.WithUserID(req.UserID, true),
		orderstatement1.WithGoodID(req.GoodID, true),
		orderstatement1.WithAppGoodID(req.AppGoodID, true),
		orderstatement1.WithOrderID(req.OrderID, true),
		orderstatement1.WithOrderUserID(req.OrderUserID, true),
		orderstatement1.WithDirectContributorID(req.DirectContributorID, true),
		orderstatement1.WithGoodCoinTypeID(req.GoodCoinTypeID, true),
		orderstatement1.WithUnits(req.Units, true),
		orderstatement1.WithGoodValueUSD(req.GoodValueUSD, true),
		orderstatement1.WithPaymentAmountUSD(req.PaymentAmountUSD, true),
		orderstatement1.WithCommissionAmountUSD(req.CommissionAmountUSD, true),
		orderstatement1.WithAppConfigID(req.AppConfigID, true),
		orderstatement1.WithCommissionConfigID(req.CommissionConfigID, false),
		orderstatement1.WithCommissionConfigType(req.CommissionConfigType, true),
		orderstatement1.WithPaymentStatements(req.PaymentStatements, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.CreateStatement(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStatementResponse{}, nil
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
	multiHandler := &orderstatement1.MultiHandler{}
	for _, req := range reqs {
		handler, err := orderstatement1.NewHandler(
			ctx,
			orderstatement1.WithEntID(req.EntID, false),
			orderstatement1.WithAppID(req.AppID, true),
			orderstatement1.WithUserID(req.UserID, true),
			orderstatement1.WithGoodID(req.GoodID, true),
			orderstatement1.WithAppGoodID(req.AppGoodID, true),
			orderstatement1.WithOrderID(req.OrderID, true),
			orderstatement1.WithOrderUserID(req.OrderUserID, true),
			orderstatement1.WithDirectContributorID(req.DirectContributorID, true),
			orderstatement1.WithGoodCoinTypeID(req.GoodCoinTypeID, true),
			orderstatement1.WithUnits(req.Units, true),
			orderstatement1.WithGoodValueUSD(req.GoodValueUSD, true),
			orderstatement1.WithPaymentAmountUSD(req.PaymentAmountUSD, true),
			orderstatement1.WithCommissionAmountUSD(req.CommissionAmountUSD, true),
			orderstatement1.WithAppConfigID(req.AppConfigID, true),
			orderstatement1.WithCommissionConfigID(req.CommissionConfigID, false),
			orderstatement1.WithCommissionConfigType(req.CommissionConfigType, true),
			orderstatement1.WithPaymentStatements(req.PaymentStatements, true),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"CreateStatements",
				"In", in,
				"Err", err,
			)
			return &npool.CreateStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}

	if err := multiHandler.CreateStatements(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"In", in,
			"Err", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStatementsResponse{}, nil
}
