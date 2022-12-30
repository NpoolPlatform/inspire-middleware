package commission

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"
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
	if conds.UserID != nil {
		if _, err := uuid.Parse(conds.GetUserID().GetValue()); err != nil {
			return err
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return err
		}
	}

	switch mgrpb.SettleType(conds.GetSettleType().GetValue()) {
	case mgrpb.SettleType_GoodOrderPercent:
	case mgrpb.SettleType_LimitedOrderPercent:
		fallthrough //nolint
	case mgrpb.SettleType_AmountThreshold:
		return fmt.Errorf("not implemented")
	default:
		return fmt.Errorf("invalid settle type")
	}

	return nil
}

func (s *Server) GetCommissions(ctx context.Context, in *npool.GetCommissionsRequest) (*npool.GetCommissionsResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCommissionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := commission1.GetCommissions(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		return &npool.GetCommissionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommissionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCommissionOnly(ctx context.Context, in *npool.GetCommissionOnlyRequest) (*npool.GetCommissionOnlyResponse, error) {
	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCommissionOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := commission1.GetCommissionOnly(ctx, in.GetConds())
	if err != nil {
		return &npool.GetCommissionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCommissionOnlyResponse{
		Info: info,
	}, nil
}
