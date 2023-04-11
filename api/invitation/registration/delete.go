package registration

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteRegistration(
	ctx context.Context,
	in *npool.DeleteRegistrationRequest,
) (
	*npool.DeleteRegistrationResponse,
	error,
) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgrcli.DeleteRegistration(ctx, in.GetInfo().GetID())
	if err != nil {
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRegistrationResponse{
		Info: info,
	}, nil
}
