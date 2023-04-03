package main

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/action"
	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/api"
	"github.com/NpoolPlatform/inspire-middleware/pkg/feeder"
	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"
	"github.com/NpoolPlatform/inspire-middleware/pkg/watcher"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"

	"github.com/NpoolPlatform/inspire-middleware/pkg/pubsub"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		return action.Run(
			c.Context,
			func(ctx context.Context) error {
				if err := db.Init(); err != nil {
					return err
				}
				if err := registration.CreateSubordinateProcedure(c.Context); err != nil {
					return err
				}
				if err := registration.CreateSuperiorProcedure(c.Context); err != nil {
					return err
				}
				pubsub.Subscrib(ctx)
				return nil
			},
			rpcRegister,
			rpcGatewayRegister,
			func(ctx context.Context) error {
				go watcher.Watch(ctx)
				go feeder.Watch(ctx)
				return nil
			},
		)
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
