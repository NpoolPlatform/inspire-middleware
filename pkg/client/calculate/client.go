package calculate

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

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

func Calculate(ctx context.Context, in *npool.CalculateRequest) ([]*statementmwpb.Statement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.Calculate(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*statementmwpb.Statement), nil
}
