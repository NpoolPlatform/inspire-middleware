package calculate

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

	"google.golang.org/grpc"
)

type Server struct {
	calculate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	calculate.RegisterMiddlewareServer(server, &Server{})
}
