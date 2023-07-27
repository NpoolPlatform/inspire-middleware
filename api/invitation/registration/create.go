package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRegistration(ctx context.Context, in *npool.CreateRegistrationRequest) (*npool.CreateRegistrationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateRegistration",
			"In", in,
		)
		return &npool.CreateRegistrationResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithID(req.ID),
		registration1.WithAppID(req.AppID),
		registration1.WithInviterID(req.InviterID),
		registration1.WithInviteeID(req.InviteeID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.CreateRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateRegistration(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.CreateRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRegistrationResponse{
		Info: info,
	}, nil
}
