//nolint:dupl
package registration

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	reg1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateConds(conds *mgrpb.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}
	if conds.InviterID != nil {
		if _, err := uuid.Parse(conds.GetInviterID().GetValue()); err != nil {
			return err
		}
	}
	if conds.InviteeID != nil {
		if _, err := uuid.Parse(conds.GetInviteeID().GetValue()); err != nil {
			return err
		}
	}
	for _, id := range conds.GetInviterIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}
	for _, id := range conds.GetInviteeIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetRegistrations(
	ctx context.Context,
	in *npool.GetRegistrationsRequest,
) (
	*npool.GetRegistrationsResponse,
	error,
) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetRegistrationsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := reg1.GetRegistrations(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetRegistrationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetRegistrationOnly(
	ctx context.Context,
	in *npool.GetRegistrationOnlyRequest,
) (
	*npool.GetRegistrationOnlyResponse,
	error,
) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetRegistrationOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := reg1.GetRegistrationOnly(ctx, in.GetConds())
	if err != nil {
		return &npool.GetRegistrationOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRegistration(
	ctx context.Context,
	in *npool.GetRegistrationRequest,
) (
	*npool.GetRegistrationResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.GetRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := reg1.GetRegistration(ctx, in.GetID())
	if err != nil {
		return &npool.GetRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationResponse{
		Info: info,
	}, nil
}
