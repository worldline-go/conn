package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var (
	ConnMaxLifetime = 15 * time.Minute
	MaxIdleConns    = 3
	MaxOpenConns    = 3
)

var ErrUnsupportedDBType = fmt.Errorf("unsupported database type")

var DBConnections = make(map[string]func(ctx context.Context, dbType, dbDataSource string) (*sql.DB, error))

func ConnectWithConfig(ctx context.Context, cfg *Config) (*sql.DB, error) {
	opts := []Option{
		WithConnMaxLifetime(cfg.ConnMaxLifetime),
		WithMaxIdleConns(cfg.MaxIdleConns),
		WithMaxOpenConns(cfg.MaxOpenConns),
	}

	return Connect(ctx, cfg.Type, cfg.Datasource, opts...)
}

// Connect attempts to connect to database server.
func Connect(ctx context.Context, dbType, dbDataSource string, opts ...Option) (*sql.DB, error) {
	o := &option{
		ConnMaxLifetime: ConnMaxLifetime,
		MaxIdleConns:    MaxIdleConns,
		MaxOpenConns:    MaxOpenConns,
	}

	for _, opt := range opts {
		opt(o)
	}

	var db *sql.DB
	if connectFunc, ok := DBConnections[dbType]; ok {
		var err error
		db, err = connectFunc(ctx, dbType, dbDataSource)
		if err != nil {
			return nil, fmt.Errorf("failed to connect using %s: %w", dbType, err)
		}
	}

	if db == nil {
		var err error
		db, err = sql.Open(dbType, dbDataSource)
		if err != nil {
			return db, err
		}

		if err = db.PingContext(ctx); err != nil {
			return nil, err
		}
	}

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	return db, nil
}
