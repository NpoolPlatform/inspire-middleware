package commission

import (
	"context"

	gop "github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"
)

func CloneCommissions(ctx context.Context, appID, fromGoodID, toGoodID string) error {
	if err := gop.CloneGoodOrderPercents(ctx, appID, fromGoodID, toGoodID); err != nil {
		return err
	}
	return nil
}
