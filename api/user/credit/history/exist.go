package history

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	history1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/credit/history"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
)

func (s *Server) ExistUserCreditHistoryConds(ctx context.Context, in *npool.ExistUserCreditHistoryCondsRequest) (*npool.ExistUserCreditHistoryCondsResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserCreditHistoryConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCreditHistoryCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistUserCreditHistoryConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserCreditHistoryConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCreditHistoryCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistUserCreditHistoryCondsResponse{
		Info: exist,
	}, nil
}
