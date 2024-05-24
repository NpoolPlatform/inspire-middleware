//nolint:dupl
package reward

import (
	"context"

	reward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/reward"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
)

func (s *Server) CreateUserReward(ctx context.Context, in *npool.CreateUserRewardRequest) (*npool.CreateUserRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateUserReward",
			"In", in,
		)
		return &npool.CreateUserRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithEntID(req.EntID, false),
		reward1.WithAppID(req.AppID, true),
		reward1.WithUserID(req.UserID, true),
		reward1.WithActionCredits(req.ActionCredits, true),
		reward1.WithCouponAmount(req.CouponAmount, true),
		reward1.WithCouponCashableAmount(req.CouponCashableAmount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateUserReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserRewardResponse{}, nil
}
