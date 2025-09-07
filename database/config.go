package database

import "time"

type Config struct {
	Datasource string `cfg:"datasource" log:"-"`
	Type       string `cfg:"type"`

	ConnMaxLifetime time.Duration `cfg:"conn_max_lifetime" default:"15m"`
	MaxIdleConns    int           `cfg:"max_idle_conns"    default:"3"`
	MaxOpenConns    int           `cfg:"max_open_conns"    default:"3"`

	Migration Migration `cfg:"migration"`
}

type Migration struct {
	Disabled   bool   `cfg:"disabled"`
	Datasource string `cfg:"datasource" log:"-"`
	Type       string `cfg:"type"`
	Schema     string `cfg:"schema"`
	Table      string `cfg:"table"`
}
