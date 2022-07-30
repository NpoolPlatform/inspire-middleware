package commission

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/gw/v1/commission"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPercents(ctx context.Context, in *npool.GetPercentsRequest) (*npool.GetPercentsResponse, error) {
	return &npool.GetPercentsResponse{}, status.Error(codes.Internal, "NOT IMPLEMENTED")
}
