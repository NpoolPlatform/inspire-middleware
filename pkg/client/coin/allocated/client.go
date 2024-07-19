package allocated

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
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

func CreateCoinAllocated(ctx context.Context, req *npool.CoinAllocatedReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateCoinAllocated(_ctx, &npool.CreateCoinAllocatedRequest{
			Info: req,
		})
	})
	return err
}

func GetCoinAllocated(ctx context.Context, id string) (*npool.CoinAllocated, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinAllocated(ctx, &npool.GetCoinAllocatedRequest{
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
	return info.(*npool.CoinAllocated), nil
}

func GetCoinAllocateds(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.CoinAllocated, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinAllocateds(ctx, &npool.GetCoinAllocatedsRequest{
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
	return _infos.([]*npool.CoinAllocated), total, nil
}

func ExistCoinAllocatedConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistCoinAllocatedConds(ctx, &npool.ExistCoinAllocatedCondsRequest{
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

func GetCoinAllocatedOnly(ctx context.Context, conds *npool.Conds) (*npool.CoinAllocated, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCoinAllocateds(ctx, &npool.GetCoinAllocatedsRequest{
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
	if len(infos.([]*npool.CoinAllocated)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.CoinAllocated)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.CoinAllocated)[0], nil
}

func DeleteCoinAllocated(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteCoinAllocated(_ctx, &npool.DeleteCoinAllocatedRequest{
			Info: &npool.CoinAllocatedReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
