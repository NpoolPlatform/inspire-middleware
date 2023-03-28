package sub

import (
	"encoding/json"
	"fmt"
	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	constant "github.com/NpoolPlatform/inspire-middleware/pkg/message/const"
)

type SignMsg struct {
	AppID     string `json:"app_id"`
	InviterID string `json:"inviter_id"`
	InviteeID string `json:"invitee_id"`
}

func Sign(h func(*SignMsg) error) error {
	myClient, err := msgcli.New(constant.ServiceName)
	if err != nil {
		return err
	}
	err = myClient.DeclareSub(constant.ServiceName, "sign")
	if err != nil {
		return err
	}
	msgs, err := myClient.Consume(constant.ServiceName)
	if err != nil {
		return err
	}
	for d := range msgs {
		msg := SignMsg{}
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			return fmt.Errorf("parse message example error: %v", err)
		}

		if h != nil {
			err = h(&msg)
			if err != nil {
				return err
			}
		}
	}

	return fmt.Errorf("WE SHOULD NOT BE HERE")
}
