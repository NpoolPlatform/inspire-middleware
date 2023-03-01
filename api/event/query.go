package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	mgrapi "github.com/NpoolPlatform/inspire-manager/api/event"
	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEventOnly(ctx context.Context, in *npool.GetEventOnlyRequest) (*npool.GetEventOnlyResponse, error) {
	if err := mgrapi.ValidateConds(in.GetConds()); err != nil {
		logger.Sugar().Errorw("GetEventOnly", "Conds", in.GetConds(), "Error", err)
		return &npool.GetEventOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgrcli.GetEventOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetEventOnly", "Error", err)
		return &npool.GetEventOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEventOnlyResponse{
		Info: info,
	}, nil
}
