package invitationcode

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

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

func CreateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateInvitationCode(ctx, &npool.CreateInvitationCodeRequest{
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
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCodeOnly(ctx context.Context, conds *npool.Conds) (*npool.InvitationCode, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCodes(ctx, &npool.GetInvitationCodesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.InvitationCode)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.InvitationCode)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.InvitationCode)[0], nil
}

func UpdateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateInvitationCode(ctx, &npool.UpdateInvitationCodeRequest{
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
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCode(ctx context.Context, id string) (*npool.InvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCode(ctx, &npool.GetInvitationCodeRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCodes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.InvitationCode, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetInvitationCodes(ctx, &npool.GetInvitationCodesRequest{
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
	return infos.([]*npool.InvitationCode), total, nil
}

func DeleteInvitationCode(ctx context.Context, id string) (*npool.InvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteInvitationCode(ctx, &npool.DeleteInvitationCodeRequest{
			Info: &npool.InvitationCodeReq{
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
	return info.(*npool.InvitationCode), nil
}
