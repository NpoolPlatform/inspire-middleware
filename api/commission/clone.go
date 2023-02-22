package commission

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

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

	switch in.GetSettleType() {
	case mgrpb.SettleType_GoodOrderPercent:
	case mgrpb.SettleType_GoodOrderValuePercent:
	default:
		return &npool.CloneCommissionsResponse{}, status.Error(codes.InvalidArgument, "SettleType is invalid")
	}
	err := comm1.CloneCommissions(
		ctx,
		in.GetAppID(),
		in.GetFromGoodID(),
		in.GetToGoodID(),
		in.GetSettleType(),
	)
	if err != nil {
		return &npool.CloneCommissionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CloneCommissionsResponse{}, nil
}
