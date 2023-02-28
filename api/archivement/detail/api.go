package detail

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/detail"

	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterMiddlewareServer(server, &Server{})
}
