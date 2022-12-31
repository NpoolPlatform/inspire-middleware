//nolint:dupl
package invitationcode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetInvitationCode(
	ctx context.Context,
	in *npool.GetInvitationCodeRequest,
) (
	*npool.GetInvitationCodeResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.GetInvitationCode(ctx, in.GetAppID(), in.GetUserID())
	if err != nil {
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetInvitationCodeResponse{
		Info: info,
	}, nil
}
