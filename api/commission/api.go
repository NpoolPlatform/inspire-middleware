package commission

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	commission.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	commission.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return commission.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
