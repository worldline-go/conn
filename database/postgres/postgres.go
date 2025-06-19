package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	ConnMaxLifetime = 15 * time.Minute
	MaxIdleConns    = 3
	MaxOpenConns    = 3
)

// Connect attempts to connect to database server.
func Connect(ctx context.Context, dbDataSource string) (*sqlx.DB, error) {
	pool, err := pgxpool.New(ctx, dbDataSource)
	if err != nil {
		return nil, err
	}

	dbStd := stdlib.OpenDBFromPool(pool)
	db := sqlx.NewDb(dbStd, "pgx")

	db.SetConnMaxLifetime(ConnMaxLifetime)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return db, nil
}
