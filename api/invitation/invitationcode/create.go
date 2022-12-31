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

func (s *Server) CreateInvitationCode(
	ctx context.Context,
	in *npool.CreateInvitationCodeRequest,
) (
	*npool.CreateInvitationCodeResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.CreateInvitationCode(ctx, in.GetAppID(), in.GetUserID())
	if err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateInvitationCodeResponse{
		Info: info,
	}, nil
}
