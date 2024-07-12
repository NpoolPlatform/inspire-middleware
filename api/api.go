package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1"

	achievement "github.com/NpoolPlatform/inspire-middleware/api/achievement"
	goodachievement "github.com/NpoolPlatform/inspire-middleware/api/achievement/good"
	goodcoinachievement "github.com/NpoolPlatform/inspire-middleware/api/achievement/good/coin"
	orderstatement "github.com/NpoolPlatform/inspire-middleware/api/achievement/statement/order"
	orderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/api/achievement/statement/order/payment"
	"github.com/NpoolPlatform/inspire-middleware/api/achievement/user"
	appcommissionconfig "github.com/NpoolPlatform/inspire-middleware/api/app/commission/config"
	appconfig "github.com/NpoolPlatform/inspire-middleware/api/app/config"
	appgoodcommissionconfig "github.com/NpoolPlatform/inspire-middleware/api/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/api/calculate"
	coinallocated "github.com/NpoolPlatform/inspire-middleware/api/coin/allocated"
	coinconfig "github.com/NpoolPlatform/inspire-middleware/api/coin/config"
	"github.com/NpoolPlatform/inspire-middleware/api/commission"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/allocated"
	cashcontrol "github.com/NpoolPlatform/inspire-middleware/api/coupon/app/cashcontrol"
	scope1 "github.com/NpoolPlatform/inspire-middleware/api/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/api/coupon/scope"
	"github.com/NpoolPlatform/inspire-middleware/api/event"
	"github.com/NpoolPlatform/inspire-middleware/api/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/api/invitation/registration"
	taskconfig "github.com/NpoolPlatform/inspire-middleware/api/task/config"
	taskuser "github.com/NpoolPlatform/inspire-middleware/api/task/user"
	usercoinreward "github.com/NpoolPlatform/inspire-middleware/api/user/coin/reward"
	usercredithistory "github.com/NpoolPlatform/inspire-middleware/api/user/credit/history"
	userreward "github.com/NpoolPlatform/inspire-middleware/api/user/reward"

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
	goodachievement.Register(server)
	achievement.Register(server)
	goodcoinachievement.Register(server)
	calculate.Register(server)
	event.Register(server)
	orderstatement.Register(server)
	orderpaymentstatement.Register(server)
	cashcontrol.Register(server)
	appconfig.Register(server)
	appcommissionconfig.Register(server)
	appgoodcommissionconfig.Register(server)
	user.Register(server)
	taskconfig.Register(server)
	taskuser.Register(server)
	coinconfig.Register(server)
	coinallocated.Register(server)
	usercoinreward.Register(server)
	usercredithistory.Register(server)
	userreward.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := achievement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := goodachievement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := goodcoinachievement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := orderstatement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := orderpaymentstatement.RegisterGateway(mux, endpoint, opts); err != nil {
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
	if err := appgoodcommissionconfig.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := user.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := taskconfig.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := taskuser.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coinconfig.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coinallocated.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usercoinreward.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usercredithistory.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := userreward.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
