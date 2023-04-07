package invitationcode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteInvitationCode(
	ctx context.Context,
	in *npool.DeleteInvitationCodeRequest,
) (
	*npool.DeleteInvitationCodeResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.DeleteInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.DeleteInvitationCode(ctx, in.GetInfo().GetID())
	if err != nil {
		return &npool.DeleteInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.DeleteInvitationCodeResponse{
		Info: info,
	}, nil
}
