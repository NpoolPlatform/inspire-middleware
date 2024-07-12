package reward

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	reward.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	reward.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return reward.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
