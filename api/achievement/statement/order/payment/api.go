package orderpaymentstatement

import (
	"context"

	orderpaymentstatement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	orderpaymentstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	orderpaymentstatement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return orderpaymentstatement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
