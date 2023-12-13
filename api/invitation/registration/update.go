package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateRegistration(ctx context.Context, in *npool.UpdateRegistrationRequest) (*npool.UpdateRegistrationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateRegistration",
			"In", in,
		)
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithID(req.ID, true),
		registration1.WithAppID(req.AppID, true),
		registration1.WithInviterID(req.InviterID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateRegistration(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.UpdateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateRegistrationResponse{
		Info: info,
	}, nil
}
