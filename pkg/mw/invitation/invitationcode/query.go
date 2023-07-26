package invitationcode

import (
	"context"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/invitation/invitationcode"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func GetInvitationCode(ctx context.Context, id string) (*mgrpb.InvitationCode, error) {
	return mgrcli.GetInvitationCode(ctx, id)
}

func GetInvitationCodes(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.InvitationCode, uint32, error) {
	return mgrcli.GetInvitationCodes(ctx, conds, offset, limit)
}

func GetInvitationCodeOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.InvitationCode, error) {
	return mgrcli.GetInvitationCodeOnly(ctx, conds)
}
