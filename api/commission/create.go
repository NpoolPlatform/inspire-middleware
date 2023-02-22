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

//nolint:gocyclo
func ValidateCreate(ctx context.Context, in *npool.CommissionReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		return err
	}

	switch in.GetSettleType() {
	case mgrpb.SettleType_GoodOrderValuePercent:
		fallthrough //nolint
	case mgrpb.SettleType_GoodOrderPercent:
		if _, err := uuid.Parse(in.GetGoodID()); err != nil {
			return err
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
		fallthrough //nolint
	case mgrpb.SettleType_GoodOrderValuePercent:
		fallthrough //nolint
	case mgrpb.SettleType_LimitedOrderPercent:
		percent, err := decimal.NewFromString(in.GetPercent())
		if err != nil {
			return err
		}
		if percent.Cmp(decimal.NewFromInt(100)) >= 0 { //nolint
			return fmt.Errorf("invalid percent")
		}
	case mgrpb.SettleType_AmountThreshold:
		if _, err := decimal.NewFromString(in.GetThreshold()); err != nil {
			return err
		}
	default:
	}

	return nil
}

func (s *Server) CreateCommission(ctx context.Context, in *npool.CreateCommissionRequest) (*npool.CreateCommissionResponse, error) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := commission1.CreateCommission(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCommissionResponse{
		Info: info,
	}, nil
}
