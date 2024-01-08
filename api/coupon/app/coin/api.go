package coin

import (
	coin1 "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"google.golang.org/grpc"
)

type Server struct {
	coin1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coin1.RegisterMiddlewareServer(server, &Server{})
}
