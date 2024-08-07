package event

import (
	"context"
	"encoding/json"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func Prepare(body string) (interface{}, error) {
	req := eventmwpb.RewardEventRequest{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, wlog.WrapError(err)
	}
	return &req, nil
}

func Apply(ctx context.Context, req interface{}, publisher *pubsub.Publisher) error {
	_req := req.(*eventmwpb.RewardEventRequest)

	handler, err := event1.NewHandler(
		ctx,
		event1.WithAppID(&_req.AppID, true),
		event1.WithUserID(&_req.UserID, true),
		event1.WithEventType(&_req.EventType, true),
		event1.WithGoodID(_req.GoodID, false),
		event1.WithAppGoodID(_req.AppGoodID, false),
		event1.WithConsecutive(&_req.Consecutive, true),
		event1.WithAmount(_req.Amount, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	credits, err := handler.RewardEvent(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(credits) == 0 {
		return nil
	}
	if err := publisher.Update(
		basetypes.MsgID_IncreaseUserActionCreditsReq.String(),
		nil,
		nil,
		nil,
		credits,
	); err != nil {
		return wlog.WrapError(err)
	}
	if err := publisher.Publish(); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
