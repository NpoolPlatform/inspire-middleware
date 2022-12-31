package invitationcode

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func GetInvitationCodeOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.InvitationCode, error) {
	return mgrcli.GetInvitationCodeOnly(ctx, conds)
}
