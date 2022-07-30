package invitation

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

	"google.golang.org/grpc"
)

type Server struct {
	invitation.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	invitation.RegisterMiddlewareServer(server, &Server{})
}
