package listener

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/inspire-middleware/pkg/message/sub"
)

func signup() {
	for {
		err := sub.Signup(func(msg *sub.SignupMsg) error {
			fmt.Println(msg)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume signup: %v", err)
			return
		}
	}
}
