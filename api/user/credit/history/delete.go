package history

import (
	"context"

	history1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/credit/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
)

func (s *Server) DeleteUserCreditHistory(ctx context.Context, in *npool.DeleteUserCreditHistoryRequest) (*npool.DeleteUserCreditHistoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteUserCreditHistory",
			"In", in,
		)
		return &npool.DeleteUserCreditHistoryResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := history1.NewHandler(
		ctx,
		history1.WithID(req.ID, false),
		history1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteUserCreditHistory(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserCreditHistoryResponse{}, nil
}
