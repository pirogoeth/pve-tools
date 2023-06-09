package types

import (
	"context"
	"database/sql"
	"encoding/json"
)

// Opener is a function that opens a database connection.
type Opener func(ctx context.Context, cfgJson *json.RawMessage) (*sql.DB, error)

// type
