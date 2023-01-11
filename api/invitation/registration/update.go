//nolint:dupl
package registration

import (
	"context"
	"fmt"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	ivcodemgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	ivcodemgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

	reg1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func ValidateUpdate(ctx context.Context, info *mgrpb.RegistrationReq) error {
	if _, err := uuid.Parse(info.GetID()); err != nil {
		return err
	}
	if info.InviterID != nil {
		if _, err := uuid.Parse(info.GetInviterID()); err != nil {
			return err
		}
	}

	info1, err := mgrcli.GetRegistration(ctx, info.GetID())
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid registration")
	}

	exist, err := ivcodemgrcli.ExistInvitationCodeConds(ctx, &ivcodemgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: info1.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GetInviterID(),
		},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("permission denied")
	}

	return nil
}

func (s *Server) UpdateRegistration(
	ctx context.Context,
	in *npool.UpdateRegistrationRequest,
) (
	*npool.UpdateRegistrationResponse,
	error,
) {
	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := reg1.UpdateRegistration(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRegistrationResponse{
		Info: info,
	}, nil
}
