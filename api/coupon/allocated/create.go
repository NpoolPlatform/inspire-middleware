package allocated

import (
	"context"
	"fmt"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreate(ctx context.Context, info *npool.CouponReq) error { //nolint
	switch info.GetCouponType() {
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
		return fmt.Errorf("not implemented")
	default:
		return fmt.Errorf("unknown coupon type")
	}

	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetCouponID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return err
	}

	coup, err := coupon1.GetCoupon(ctx, info.GetCouponID(), info.GetCouponType())
	if err != nil {
		return err
	}
	if coup == nil {
		return fmt.Errorf("invalid coupon")
	}

	allocated, err := decimal.NewFromString(coup.Allocated)
	if err != nil {
		return err
	}

	switch info.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		value, err := decimal.NewFromString(coup.Value)
		if err != nil {
			return err
		}
		circulation, err := decimal.NewFromString(coup.Circulation)
		if err != nil {
			return err
		}
		allocated = value.Mul(allocated)
		if allocated.Add(value).Cmp(circulation) > 0 {
			return fmt.Errorf("overflow")
		}
	case allocatedmgrpb.CouponType_Discount:
	case allocatedmgrpb.CouponType_ThresholdDiscount:
	case allocatedmgrpb.CouponType_GoodDiscount:
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		circulation, err := decimal.NewFromString(coup.Circulation)
		if err != nil {
			return err
		}
		if allocated.Cmp(circulation) >= 0 {
			return fmt.Errorf("overflow")
		}
	default:
	}

	return nil
}

func (s *Server) CreateCoupon(ctx context.Context, in *npool.CreateCouponRequest) (*npool.CreateCouponResponse, error) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := allocated1.CreateCoupon(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCouponResponse{
		Info: info,
	}, nil
}
