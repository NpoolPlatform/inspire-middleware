package scope

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

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

func CreateAppGoodScope(ctx context.Context, in *npool.ScopeReq) (*npool.Scope, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CreateAppGoodScope(ctx, &npool.CreateAppGoodScopeRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Scope), nil
}

func GetAppGoodScopes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Scope, uint32, error) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppGoodScopes(ctx, &npool.GetAppGoodScopesRequest{
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
	return infos.([]*npool.Scope), total, nil
}

func GetAppGoodScope(ctx context.Context, id string) (*npool.Scope, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppGoodScope(ctx, &npool.GetAppGoodScopeRequest{
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
	return info.(*npool.Scope), nil
}

func DeleteAppGoodScope(ctx context.Context, id uint32) (*npool.Scope, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.DeleteAppGoodScope(ctx, &npool.DeleteAppGoodScopeRequest{
			Info: &npool.ScopeReq{
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
	return info.(*npool.Scope), nil
}

func ExistAppGoodScopeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistAppGoodScopeConds(ctx, &npool.ExistAppGoodScopeCondsRequest{
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

func VerifyCouponScopes(ctx context.Context, in []*npool.ScopeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		_, err := cli.VerifyCouponScopes(ctx, &npool.VerifyCouponScopesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
