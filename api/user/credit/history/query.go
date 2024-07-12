package history

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	history1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/credit/history"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
)

func (s *Server) GetUserCreditHistory(ctx context.Context, in *npool.GetUserCreditHistoryRequest) (*npool.GetUserCreditHistoryResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetUserCreditHistory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCreditHistoryResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserCreditHistories(ctx context.Context, in *npool.GetUserCreditHistoriesRequest) (*npool.GetUserCreditHistoriesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.GetConds()),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetUserCreditHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCreditHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
