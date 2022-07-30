package api

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire"

	"github.com/NpoolPlatform/inspire-middleware/api/invitation"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterMiddlewareServer(server, &Server{})
	invitation.Register(server)
}

func RegisterMiddleware(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
