package allocated

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoupon(ctx context.Context, in *npool.CreateCouponRequest) (*npool.CreateCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(req.EntID, false),
		allocated1.WithAppID(req.AppID, true),
		allocated1.WithCouponID(req.CouponID, true),
		allocated1.WithUserID(req.UserID, true),
		allocated1.WithExtra(req.Extra, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	err = handler.CreateCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCouponResponse{
		Info: nil,
	}, nil
}

func (s *Server) CreateDirectCoupon(ctx context.Context, in *npool.CreateDirectCouponRequest) (*npool.CreateDirectCouponResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDirectCoupon",
			"In", in,
		)
		return &npool.CreateDirectCouponResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithEntID(req.EntID, false),
		allocated1.WithAppID(req.AppID, true),
		allocated1.WithCouponID(req.CouponID, true),
		allocated1.WithUserID(req.UserID, true),
		allocated1.WithCashable(req.Cashable, true),
		allocated1.WithExtra(req.Extra, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDirectCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateDirectCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	err = handler.CreateCoupon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDirectCoupon",
			"In", in,
			"Err", err,
		)
		return &npool.CreateDirectCouponResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDirectCouponResponse{
		Info: nil,
	}, nil
}
