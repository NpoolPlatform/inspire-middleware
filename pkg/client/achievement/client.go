package achievement

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

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

func ExpropriateAchievement(ctx context.Context, orderID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.ExpropriateAchievement(ctx, &npool.ExpropriateAchievementRequest{
			OrderID: orderID,
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

func GetAchievements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Achievement, uint32, error) {
	total := uint32(0)
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAchievements(ctx, &npool.GetAchievementsRequest{
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
	return infos.([]*npool.Achievement), total, nil
}
