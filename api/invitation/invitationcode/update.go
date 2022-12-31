package invitationcode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateInvitationCode(
	ctx context.Context,
	in *npool.UpdateInvitationCodeRequest,
) (
	*npool.UpdateInvitationCodeResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.UpdateInvitationCode(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.UpdateInvitationCodeResponse{
		Info: info,
	}, nil
}
