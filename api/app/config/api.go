package config

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
	"google.golang.org/grpc"
)

type Server struct {
	config.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	config.RegisterMiddlewareServer(server, &Server{})
}
