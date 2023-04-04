package handler

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	"github.com/google/uuid"
)

type PostHandler func(ctx context.Context, tx *ent.Tx, err error) error
