package coupon

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateConds(conds *npool.Conds) error {
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
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetCoupons(ctx context.Context, in *npool.GetCouponsRequest) (*npool.GetCouponsResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := coupon1.GetCoupons(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		return &npool.GetCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCoupon(ctx context.Context, in *npool.GetCouponRequest) (*npool.GetCouponResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.GetCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
	case allocatedmgrpb.CouponType_Discount:
	case allocatedmgrpb.CouponType_SpecialOffer:
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("unknown coupon type")
	}

	info, err := coupon1.GetCoupon(ctx, in.GetID(), in.GetCouponType())
	if err != nil {
		return &npool.GetCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponResponse{
		Info: info,
	}, nil
}
