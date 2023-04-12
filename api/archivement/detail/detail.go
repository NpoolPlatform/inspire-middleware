package detail

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/servicename"
	"go.opentelemetry.io/otel"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/detail"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/archivement/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	scodes "go.opentelemetry.io/otel/codes"
)

func (s *Server) GetDetails(ctx context.Context, in *npool.GetDetailsRequest) (*npool.GetDetailsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateOrder")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	infos, total, err := mgrcli.GetDetails(ctx, in.GetConds(), int32(in.GetOffset()), int32(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetDetails", "error", err)
		return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetDetailsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
