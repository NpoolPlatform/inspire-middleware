//nolint:dupl
package coupon

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUpdate(ctx context.Context, info *npool.CouponReq) error { //nolint
	if _, err := uuid.Parse(info.GetID()); err != nil {
		return err
	}

	switch info.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_SpecialOffer:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_Discount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
	default:
		return fmt.Errorf("invalid coupon type")
	}

	coup, err := coupon1.GetCoupon(ctx, info.GetID(), info.GetCouponType())
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
	if allocated.Cmp(decimal.NewFromInt(0)) > 0 {
		if info.Value != nil || info.Circulation != nil {
			return fmt.Errorf("permission denied")
		}
	}

	value := decimal.NewFromInt(0)
	circulation := decimal.NewFromInt(0)

	if info.Value != nil {
		value, err = decimal.NewFromString(info.GetValue())
		if err != nil {
			return err
		}
	}

	if info.Circulation != nil {
		circulation, err = decimal.NewFromString(info.GetCirculation())
		if err != nil {
			return err
		}
	}

	switch info.GetCouponType() {
	case allocatedmgrpb.CouponType_FixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_SpecialOffer:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_ThresholdDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodFixAmount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		if info.Value != nil {
			if circulation.Cmp(value) < 0 {
				return fmt.Errorf("value overflow")
			}
		}
	case allocatedmgrpb.CouponType_Discount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodDiscount:
		fallthrough //nolint
	case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		if info.Value != nil {
			if value.Cmp(decimal.NewFromInt(100)) >= 0 { //nolint
				return fmt.Errorf("value overflow")
			}
		}
		if info.Circulation != nil {
			if circulation.Cmp(decimal.NewFromInt(0)) <= 0 {
				return fmt.Errorf("invalid circulation")
			}
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
		if info.UserID != nil {
			if _, err := uuid.Parse(info.GetUserID()); err != nil {
				return err
			}
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
		if info.GoodID != nil {
			if _, err := uuid.Parse(info.GetGoodID()); err != nil {
				return err
			}
		}
	}

	if info.Threshold != nil {
		if _, err := decimal.NewFromString(info.GetThreshold()); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) UpdateCoupon(ctx context.Context, in *npool.UpdateCouponRequest) (*npool.UpdateCouponResponse, error) {
	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateCoupon", "error", err)
		return &npool.UpdateCouponResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := coupon1.UpdateCoupon(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoupon", "error", err)
		return &npool.UpdateCouponResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCouponResponse{
		Info: info,
	}, nil
}
