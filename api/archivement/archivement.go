package archivement

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/gw/v1/archivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoinArchivements(ctx context.Context, in *npool.GetCoinArchivementsRequest) (*npool.GetCoinArchivementsResponse, error) {
	return &npool.GetCoinArchivementsResponse{}, status.Error(codes.Internal, "NOT IMPLEMENTED")
}
