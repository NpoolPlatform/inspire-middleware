package achievement

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	achievement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	achievement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return achievement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
