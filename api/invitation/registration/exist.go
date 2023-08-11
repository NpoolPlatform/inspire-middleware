package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistRegistrationConds(ctx context.Context, in *npool.ExistRegistrationCondsRequest) (*npool.ExistRegistrationCondsResponse, error) {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRegistrationConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistRegistrationCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistRegistrationConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRegistrationConds",
			"In", in,
			"Err", err,
		)
		return &npool.ExistRegistrationCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRegistrationCondsResponse{
		Info: exist,
	}, nil
}
