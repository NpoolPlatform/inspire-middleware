package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RewardEvent(ctx context.Context, in *npool.RewardEventRequest) (*npool.RewardEventResponse, error) { //nolint
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("RewardEvent", "AppID", in.GetAppID(), "Error", err)
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("RewardEvent", "UserID", in.GetUserID(), "Error", err)
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	switch in.GetEventType() {
	case basetypes.UsedFor_Signup:
	case basetypes.UsedFor_Signin:
	case basetypes.UsedFor_Update:
	case basetypes.UsedFor_Contact:
	case basetypes.UsedFor_SetWithdrawAddress:
	case basetypes.UsedFor_Withdraw:
	case basetypes.UsedFor_CreateInvitationCode:
	case basetypes.UsedFor_SetCommission:
	case basetypes.UsedFor_SetTransferTargetUser:
	case basetypes.UsedFor_WithdrawalRequest:
	case basetypes.UsedFor_WithdrawalCompleted:
	case basetypes.UsedFor_DepositReceived:
	case basetypes.UsedFor_KYCApproved:
	case basetypes.UsedFor_KYCRejected:
	case basetypes.UsedFor_Purchase:
		if _, err := uuid.Parse(in.GetGoodID()); err != nil {
			logger.Sugar().Errorw("RewardEvent", "GoodID", in.GetGoodID(), "Error", err)
			return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	default:
		logger.Sugar().Errorw("RewardEvent", "EventType", in.GetEventType())
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, "EventType is invalid")
	}
	if in.GetConsecutive() <= 0 {
		logger.Sugar().Errorw("RewardEvent", "Consecutive", in.GetConsecutive())
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, "Consecutive is invalid")
	}
	amount, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		logger.Sugar().Errorw("RewardEvent", "Amount", in.GetAmount(), "Error", err)
		return &npool.RewardEventResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := event1.RewardEvent(
		ctx,
		in.GetAppID(),
		in.GetUserID(),
		in.GetEventType(),
		in.GoodID,
		in.GetConsecutive(),
		amount,
	)
	if err != nil {
		logger.Sugar().Errorw("RewardEvent", "Error", err)
		return &npool.RewardEventResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.RewardEventResponse{
		Info: info.String(),
	}, nil
}
