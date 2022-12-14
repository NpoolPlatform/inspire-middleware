package coupon

import (
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/coupon"

	"google.golang.org/grpc"
)

type Server struct {
	coupon.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coupon.RegisterMiddlewareServer(server, &Server{})
}
