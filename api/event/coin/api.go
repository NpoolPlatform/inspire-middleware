package coin

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	coin.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coin.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return coin.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
