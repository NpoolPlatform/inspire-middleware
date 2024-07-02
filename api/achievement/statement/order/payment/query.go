package orderpaymentstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orderpaymentstatement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order/payment"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetStatements(ctx context.Context, in *npool.GetStatementsRequest) (*npool.GetStatementsResponse, error) {
	handler, err := orderpaymentstatement1.NewHandler(
		ctx,
		orderpaymentstatement1.WithConds(in.GetConds()),
		orderpaymentstatement1.WithOffset(in.GetOffset()),
		orderpaymentstatement1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatements",
			"In", in,
			"Err", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatements",
			"In", in,
			"Err", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
