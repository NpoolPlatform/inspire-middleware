package invitationcode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetInvitationCode(ctx context.Context, in *npool.GetInvitationCodeRequest) (*npool.GetInvitationCodeResponse, error) {
	handler, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetInvitationCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetInvitationCode",
			"In", in,
			"Err", err,
		)
		return &npool.GetInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitationCodeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetInvitationCodes(ctx context.Context, in *npool.GetInvitationCodesRequest) (*npool.GetInvitationCodesResponse, error) {
	handler, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithConds(in.GetConds()),
		invitationcode1.WithOffset(in.GetOffset()),
		invitationcode1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetInvitationCodes",
			"In", in,
			"Err", err,
		)
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetInvitationCodes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetInvitationCodes",
			"In", in,
			"Err", err,
		)
		return &npool.GetInvitationCodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetInvitationCodesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
