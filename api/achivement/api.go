package achivement

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"

	"google.golang.org/grpc"
)

type Server struct {
	achivement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	achivement.RegisterMiddlewareServer(server, &Server{})
}
