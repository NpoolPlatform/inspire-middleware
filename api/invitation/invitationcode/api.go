package invitationcode

import (
	"context"

	invitationcode1 "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	invitationcode1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	invitationcode1.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return invitationcode1.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}