package orderstatement

import (
	"context"
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

func CreateStatement(ctx context.Context, req *npool.StatementReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateStatement(ctx, &npool.CreateStatementRequest{
			Info: req,
		})
	})
	return err
}

func CreateStatements(ctx context.Context, reqs []*npool.StatementReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateStatements(ctx, &npool.CreateStatementsRequest{
			Infos: reqs,
		})
	})
	return err
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

func ExistStatementConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		in, err := cli.ExistStatementConds(ctx, &npool.ExistStatementCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return in.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func DeleteStatement(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteStatement(ctx, &npool.DeleteStatementRequest{
			Info: &npool.StatementReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func DeleteStatements(ctx context.Context, in []*npool.StatementReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteStatements(ctx, &npool.DeleteStatementsRequest{
			Infos: in,
		})
	})
	return err
}
