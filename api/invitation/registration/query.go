package registration

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRegistration(ctx context.Context, in *npool.GetRegistrationRequest) (*npool.GetRegistrationResponse, error) {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.GetRegistrationResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetRegistration(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRegistration",
			"In", in,
			"Err", err,
		)
		return &npool.GetRegistrationResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRegistrations(ctx context.Context, in *npool.GetRegistrationsRequest) (*npool.GetRegistrationsResponse, error) {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithConds(in.GetConds()),
		registration1.WithOffset(in.GetOffset()),
		registration1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRegistrations",
			"In", in,
			"Err", err,
		)
		return &npool.GetRegistrationsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetRegistrations(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRegistrations",
			"In", in,
			"Err", err,
		)
		return &npool.GetRegistrationsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetRegistrationsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
