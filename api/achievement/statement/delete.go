package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

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
		return &npool.DeleteStatementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteStatementResponse{
		Info: info,
	}, nil
}
