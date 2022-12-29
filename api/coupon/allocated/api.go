package allocated

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"google.golang.org/grpc"
)

type Server struct {
	allocated.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	allocated.RegisterMiddlewareServer(server, &Server{})
}
