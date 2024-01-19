package cashcontrol

import (
	cashcontrol1 "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"

	"google.golang.org/grpc"
)

type Server struct {
	cashcontrol1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	cashcontrol1.RegisterMiddlewareServer(server, &Server{})
}
