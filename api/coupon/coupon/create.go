package coupon

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreate(info *npool.CouponReq) error { //nolint
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}
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

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}

	value, err := decimal.NewFromString(info.GetValue())
	if err != nil {
		return err
	}

	switch info.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		circulation, err := decimal.NewFromString(info.GetCirculation())
		if err != nil {
			return err
		}
		if circulation.Cmp(value) < 0 {
			return fmt.Errorf("value overflow")
		}
	case allocatedmgrpb.CouponType_Discount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		circulation, err := decimal.NewFromString(info.GetCirculation())
		if err != nil {
			return err
		}
		if value.Cmp(decimal.NewFromInt(100)) >= 0 { //nolint
			return fmt.Errorf("value overflow")
		}
		if circulation.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid circulation")
		}
	default:
	}

	if info.GetDurationDays() <= 0 {
		return fmt.Errorf("invalid durationdays")
	}

	if info.GetMessage() == "" {
		return fmt.Errorf("invalid message")
	}
	if info.GetName() == "" {
		return fmt.Errorf("invalid name")
	}

	if info.GetCouponType() == allocatedmgrpb.CouponType_SpecialOffer {
		if _, err := uuid.Parse(info.GetUserID()); err != nil {
			return err
		}
	}

	switch info.GetCouponType() {
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		if _, err := uuid.Parse(info.GetGoodID()); err != nil {
			return err
		}
	}

	if info.Threshold != nil {
		if _, err := decimal.NewFromString(info.GetThreshold()); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) CreateCoupon(ctx context.Context, in *npool.CreateCouponRequest) (*npool.CreateCouponResponse, error) {
	if err := ValidateCreate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateCoupon", "error", err)
		return &npool.CreateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := coupon1.CreateCoupon(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoupon", "error", err)
		return &npool.CreateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCouponResponse{
		Info: info,
	}, nil
}
