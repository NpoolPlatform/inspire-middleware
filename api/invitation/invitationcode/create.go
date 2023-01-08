//nolint:dupl
package invitationcode

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreate(info *mgrpb.InvitationCodeReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}
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

	info, err := ivcode.GetInvitationCodeOnly(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetAppID(),
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetUserID(),
		},
	})
	if err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if info != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.Internal, "InvitationCode exist")
	}

	info, err = ivcode.CreateInvitationCode(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateInvitationCodeResponse{
		Info: info,
	}, nil
}
