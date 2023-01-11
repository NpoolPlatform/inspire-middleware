package archivement

import (
	"context"
	"errors"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	detailmgrapi "github.com/NpoolPlatform/inspire-manager/api/archivement/detail"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"

	errno "github.com/NpoolPlatform/go-service-framework/pkg/errno"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if err := detailmgrapi.Validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := archivement1.BookKeeping(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		if errors.Is(err, errno.ErrAlreadyExists) {
			return &npool.BookKeepingResponse{}, status.Error(codes.AlreadyExists, err.Error())
		}
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}
