//nolint:dupl
package commission

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateUpdate(ctx context.Context, in *npool.CommissionReq) error { //nolint
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return err
	}

	if in.SettleType != nil {
		switch in.GetSettleType() {
		case mgrpb.SettleType_GoodOrderPercent:
			fallthrough //nolint
		case mgrpb.SettleType_GoodOrderValuePercent:
			if in.GoodID != nil {
				if _, err := uuid.Parse(in.GetGoodID()); err != nil {
					return err
				}
			}
		case mgrpb.SettleType_LimitedOrderPercent:
			fallthrough //nolint
		case mgrpb.SettleType_AmountThreshold:
			return fmt.Errorf("not implemented")
		default:
			return fmt.Errorf("invalid settle type")
		}

		switch in.GetSettleType() {
		case mgrpb.SettleType_GoodOrderPercent:
		case mgrpb.SettleType_GoodOrderValuePercent:
			fallthrough //nolint
		case mgrpb.SettleType_LimitedOrderPercent:
			if in.Percent != nil {
				percent, err := decimal.NewFromString(in.GetPercent())
				if err != nil {
					return err
				}
				if percent.Cmp(decimal.NewFromInt(100)) >= 0 { //nolint
					return fmt.Errorf("invalid percent")
				}
			}
		case mgrpb.SettleType_AmountThreshold:
			if _, err := decimal.NewFromString(in.GetThreshold()); err != nil {
				return err
			}
		default:
		}
	}

	return nil
}

func (s *Server) UpdateCommission(ctx context.Context, in *npool.UpdateCommissionRequest) (*npool.UpdateCommissionResponse, error) {
	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		return &npool.UpdateCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := commission1.UpdateCommission(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCommissionResponse{
		Info: info,
	}, nil
}
