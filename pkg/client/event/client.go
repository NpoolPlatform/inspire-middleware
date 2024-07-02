package event

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateEvent(ctx context.Context, req *npool.EventReq) (*npool.Event, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateEvent(ctx, &npool.CreateEventRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Event), nil
}

func UpdateEvent(ctx context.Context, req *npool.EventReq) (*npool.Event, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.UpdateEvent(ctx, &npool.UpdateEventRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Event), nil
}

func GetEvent(ctx context.Context, id string) (*npool.Event, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEvent(ctx, &npool.GetEventRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Event), nil
}

func GetEvents(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Event, uint32, error) {
	total := uint32(0)

	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEvents(ctx, &npool.GetEventsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Event), total, nil
}

func GetEventOnly(ctx context.Context, conds *npool.Conds) (*npool.Event, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEvents(ctx, &npool.GetEventsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}

		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Event)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Event)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Event)[0], nil
}

func DeleteEvent(ctx context.Context, id uint32) (*npool.Event, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteEvent(ctx, &npool.DeleteEventRequest{
			Info: &npool.EventReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Event), nil
}

func RewardEvent(ctx context.Context, req *npool.RewardEventRequest) ([]*npool.Credit, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.RewardEvent(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Credit), nil
}
