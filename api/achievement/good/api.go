package goodachievement

import (
	"context"

	goodachievement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/good"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodachievement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodachievement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return goodachievement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
