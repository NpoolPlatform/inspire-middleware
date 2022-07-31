package invitation

import (
	invitation1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"
)

func Ent2Grpc(row *invitation1.Invitation) *npool.Invitation {
	if row == nil {
		return nil
	}

	return &npool.Invitation{
		InviterID: row.InviterID.String(),
		InviteeID: row.InviteeID.String(),
		CreatedAt: row.CreatedAt,
		Kol:       row.InvitationCode != "",
	}
}

func Ent2GrpcMany(rows []*invitation1.Invitation) []*npool.Invitation {
	infos := []*npool.Invitation{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
