package commission

import (
	"context"

	commission "github.com/NpoolPlatform/message/npool/inspire/gw/v1/commission"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	commission.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	commission.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := commission.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
