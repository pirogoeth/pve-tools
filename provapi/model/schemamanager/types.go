package schemamanager

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	ErrInitialCreate    = fmt.Errorf("initial table creation needed")
	ErrMigrationsNeeded = fmt.Errorf("DB migrations needed")
)

type Migrator func(context.Context, *sql.Tx, int) error

type Config struct {
	SchemaName           string
	SchemaCurrentVersion int
	SchemaLatestVersion  int
	MigrationFn          Migrator
}
