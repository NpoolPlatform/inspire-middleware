package allocated

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/credit/allocated"
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

func CreateCreditAllocated(ctx context.Context, req *npool.CreditAllocatedReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateCreditAllocated(_ctx, &npool.CreateCreditAllocatedRequest{
			Info: req,
		})
	})
	return err
}

func GetCreditAllocated(ctx context.Context, id string) (*npool.CreditAllocated, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCreditAllocated(ctx, &npool.GetCreditAllocatedRequest{
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
	return info.(*npool.CreditAllocated), nil
}

func GetCreditAllocateds(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.CreditAllocated, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCreditAllocateds(ctx, &npool.GetCreditAllocatedsRequest{
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
	return _infos.([]*npool.CreditAllocated), total, nil
}

func ExistCreditAllocatedConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistCreditAllocatedConds(ctx, &npool.ExistCreditAllocatedCondsRequest{
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

func GetCreditAllocatedOnly(ctx context.Context, conds *npool.Conds) (*npool.CreditAllocated, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetCreditAllocateds(ctx, &npool.GetCreditAllocatedsRequest{
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
	if len(infos.([]*npool.CreditAllocated)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.CreditAllocated)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.CreditAllocated)[0], nil
}

func DeleteCreditAllocated(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteCreditAllocated(_ctx, &npool.DeleteCreditAllocatedRequest{
			Info: &npool.CreditAllocatedReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
