package schemamanager

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pirogoeth/pve-tools/pkg/logging"
)

type manager struct {
	cfg *Config
	db  *sql.DB
}

func Configure(ctx context.Context, db *sql.DB, cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("cfg can not be nil")
	}
	if db == nil {
		return fmt.Errorf("db can not be nil")
	}

	m := &manager{
		cfg: cfg,
		db:  db,
	}
	return m.ensureSchema(ctx, db)
}

// EnsureSchema can be used by other models to ensure that their schema is up-to-date.
func (m *manager) ensureSchema(ctx context.Context, db *sql.DB) error {
	// Get the current schema version.
	var currentVersion int
	if err := db.QueryRowContext(ctx, "select version from schema where name = ?", m.cfg.SchemaName).Scan(&currentVersion); err != nil {
		if err == sql.ErrNoRows {
			// The schema hasn't been recorded yet, so we need to insert it.
			if _, err := db.ExecContext(ctx, "insert into schema (name, version) values (?, ?)", m.cfg.SchemaName, m.cfg.SchemaLatestVersion); err != nil {
				return fmt.Errorf("failed to insert schema: %w", err)
			}

			// We need to perform the initial migrations.
			currentVersion = 0
		} else {
			return fmt.Errorf("failed to query schema: %w", err)
		}
	}

	// The schema has been recorded, so we need to check if it's up-to-date.
	if currentVersion != m.cfg.SchemaLatestVersion {
		return m.migrateToSchemaVersion(ctx, db)
	}

	return nil
}

func (m *manager) migrateToSchemaVersion(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	logging.Infof("applying database migration for schema %s, version %d -> %d",
		m.cfg.SchemaName, m.cfg.SchemaCurrentVersion, m.cfg.SchemaLatestVersion)

	for version := m.cfg.SchemaCurrentVersion + 1; version <= m.cfg.SchemaLatestVersion; version++ {
		if err := m.cfg.MigrationFn(ctx, tx, version); err != nil {
			// If we fail to apply a migration, we need to rollback the transaction.
			logging.Errlog(tx.Rollback())
			return fmt.Errorf("failed to apply migration %d: %w", version, err)
		}
	}

	return nil
}
