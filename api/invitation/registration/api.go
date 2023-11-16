package registration

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	registration.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	registration.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return registration.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
