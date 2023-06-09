package schema

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pirogoeth/pve-tools/pkg/logging"
	"github.com/pirogoeth/pve-tools/provapi/model/schemamanager"
)

const (
	// SchemaVersion is the current version of the schema table.
	SchemaVersion = 1
	// SchemaName is the name of the schema version within the schema table.
	SchemaName = "schema"
	// SchemaDescription is a brief discription of the schema.
	SchemaDescription = "table used by migration system for storing schemas"
)

// Schema is a representation of an item in the schema table.
type Schema struct {
	Name        string
	Description string
	Version     int
}

func Init(ctx context.Context, db *sql.DB) error {
	// For the sake of getting somewhere, let's _ALWAYS_ make sure the first
	// migration has been run, straight from the beginning.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := applyMigration(ctx, tx, 1); err != nil {
		return fmt.Errorf("failed to apply initial migration: %w", err)
	}

	logging.Errlog(tx.Commit())

	curSchema, err := GetSchema(ctx, db, SchemaName)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("failed to get schema version: %w", err)
		}

		curSchema = &Schema{
			Name:    SchemaName,
			Version: 1,
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
		schema := `create table if not exists schema (
			name text not null primary key unique,
			version integer not null
		);`
		if _, err := tx.ExecContext(ctx, schema); err != nil {
			return fmt.Errorf("failed to create imagefamily table: %w", err)
		}
	default:
		return fmt.Errorf("unknown schema version: %d", versionNumber)
	}

	return nil
}

func GetSchema(ctx context.Context, db *sql.DB, name string) (*Schema, error) {
	var schema Schema
	if err := db.QueryRowContext(ctx, "select name, version from schema where name = ?", name).Scan(&schema.Name, &schema.Version); err != nil {
		return nil, err
	}

	return &schema, nil
}
