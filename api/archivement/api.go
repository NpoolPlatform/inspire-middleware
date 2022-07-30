package archivement

import (
	"context"

	archivement "github.com/NpoolPlatform/message/npool/inspire/gw/v1/archivement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	archivement.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	archivement.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := archivement.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
