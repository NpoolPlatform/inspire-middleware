package invitationcode

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func CreateInvitationCode(ctx context.Context, in *mgrpb.InvitationCodeReq) (*mgrpb.InvitationCode, error) {
	return mgrcli.CreateInvitationCode(ctx, in)
}