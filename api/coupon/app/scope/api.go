package scope

import (
	appgoodscope "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"google.golang.org/grpc"
)

type Server struct {
	appgoodscope.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appgoodscope.RegisterMiddlewareServer(server, &Server{})
}
