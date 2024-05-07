package history

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return history.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
