package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1"

	"github.com/NpoolPlatform/inspire-middleware/api/achievement"
	"github.com/NpoolPlatform/inspire-middleware/api/achievement/statement"
	"github.com/NpoolPlatform/inspire-middleware/api/calculate"
	"github.com/NpoolPlatform/inspire-middleware/api/commission"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/allocated"
	scope1 "github.com/NpoolPlatform/inspire-middleware/api/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/scope"
	"github.com/NpoolPlatform/inspire-middleware/api/event"
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
	scope.Register(server)
	scope1.Register(server)
	commission.Register(server)
	achievement.Register(server)
	calculate.Register(server)
	event.Register(server)
	statement.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := achievement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := statement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := calculate.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := allocated.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := commission.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coupon.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := allocated.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := event.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := invitationcode.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := registration.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
