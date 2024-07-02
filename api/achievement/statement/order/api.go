package orderstatement

import (
	"context"

	orderstatement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	orderstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	orderstatement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return orderstatement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
