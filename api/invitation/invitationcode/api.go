package invitationcode

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	"google.golang.org/grpc"
)

type Server struct {
	invitationcode.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	invitationcode.RegisterMiddlewareServer(server, &Server{})
}
