//nolint:dupl
package invitationcode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteInvitationCode(ctx context.Context, in *npool.DeleteInvitationCodeRequest) (*npool.DeleteInvitationCodeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteInvitationCode",
			"In", in,
		)
		return &npool.DeleteInvitationCodeResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithID(req.ID, false),
		invitationcode1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteInvitationCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteInvitationCodeResponse{
		Info: info,
	}, nil
}
