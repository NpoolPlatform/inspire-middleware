package invitation

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

	_ "github.com/google/uuid"
)

func GetInvitees(ctx context.Context, appID, userID string, offset, limit int32) ([]*npool.Invitation, uint32, error) {
	return nil, 0, fmt.Errorf("NOT IMPLEMENTED")
}

func GetInviters(ctx context.Context, appID, userID string, offset, limit int32) ([]*npool.Invitation, uint32, error) {
	return nil, 0, fmt.Errorf("NOT IMPLEMENTED")
}
