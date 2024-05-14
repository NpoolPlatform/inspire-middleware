package orderstatement

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
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

func CreateStatement(ctx context.Context, req *npool.StatementReq) (*npool.Statement, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateStatement(ctx, &npool.CreateStatementRequest{
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
	return info.(*npool.Statement), nil
}

func CreateStatements(ctx context.Context, reqs []*npool.StatementReq) ([]*npool.Statement, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateStatements(ctx, &npool.CreateStatementsRequest{
			Infos: reqs,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Statement), nil
}

func GetStatements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Statement, uint32, error) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		in, err := cli.GetStatements(ctx, &npool.GetStatementsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = in.GetTotal()
		return in.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Statement), total, nil
}

func DeleteStatement(ctx context.Context, id uint32) (*npool.Statement, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteStatement(ctx, &npool.DeleteStatementRequest{
			Info: &npool.StatementReq{
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
	return info.(*npool.Statement), nil
}

func DeleteStatements(ctx context.Context, in []*npool.StatementReq) ([]*npool.Statement, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteStatements(ctx, &npool.DeleteStatementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete statements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete statements: %v", err)
	}
	return infos.([]*npool.Statement), nil
}
