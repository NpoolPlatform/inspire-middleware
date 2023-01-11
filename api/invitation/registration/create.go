//nolint:dupl
package registration

import (
	"context"
	"fmt"

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

func ValidateCreate(ctx context.Context, info *mgrpb.RegistrationReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetInviterID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetInviteeID()); err != nil {
		return err
	}

	exist, err := ivcodemgrcli.ExistInvitationCodeConds(ctx, &ivcodemgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GetAppID(),
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

func (s *Server) CreateRegistration(
	ctx context.Context,
	in *npool.CreateRegistrationRequest,
) (
	*npool.CreateRegistrationResponse,
	error,
) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := reg1.CreateRegistration(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRegistrationResponse{
		Info: info,
	}, nil
}
