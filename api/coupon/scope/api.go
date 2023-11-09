package scope

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	"google.golang.org/grpc"
)

type Server struct {
	scope.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	scope.RegisterMiddlewareServer(server, &Server{})
}
