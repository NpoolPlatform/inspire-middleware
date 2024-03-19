package user

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return user.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
