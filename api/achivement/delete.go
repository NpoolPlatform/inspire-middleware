package achivement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAchivement(ctx context.Context, in *npool.DeleteAchivementRequest) (*npool.DeleteAchivementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAchivement",
			"In", in,
		)
		return &npool.DeleteAchivementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := achivement1.NewHandler(
		ctx,
		achivement1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchivement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchivementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteAchivement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAchivement",
			"In", in,
			"Err", err,
		)
		return &npool.DeleteAchivementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAchivementResponse{
		Info: info,
	}, nil
}
