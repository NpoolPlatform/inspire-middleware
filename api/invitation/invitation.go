package invitation

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

	converter "github.com/NpoolPlatform/inspire-middleware/pkg/converter/invitation"
	invitation1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetInvitees(ctx context.Context, in *npool.GetInviteesRequest) (*npool.GetInviteesResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetInvitees", "AppID", in.GetAppID(), "error", err)
		return &npool.GetInviteesResponse{}, status.Error(codes.Internal, "AppID is invalid")
	}

	if len(in.GetUserIDs()) == 0 {
		logger.Sugar().Errorw("GetInvitees", "error", "UserIDs is empty")
		return &npool.GetInviteesResponse{}, status.Error(codes.Internal, "UserIDs is invalid")
	}

	for _, user := range in.GetUserIDs() {
		if _, err := uuid.Parse(user); err != nil {
			logger.Sugar().Errorw("GetInvitees", "UserID", user, "error", err)
			return &npool.GetInviteesResponse{}, status.Error(codes.Internal, "UserID is invalid")
		}
	}

	limit := int32(constant.DefaultLimitRows)
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := invitation1.GetInvitees(ctx, in.GetAppID(), in.GetUserIDs(), in.GetOffset(), limit)

	if err != nil {
		logger.Sugar().Errorw("GetInvitees",
			"AppID", in.GetAppID(), "UserID", in.GetUserIDs(),
			"Offset", in.GetOffset(), "Limit", in.GetLimit(),
			"error", err)
		return &npool.GetInviteesResponse{}, status.Error(codes.Internal, "fail get invitees")
	}

	return &npool.GetInviteesResponse{
		Infos: converter.Ent2GrpcMany(infos),
		Total: total,
	}, nil
}

func (s *Server) GetInviters(ctx context.Context, in *npool.GetInvitersRequest) (*npool.GetInvitersResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetInviters", "AppID", in.GetAppID(), "error", err)
		return &npool.GetInvitersResponse{}, status.Error(codes.Internal, "AppID is invalid")
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetInviters", "UserID", in.GetUserID(), "error", err)
		return &npool.GetInvitersResponse{}, status.Error(codes.Internal, "UserID is invalid")
	}

	infos, total, err := invitation1.GetInviters(ctx, in.GetAppID(), in.GetUserID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetInviters",
			"AppID", in.GetAppID(), "UserID", in.GetUserID(),
			"Offset", in.GetOffset(), "Limit", in.GetLimit(),
			"error", err)
		return &npool.GetInvitersResponse{}, status.Error(codes.Internal, "fail get invitees")
	}

	return &npool.GetInvitersResponse{
		Infos: converter.Ent2GrpcMany(infos),
		Total: total,
	}, nil
}
