package invitation

import (
	"github.com/google/uuid"
)

type Invitation struct {
	InviterID      uuid.UUID `sql:"inviter_id"`
	InviteeID      uuid.UUID `sql:"invitee_id"`
	CreatedAt      uint32    `sql:"create_at"`
	InvitationCode string    `sql:"invitation_code"`
}
