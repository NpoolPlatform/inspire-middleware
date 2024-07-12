package coin

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"
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

func CreateEventCoin(ctx context.Context, req *npool.EventCoinReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateEventCoin(_ctx, &npool.CreateEventCoinRequest{
			Info: req,
		})
	})
	return err
}

func GetEventCoin(ctx context.Context, id string) (*npool.EventCoin, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoin(ctx, &npool.GetEventCoinRequest{
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
	return info.(*npool.EventCoin), nil
}

func GetEventCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.EventCoin, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoins(ctx, &npool.GetEventCoinsRequest{
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
	return _infos.([]*npool.EventCoin), total, nil
}

func ExistEventCoinConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistEventCoinConds(ctx, &npool.ExistEventCoinCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func GetEventCoinOnly(ctx context.Context, conds *npool.Conds) (*npool.EventCoin, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetEventCoins(ctx, &npool.GetEventCoinsRequest{
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
	if len(infos.([]*npool.EventCoin)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.EventCoin)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.EventCoin)[0], nil
}

func DeleteEventCoin(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteEventCoin(_ctx, &npool.DeleteEventCoinRequest{
			Info: &npool.EventCoinReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateEventCoin(ctx context.Context, req *npool.EventCoinReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateEventCoin(_ctx, &npool.UpdateEventCoinRequest{
			Info: req,
		})
	})
	return err
}
