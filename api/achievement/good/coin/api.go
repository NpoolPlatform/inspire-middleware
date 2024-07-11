package goodcoinachievement

import (
	"context"

	goodcoinachievement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/good/coin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodcoinachievement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodcoinachievement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return goodcoinachievement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
