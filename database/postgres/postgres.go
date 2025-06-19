package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/worldline-go/conn/database"
)

func init() {
	database.DBConnections["pgx"] = Connect
}

// Connect attempts to connect to database server.
func Connect(ctx context.Context, dbType, dbDataSource string) (*sqlx.DB, error) {
	if dbType != "pgx" {
		return nil, database.ErrUnsupportedDBType
	}

	pool, err := pgxpool.New(ctx, dbDataSource)
	if err != nil {
		return nil, err
	}

	dbStd := stdlib.OpenDBFromPool(pool)
	db := sqlx.NewDb(dbStd, dbType)

	return db, nil
}
