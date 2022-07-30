//nolint:,dupl
package invitation

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

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

func GetInvitees(ctx context.Context, appID string, inviters []string, offset, limit int32) ([]*npool.Invitation, uint32, error) {
	total := uint32(0)

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetInvitees(ctx, &npool.GetInviteesRequest{
			AppID:   appID,
			UserIDs: inviters,
			Offset:  offset,
			Limit:   limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.GetTotal()

		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Invitation), total, nil
}
