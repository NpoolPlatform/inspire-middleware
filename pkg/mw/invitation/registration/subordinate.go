//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

func CreateSubordinateProcedure(ctx context.Context) error {
	conn, err := mysql.GetConn()
	if err != nil {
		return wlog.WrapError(err)
	}

	const procedure = `
		DROP PROCEDURE IF EXISTS get_subordinates;
		SET SESSION GROUP_CONCAT_MAX_LEN = 1024000;
		CREATE PROCEDURE get_subordinates (IN inviters TEXT)
		BEGIN
		  DECLARE subordinates TEXT;
		  DECLARE my_inviters TEXT;
		  SET subordinates = 'N/A';
		  SET my_inviters = inviters;
		  WHILE my_inviters is not null DO
		    if subordinates = 'N/A' THEN
			  SET subordinates = my_inviters;
			else
			  SET subordinates = CONCAT(subordinates, ',', my_inviters);
			END if;
		    SELECT GROUP_CONCAT(DISTINCT invitee_id) INTO my_inviters FROM registrations WHERE FIND_IN_SET(inviter_id, my_inviters) AND deleted_at=0;
		  END WHILE;
		  SELECT subordinates;
		END
	`
	_, err = conn.ExecContext(ctx, procedure)
	if err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *queryHandler) getInviterIDs(ctx context.Context) error {
	if h.Conds.InviterIDs == nil {
		return wlog.Errorf("invalid inviterids")
	}

	inviterIDs, ok := h.Conds.InviterIDs.Val.([]uuid.UUID)
	if !ok {
		return wlog.Errorf("invalid inviterids")
	}
	_inviterIDs := ""
	for _, id := range inviterIDs {
		if _inviterIDs != "" {
			_inviterIDs = fmt.Sprintf("%v,", _inviterIDs)
		}
		_inviterIDs = fmt.Sprintf("%v%v", _inviterIDs, id)
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		rows, err := cli.QueryContext(
			ctx,
			fmt.Sprintf("CALL get_subordinates(\"%v\")", _inviterIDs),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		defer rows.Close()

		subordinates := ""
		for rows.Next() {
			if err := rows.Scan(&subordinates); err != nil {
				return wlog.WrapError(err)
			}
		}

		__inviterIDs := strings.Split(subordinates, ",") //nolint
		for _, id := range __inviterIDs {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			inviterIDs = append(inviterIDs, _id)
		}
		return nil
	})
	if err != nil {
		return wlog.WrapError(err)
	}

	h.Conds.InviterIDs.Val = inviterIDs

	return nil
}

func (h *Handler) GetSubordinates(ctx context.Context) ([]*npool.Registration, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	if err := handler.getInviterIDs(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	infos, total, err := h.GetRegistrations(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	return infos, total, nil
}
