package orderstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orderstatement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteStatement(ctx context.Context, in *npool.DeleteStatementRequest) (*npool.DeleteStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := orderstatement1.NewHandler(
		ctx,
		orderstatement1.WithID(req.ID, false),
		orderstatement1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteStatement(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteStatementResponse{}, nil
}

func (s *Server) DeleteStatements(ctx context.Context, in *npool.DeleteStatementsRequest) (*npool.DeleteStatementsResponse, error) {
	reqs := in.GetInfos()
	multiHandler := &orderstatement1.MultiHandler{}

	for _, req := range reqs {
		if req == nil {
			logger.Sugar().Errorw(
				"DeleteStatement",
				"In", in,
			)
			return &npool.DeleteStatementsResponse{}, status.Error(codes.Aborted, "invalid infos")
		}
		handler, err := orderstatement1.NewHandler(
			ctx,
			orderstatement1.WithID(req.ID, false),
			orderstatement1.WithEntID(req.EntID, false),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"DeleteStatement",
				"In", in,
				"Err", err,
			)
			return &npool.DeleteStatementsResponse{}, status.Error(codes.Aborted, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}

	if err := multiHandler.DeleteStatements(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteStatements",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteStatementsResponse{}, nil
}
