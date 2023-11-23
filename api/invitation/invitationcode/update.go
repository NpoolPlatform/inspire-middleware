//nolint
package invitationcode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateInvitationCode(ctx context.Context, in *npool.UpdateInvitationCodeRequest) (*npool.UpdateInvitationCodeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateInvitationCode",
			"In", in,
		)
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithID(req.ID, true),
		invitationcode1.WithDisabled(req.Disabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateInvitationCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateInvitationCodeResponse{
		Info: info,
	}, nil
}
