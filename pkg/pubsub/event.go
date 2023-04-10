package pubsub

import (
	"context"
	"encoding/json"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	event "github.com/NpoolPlatform/inspire-middleware/pkg/event"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func prepareRewardEvent(body string) (interface{}, error) {
	req := eventmwpb.RewardEventRequest{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func handleRewardEvent(ctx context.Context, req interface{}) error {
	_req := req.(*eventmwpb.RewardEventRequest)

	handler, err := event.NewHandler(
		ctx,
		event.WithAppID(_req.GetAppID()),
		event.WithUserID(_req.GetUserID()),
		event.WithEventType(_req.GetEventType()),
		event.WithGoodID(_req.GoodID),
		event.WithConsecutive(_req.GetConsecutive()),
		event.WithAmount(_req.GetAmount()),
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
