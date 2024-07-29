package coupon

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coupon"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	coupon.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coupon.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return coupon.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
