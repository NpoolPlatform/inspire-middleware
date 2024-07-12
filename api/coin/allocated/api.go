package allocated

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	allocated.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	allocated.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return allocated.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
