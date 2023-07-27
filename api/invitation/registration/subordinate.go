package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSubordinates(ctx context.Context, in *npool.GetSubordinatesRequest) (*npool.GetSubordinatesResponse, error) {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithConds(in.GetConds()),
		registration1.WithOffset(in.GetOffset()),
		registration1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubordinates",
			"In", in,
			"Err", err,
		)
		return &npool.GetSubordinatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSubordinates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubordinates",
			"In", in,
			"Err", err,
		)
		return &npool.GetSubordinatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubordinatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
