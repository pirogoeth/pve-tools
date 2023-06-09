package imagefamily

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pirogoeth/pve-tools/provapi/model/schema"
	"github.com/pirogoeth/pve-tools/provapi/model/schemamanager"
)

const (
	// SchemaSchema is the schema for the schema table.
	// SchemaVersion is the current version of the schema table.
	SchemaVersion = 1
	// SchemaName is the name of the schema version within the schema table.
	SchemaName = "imagefamily"
)

// Init initializes the imagefamily table, performs migrations, etc.
func Init(ctx context.Context, db *sql.DB) error {
	curSchema, err := schema.GetSchema(ctx, db, SchemaName)
	if err != nil {
		if err == sql.ErrNoRows {
			curSchema = &schema.Schema{
				Name:    SchemaName,
				Version: 1,
			}
		} else {
			return fmt.Errorf("failed to get schema version: %w", err)
		}
	}

	smCfg := &schemamanager.Config{
		SchemaName:           SchemaName,
		SchemaCurrentVersion: curSchema.Version,
		SchemaLatestVersion:  SchemaVersion,
		MigrationFn:          applyMigration,
	}
	if err := schemamanager.Configure(ctx, db, smCfg); err != nil {
		return fmt.Errorf("schema manager failed to configure schema for %s: %w", SchemaName, err)
	}

	return nil
}

func applyMigration(ctx context.Context, tx *sql.Tx, versionNumber int) error {
	switch versionNumber {
	case 1:
		// Initial migration. Create the table.
		schema := `create table if not exists imagefamily (
			name text not null primary key unique,
			description text,
			templatePattern text not null
		);`
		if _, err := tx.ExecContext(ctx, schema); err != nil {
			return fmt.Errorf("failed to create imagefamily table: %w", err)
		}
	default:
		return fmt.Errorf("unknown schema version: %d", versionNumber)
	}

	return nil
}
