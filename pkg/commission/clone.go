package commission

import (
	"context"

	gop "github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"
)

func CloneCommissions(ctx context.Context, appID, fromGoodID, toGoodID, value string) error {
	if err := gop.CloneGoodOrderPercents(ctx, appID, fromGoodID, toGoodID, value); err != nil {
		return err
	}
	return nil
}
