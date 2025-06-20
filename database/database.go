package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ConnMaxLifetime = 15 * time.Minute
	MaxIdleConns    = 3
	MaxOpenConns    = 3
)

var ErrUnsupportedDBType = fmt.Errorf("unsupported database type")

var DBConnections = make(map[string]func(ctx context.Context, dbType, dbDataSource string) (*sqlx.DB, error))

// Connect attempts to connect to database server.
func Connect(ctx context.Context, dbType, dbDataSource string, options ...Option) (*sqlx.DB, error) {
	o := &option{
		ConnMaxLifetime: ConnMaxLifetime,
		MaxIdleConns:    MaxIdleConns,
		MaxOpenConns:    MaxOpenConns,
	}
	for _, opt := range options {
		opt(o)
	}

	var db *sqlx.DB
	if connectFunc, ok := DBConnections[dbType]; ok {
		var err error
		db, err = connectFunc(ctx, dbType, dbDataSource)
		if err != nil {
			return nil, fmt.Errorf("failed to connect using %s: %w", dbType, err)
		}
	}

	if db == nil {
		var err error
		db, err = sqlx.ConnectContext(ctx, dbType, dbDataSource)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
	}

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	return db, nil
}
