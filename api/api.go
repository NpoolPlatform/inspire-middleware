package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1"

	"github.com/NpoolPlatform/inspire-middleware/api/accounting"
	"github.com/NpoolPlatform/inspire-middleware/api/archivement"
	"github.com/NpoolPlatform/inspire-middleware/api/archivement/detail"
	"github.com/NpoolPlatform/inspire-middleware/api/commission"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/allocated"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/coupon"
	"github.com/NpoolPlatform/inspire-middleware/api/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/api/invitation/registration"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterMiddlewareServer(server, &Server{})
	invitationcode.Register(server)
	registration.Register(server)
	coupon.Register(server)
	allocated.Register(server)
	commission.Register(server)
	archivement.Register(server)
	accounting.Register(server)
	detail.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
