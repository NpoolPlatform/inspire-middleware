package registration

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"google.golang.org/grpc"
)

type Server struct {
	registration.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	registration.RegisterMiddlewareServer(server, &Server{})
}
