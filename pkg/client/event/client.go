package event

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetEventOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.Event, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetEventOnly(ctx, &npool.GetEventOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*mgrpb.Event), nil
}
