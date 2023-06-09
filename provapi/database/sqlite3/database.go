package sqlite3

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(ctx context.Context, cfgJson *json.RawMessage) (*sql.DB, error) {
	var cfg Config
	if err := json.Unmarshal(*cfgJson, &cfg); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite3 database: %w", err)
	}

	return db, nil
}
