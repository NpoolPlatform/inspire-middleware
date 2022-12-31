//nolint:dupl
package invitationcode

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreate(info *mgrpb.InvitationCodeReq) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return err
	}
	return nil
}

func (s *Server) CreateInvitationCode(
	ctx context.Context,
	in *npool.CreateInvitationCodeRequest,
) (
	*npool.CreateInvitationCodeResponse,
	error,
) {
	if err := ValidateCreate(in.GetInfo()); err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.CreateInvitationCode(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateInvitationCodeResponse{
		Info: info,
	}, nil
}
