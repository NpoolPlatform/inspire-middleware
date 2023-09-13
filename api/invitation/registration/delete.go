package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteRegistration(ctx context.Context, in *npool.DeleteRegistrationRequest) (*npool.DeleteRegistrationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteRegistration",
			"In", in,
		)
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteRegistration(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteRegistrationResponse{
		Info: info,
	}, nil
}
