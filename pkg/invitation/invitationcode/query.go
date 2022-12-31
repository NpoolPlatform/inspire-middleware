package invitationcode

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func GetInvitationCode(ctx context.Context, appID, userID string) (*mgrpb.InvitationCode, error) {
	return mgrcli.GetInvitationCodeOnly(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
	})
}
