package invitationcode

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func DeleteInvitationCode(ctx context.Context, id string) (*mgrpb.InvitationCode, error) {
	return mgrcli.DeleteInvitationCode(ctx, id)
}
