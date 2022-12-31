package allocated

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoupon(ctx context.Context, in *npool.GetCouponRequest) (*npool.GetCouponResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.GetCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := allocated1.GetCoupon(ctx, in.GetID())
	if err != nil {
		return &npool.GetCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponResponse{
		Info: info,
	}, nil
}

func (s *Server) GetManyCoupons(ctx context.Context, in *npool.GetManyCouponsRequest) (*npool.GetManyCouponsResponse, error) {
	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			return &npool.GetManyCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, err := allocated1.GetManyCoupons(ctx, in.GetIDs())
	if err != nil {
		return &npool.GetManyCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetManyCouponsResponse{
		Infos: infos,
	}, nil
}

func ValidateConds(conds *allocatedmgrpb.Conds) error { //nolint
	switch allocatedmgrpb.CouponType(conds.GetCouponType().GetValue()) {
	case allocatedmgrpb.CouponType_FixAmount:
	case allocatedmgrpb.CouponType_Discount:
	case allocatedmgrpb.CouponType_SpecialOffer:
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
	case allocatedmgrpb.CouponType_ThresholdDiscount:
	case allocatedmgrpb.CouponType_GoodFixAmount:
	case allocatedmgrpb.CouponType_GoodDiscount:
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
	default:
		return fmt.Errorf("unknown coupon type")
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}
	if conds.UserID != nil {
		if _, err := uuid.Parse(conds.GetUserID().GetValue()); err != nil {
			return err
		}
	}
	if conds.CouponID != nil {
		if _, err := uuid.Parse(conds.GetCouponID().GetValue()); err != nil {
			return err
		}
	}
	if conds.UsedByOrderID != nil {
		if _, err := uuid.Parse(conds.GetUsedByOrderID().GetValue()); err != nil {
			return err
		}
	}
	for _, id := range conds.GetIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}
	for _, id := range conds.GetUsedByOrderIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetCoupons(ctx context.Context, in *npool.GetCouponsRequest) (*npool.GetCouponsResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := allocated1.GetCoupons(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCouponOnly(ctx context.Context, in *npool.GetCouponOnlyRequest) (*npool.GetCouponOnlyResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCouponOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := allocated1.GetCouponOnly(ctx, in.GetConds())
	if err != nil {
		return &npool.GetCouponOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponOnlyResponse{
		Info: info,
	}, nil
}
