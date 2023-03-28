package main

import (
	"github.com/NpoolPlatform/inspire-middleware/api"
	msgcli "github.com/NpoolPlatform/inspire-middleware/pkg/message/client"
	msglistener "github.com/NpoolPlatform/inspire-middleware/pkg/message/listener"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	registration "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		if err := registration.CreateSubordinateProcedure(c.Context); err != nil {
			return err
		}
		if err := registration.CreateSuperiorProcedure(c.Context); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()
		if err := msgcli.Init(); err != nil {
			return err
		}

		go msglistener.Listen()

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apicli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	_ = apicli.Register(mux)
	return nil
}
