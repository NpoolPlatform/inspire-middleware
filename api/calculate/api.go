package calculate

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	calculate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	calculate.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return calculate.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
