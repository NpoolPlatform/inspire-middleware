package invitationcode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateInvitationCode(ctx context.Context, in *npool.CreateInvitationCodeRequest) (*npool.CreateInvitationCodeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateInvitationCode",
			"In", in,
		)
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithEntID(req.EntID, false),
		invitationcode1.WithAppID(req.AppID, true),
		invitationcode1.WithUserID(req.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateInvitationCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateInvitationCodeResponse{
		Info: info,
	}, nil
}
