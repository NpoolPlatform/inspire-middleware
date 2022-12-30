package coupon

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

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
