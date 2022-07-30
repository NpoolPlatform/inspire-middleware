//nolint:dupl
package archivement

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/gw/v1/archivement"

	archivement1 "github.com/NpoolPlatform/inspire-gateway/pkg/archivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetCoinArchivements(ctx context.Context, in *npool.GetCoinArchivementsRequest) (*npool.GetCoinArchivementsResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetCoinArchivements", "AppID", in.GetAppID(), "error", err)
		return &npool.GetCoinArchivementsResponse{}, status.Error(codes.Internal, "AppID is invalid")
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetCoinArchivements", "UserID", in.GetUserID(), "error", err)
		return &npool.GetCoinArchivementsResponse{}, status.Error(codes.Internal, "UserID is invalid")
	}

	start := in.StartAt
	end := in.EndAt
	if end == 0 {
		end = uint32(time.Now().Unix())
	}

	totalAmount, selfAmount, infos, total, err := archivement1.GetCoinArchivements(ctx,
		in.GetAppID(), in.GetUserID(), in.GetWithSub(),
		start, end, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetCoinArchivements",
			"AppID", in.GetAppID(), "UserID", in.GetUserID(),
			"WithSub", in.GetWithSub(),
			"StartAt", start, "EndAt", end,
			"Offset", in.GetOffset(), "Limit", in.GetLimit(),
			"error", err)
		return &npool.GetCoinArchivementsResponse{}, status.Error(codes.Internal, "fail get coin archivements")
	}

	return &npool.GetCoinArchivementsResponse{
		TotalAmount:  totalAmount.String(),
		SelfAmount:   selfAmount.String(),
		Archivements: infos,
		Total:        total,
	}, nil
}
