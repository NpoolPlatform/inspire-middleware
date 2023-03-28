package listener

import (
	"fmt"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/inspire-middleware/pkg/message/sub"
)

func sign() {
	for {
		err := sub.Sign(func(msg *sub.SignMsg) error {
			fmt.Println(msg)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}
