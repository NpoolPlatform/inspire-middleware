package orderstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orderstatement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateStatements(ctx context.Context, in *npool.UpdateStatementsRequest) (*npool.UpdateStatementsResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"UpdateStatements",
			"In", in,
		)
		return &npool.UpdateStatementsResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	multiHandler := &orderstatement1.MultiHandler{}
	for _, req := range reqs {
		handler, err := orderstatement1.NewHandler(
			ctx,
			orderstatement1.WithID(req.ID, false),
			orderstatement1.WithEntID(req.EntID, false),
			orderstatement1.WithAppID(req.AppID, true),
			orderstatement1.WithUserID(req.UserID, true),
			orderstatement1.WithGoodID(req.GoodID, true),
			orderstatement1.WithAppGoodID(req.AppGoodID, true),
			orderstatement1.WithOrderID(req.OrderID, true),
			orderstatement1.WithOrderUserID(req.OrderUserID, true),
			orderstatement1.WithGoodCoinTypeID(req.GoodCoinTypeID, true),
			orderstatement1.WithUnits(req.Units, true),
			orderstatement1.WithGoodValueUSD(req.GoodValueUSD, true),
			orderstatement1.WithPaymentAmountUSD(req.PaymentAmountUSD, true),
			orderstatement1.WithCommissionAmountUSD(req.CommissionAmountUSD, true),
			orderstatement1.WithAppConfigID(req.AppConfigID, true),
			orderstatement1.WithCommissionConfigID(req.CommissionConfigID, true),
			orderstatement1.WithCommissionConfigType(req.CommissionConfigType, true),
			orderstatement1.WithPaymentStatements(req.PaymentStatements, true),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"UpdateStatements",
				"In", in,
				"Err", err,
			)
			return &npool.UpdateStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}

	if err := multiHandler.UpdateStatements(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateStatementsResponse{}, nil
}
