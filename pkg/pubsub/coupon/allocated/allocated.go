package allocated

import (
	"context"
	"encoding/json"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	allocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
)

func Prepare(body string) (interface{}, error) {
	req := []*allocatedmwpb.CouponReq{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, wlog.WrapError(err)
	}
	return req, nil
}

func Apply(ctx context.Context, req interface{}, publisher *pubsub.Publisher) error {
	reqs := req.([]*allocatedmwpb.CouponReq)

	handler, err := allocated1.NewHandler(
		ctx,
		allocated1.WithReqs(reqs, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := handler.UpdateCoupons(ctx); err != nil {
		logger.Sugar().Errorf("update allocated coupons failed %v", err)
		return wlog.WrapError(err)
	}
	return nil
}
