package config

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	config.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	config.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return config.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
