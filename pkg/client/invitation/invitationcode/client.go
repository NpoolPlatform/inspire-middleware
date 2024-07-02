package invitationcode

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

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

func CreateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetInvitationCodes(ctx, &npool.GetInvitationCodesRequest{
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
	if len(infos.([]*npool.InvitationCode)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.InvitationCode)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.InvitationCode)[0], nil
}

func UpdateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetInvitationCode(ctx, &npool.GetInvitationCodeRequest{
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
	return info.(*npool.InvitationCode), nil
}

func GetInvitationCodes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.InvitationCode, uint32, error) {
	var total uint32

	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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

func DeleteInvitationCode(ctx context.Context, id uint32) (*npool.InvitationCode, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
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
