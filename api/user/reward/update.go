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

func (s *Server) UpdateUserReward(ctx context.Context, in *npool.UpdateUserRewardRequest) (*npool.UpdateUserRewardResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateUserReward",
			"In", in,
		)
		return &npool.UpdateUserRewardResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := reward1.NewHandler(
		ctx,
		reward1.WithID(req.ID, false),
		reward1.WithEntID(req.EntID, false),
		reward1.WithActionCredits(req.ActionCredits, false),
		reward1.WithCouponAmount(req.CouponAmount, false),
		reward1.WithCouponCashableAmount(req.CouponCashableAmount, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateUserReward(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateUserReward",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserRewardResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateUserRewardResponse{}, nil
}
