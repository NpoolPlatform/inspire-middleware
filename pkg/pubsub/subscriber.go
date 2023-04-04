package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	pubsubmsgcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/pubsubmessage"
	entpubsubmsg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessage"

	"github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	msgpb "github.com/NpoolPlatform/message/npool/basetypes/v1"

	registrationmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/google/uuid"
)

func postHandler(ctx context.Context, tx *ent.Tx, msgID string, uid uuid.UUID, err error) error {
	// TODO: dispatch resp according to err
	return nil
}

func prepare(msgID string, body []byte) (interface{}, error) {
	var req interface{}

	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		_req := registrationmgrpb.RegistrationReq{}
		if err := json.Unmarshal(body, &req); err != nil {
			logger.Sugar().Errorw(
				"handler",
				"MsgID", msgID,
				"Sender", sender,
				"Uid", uid,
				"RespToID", respToID,
				"Body", string(body),
			)
			// For message with invalid body, nobody can process it, so just act directly
			return nil, nil
		}
		req = &_req
	default:
		return nil, nil
	}

	return req, nil
}

func stat(ctx context.Context, msgID string, uid uuid.UUID, respToID *uuid.UUID) (bool, error) {
	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		// TODO: here we need to decide if we need to ack, or resp a fail
		if err := checkMessageState(ctx, msgID, uid, respToID); err != nil {
			return err
		}
	default:
		return nil
	}
}

func process(ctx context.Context, req interface{}) error {
	var err error

	switch msgID {
	case msgpb.MessageID_CreateRegistrationInvitationReq.String():
		err = registration.CreateRegistrationV2(
			ctx,
			req.(*registrationmgrpb.RegistrationReq),
			func(ctx context.Context, tx *ent.Tx, err error) error {
				return postHandler(ctx, tx, msgID, uid, err)
			})
	default:
		return nil
	}

	return err
}

func handler(ctx context.Context, msgID, sender string, uid uuid.UUID, respToID *uuid.UUID, body []byte) error {
	req, err := prepare(msgID, body)
	if err != nil {
		return err
	}
	if req == nil {
		return nil
	}

	appliable, err := stat(ctx, msgID, uid, respToID)
	if err != nil {
		return err
	}
	if !appliable {
		return nil
	}

	return process(ctx, req)
}

func Subscribe(ctx context.Context) error {
	return pubsub.Subscribe(ctx, handler)
}

func Shutdown() error {
	return nil
}
