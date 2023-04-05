package handler

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

type MsgCommiter func(ctx context.Context, tx *ent.Tx, err error) error
