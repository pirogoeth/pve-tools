package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pirogoeth/pve-tools/provapi/model/imagefamily"
	"github.com/pirogoeth/pve-tools/provapi/model/schema"
)

func Init(ctx context.Context, db *sql.DB) error {
	if err := schema.Init(ctx, db); err != nil {
		return fmt.Errorf("failed to initialize schema model: %w", err)
	}

	if err := imagefamily.Init(ctx, db); err != nil {
		return fmt.Errorf("failed to initialize imagefamily model: %w", err)
	}

	return nil
}
