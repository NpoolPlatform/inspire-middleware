package accounting

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"

	"google.golang.org/grpc"
)

type Server struct {
	accounting.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	accounting.RegisterMiddlewareServer(server, &Server{})
}
