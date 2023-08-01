package event

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateEvent(ctx context.Context, req *npool.EventReq) (*npool.Event, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetEvent(ctx, &npool.GetEventRequest{
			ID: id,
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

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetEvents(ctx, &npool.GetEventsRequest{
			Conds: conds,
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
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Event)[0], nil
}

func DeleteEvent(ctx context.Context, id string) (*npool.Event, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
