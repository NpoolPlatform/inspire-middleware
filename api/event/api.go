package event

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	event.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	event.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return event.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
