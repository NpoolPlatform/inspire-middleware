package history

import (
	"context"

	history1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/credit/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
)

func (s *Server) CreateUserCreditHistory(ctx context.Context, in *npool.CreateUserCreditHistoryRequest) (*npool.CreateUserCreditHistoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateUserCreditHistory",
			"In", in,
		)
		return &npool.CreateUserCreditHistoryResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := history1.NewHandler(
		ctx,
		history1.WithEntID(req.EntID, false),
		history1.WithAppID(req.AppID, true),
		history1.WithUserID(req.UserID, true),
		history1.WithTaskID(req.TaskID, true),
		history1.WithEventID(req.EventID, true),
		history1.WithCredits(req.Credits, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateUserCreditHistory(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateUserCreditHistory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCreditHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserCreditHistoryResponse{}, nil
}
