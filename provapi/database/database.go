package database

import (
	"context"
	"database/sql"
	"fmt"

	apiCfg "github.com/pirogoeth/pve-tools/provapi/config"
	"github.com/pirogoeth/pve-tools/provapi/database/sqlite3"
	dbTypes "github.com/pirogoeth/pve-tools/provapi/database/types"
)

var database *sql.DB = nil

func Init(ctx context.Context, cfg *apiCfg.Config) error {
	var err error
	var opener dbTypes.Opener

	switch cfg.Database.Type {
	case "sqlite3":
		opener = sqlite3.Open
	default:
		return fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}

	database, err = opener(ctx, cfg.Database.Config)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	return nil
}

func Database() *sql.DB {
	if database == nil {
		panic("database not initialized")
	}

	return database
}
