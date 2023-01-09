package invitationcode

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	ivcode "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/invitationcode"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateConds(conds *mgrpb.Conds) error {
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}
	if conds.UserID != nil {
		if _, err := uuid.Parse(conds.GetUserID().GetValue()); err != nil {
			return err
		}
	}
	if conds.InvitationCode != nil && conds.GetInvitationCode().GetValue() == "" {
		return fmt.Errorf("invalid invitation code")
	}
	return nil
}

func (s *Server) GetInvitationCode(
	ctx context.Context,
	in *npool.GetInvitationCodeRequest,
) (
	*npool.GetInvitationCodeResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.GetInvitationCode(ctx, in.GetID())
	if err != nil {
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetInvitationCodeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetInvitationCodes(
	ctx context.Context,
	in *npool.GetInvitationCodesRequest,
) (
	*npool.GetInvitationCodesResponse,
	error,
) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := ivcode.GetInvitationCodes(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetInvitationCodesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetInvitationCodeOnly(
	ctx context.Context,
	in *npool.GetInvitationCodeOnlyRequest,
) (
	*npool.GetInvitationCodeOnlyResponse,
	error,
) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetInvitationCodeOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := ivcode.GetInvitationCodeOnly(ctx, in.GetConds())
	if err != nil {
		return &npool.GetInvitationCodeOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetInvitationCodeOnlyResponse{
		Info: info,
	}, nil
}
