package event

import (
	"context"
	"encoding/json"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	event1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func Prepare(body string) (interface{}, error) {
	req := eventmwpb.RewardEventRequest{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Apply(ctx context.Context, req interface{}, publisher *pubsub.Publisher) error {
	_req := req.(*eventmwpb.RewardEventRequest)

	handler, err := event1.NewHandler(
		ctx,
		event1.WithAppID(&_req.AppID),
		event1.WithUserID(&_req.UserID),
		event1.WithEventType(&_req.EventType),
		event1.WithGoodID(_req.GoodID),
		event1.WithConsecutive(&_req.Consecutive),
		event1.WithAmount(&_req.Amount),
	)
	if err != nil {
		return err
	}
	credits, err := handler.RewardEvent(ctx)
	if err != nil {
		return err
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
		return err
	}
	if err := publisher.Publish(); err != nil {
		return err
	}

	return nil
}
