package commission

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	comm1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneCommissions(ctx context.Context, in *npool.CloneCommissionsRequest) (*npool.CloneCommissionsResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.CloneCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetFromGoodID()); err != nil {
		return &npool.CloneCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetToGoodID()); err != nil {
		return &npool.CloneCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err := comm1.CloneCommissions(ctx, in.GetAppID(), in.GetFromGoodID(), in.GetToGoodID(), in.GetValue())
	if err != nil {
		return &npool.CloneCommissionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CloneCommissionsResponse{}, nil
}
