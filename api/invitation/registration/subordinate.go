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

func (s *Server) GetSubordinates(
	ctx context.Context,
	in *npool.GetSubordinatesRequest,
) (
	*npool.GetSubordinatesResponse,
	error,
) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetSubordinatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := reg1.GetSubordinates(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetSubordinatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubordinatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}