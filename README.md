# Conn

Connection to resources.

```sh
go get github.com/worldline-go/conn
```

## Database

| Database | Type | Package                                        |
| -------- | ---- | ---------------------------------------------- |
| Postgres | pgx  | github.com/worldline-go/conn/database/postgres |

Postgres connection using `pgxpool`, this help for refresh connection when database is restarted.

```go
// "github.com/worldline-go/conn/database"
// _ "github.com/worldline-go/conn/database/postgres"

db, err := database.Connect(ctx, "pgx", "postgres://postgres@localhost:5432/postgres")
```

## Redis

Redis connection return `redis.UniversalClient` which is a client that can connect to multiple Redis servers.

```go
// import "github.com/worldline-go/conn/connredis"
client, err := connredis.New(connredis.Config{})
```
