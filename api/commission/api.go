package commission

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"google.golang.org/grpc"
)

type Server struct {
	commission.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	commission.RegisterMiddlewareServer(server, &Server{})
}
