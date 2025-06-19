# Conn

Connection to resources.

```sh
go get github.com/worldline-go/conn
```

## Redis

Redis connection return `redis.UniversalClient` which is a client that can connect to multiple Redis servers.

```go
// import "github.com/worldline-go/conn/connredis"
client, err := connredis.New(connredis.Config{})
```

## Postgres

Postgres connection using `pgxpool`, this help for refresh connection when database is restarted.

```go
// import "github.com/worldline-go/conn/database/postgres"
db, err := postgres.Connect(ctx, "postgres://postgres@localhost:5432/postgres")
```
