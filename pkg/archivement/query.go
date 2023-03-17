package archivement

import (
	"context"

	detailmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/archivement/detail"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
)

func GetDetails(ctx context.Context, conds *detailmgrpb.Conds, offset, limit int32) ([]*detailmgrpb.Detail, uint32, error) {
	return detailmgrcli.GetDetails(ctx, conds, offset, limit)
}
