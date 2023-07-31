package achievement

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

	"google.golang.org/grpc"
)

type Server struct {
	achievement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	achievement.RegisterMiddlewareServer(server, &Server{})
}
