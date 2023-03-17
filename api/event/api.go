package event

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"google.golang.org/grpc"
)

type Server struct {
	event.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	event.RegisterMiddlewareServer(server, &Server{})
}
