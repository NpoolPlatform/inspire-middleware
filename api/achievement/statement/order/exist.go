package orderstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orderstatement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistStatementConds(ctx context.Context, in *npool.ExistStatementCondsRequest) (*npool.ExistStatementCondsResponse, error) {
	handler, err := orderstatement1.NewHandler(
		ctx,
		orderstatement1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistStatementConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistStatementCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistStatementConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistStatementConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistStatementCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStatementCondsResponse{
		Info: exist,
	}, nil
}
