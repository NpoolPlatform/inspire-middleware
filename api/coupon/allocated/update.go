//nolint:dupl
package allocated

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUpdate(ctx context.Context, info *npool.CouponReq) error { //nolint
	if _, err := uuid.Parse(info.GetID()); err != nil {
		return err
	}

	if _, err := uuid.Parse(info.GetUsedByOrderID()); err != nil {
		return err
	}

	return nil
}

func (s *Server) UpdateCoupon(ctx context.Context, in *npool.UpdateCouponRequest) (*npool.UpdateCouponResponse, error) {
	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		return &npool.UpdateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := allocated1.UpdateCoupon(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCouponResponse{
		Info: info,
	}, nil
}
