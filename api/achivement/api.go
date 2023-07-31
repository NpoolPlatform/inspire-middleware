package archivement

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	"google.golang.org/grpc"
)

type Server struct {
	archivement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	archivement.RegisterMiddlewareServer(server, &Server{})
}
