package watcher

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/pubsub"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func Watch(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			err := pubsub.AdjustmentMessageState()
			if err != nil {
				logger.Sugar().Errorw("Watch updatePubsubMessageState", "Error", err)
				return
			}
			if ctx.Err() == nil {
				logger.Sugar().Infow("Watch", "State", "Done")
				return
			}
			logger.Sugar().Errorw("Watch", "State", "Error", "Error", ctx.Err())
			return
		}
	}
}
