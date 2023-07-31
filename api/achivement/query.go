package achivement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	achivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAchivements(ctx context.Context, in *npool.GetAchivementsRequest) (*npool.GetAchivementsResponse, error) {
	handler, err := achivement1.NewHandler(
		ctx,
		achivement1.WithConds(in.GetConds()),
		achivement1.WithOffset(in.GetOffset()),
		achivement1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchivements",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchivementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAchivements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAchivements",
			"In", in,
			"Err", err,
		)
		return &npool.GetAchivementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAchivementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
