package statement

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

	"google.golang.org/grpc"
)

type Server struct {
	statement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	statement.RegisterMiddlewareServer(server, &Server{})
}
