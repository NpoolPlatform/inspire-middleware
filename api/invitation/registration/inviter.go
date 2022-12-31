//nolint:dupl
package registration

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	reg1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetInviters(
	ctx context.Context,
	in *npool.GetInvitersRequest,
) (
	*npool.GetInvitersResponse,
	error,
) {
	if len(in.GetConds().GetInviterIDs().GetValue()) == 0 {
		return &npool.GetInvitersResponse{}, status.Error(codes.InvalidArgument, "InviterIDs is invalid")
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetInvitersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := reg1.GetInviters(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetInvitersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
