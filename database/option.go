package database

import "time"

type option struct {
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns"`
}

type Option func(*option)

// WithConnMaxLifetime sets the maximum lifetime of a connection.
func WithConnMaxLifetime(d time.Duration) Option {
	return func(o *option) {
		o.ConnMaxLifetime = d
	}
}

// WithMaxIdleConns sets the maximum number of idle connections.
func WithMaxIdleConns(n int) Option {
	return func(o *option) {
		o.MaxIdleConns = n
	}
}

// WithMaxOpenConns sets the maximum number of open connections.
func WithMaxOpenConns(n int) Option {
	return func(o *option) {
		o.MaxOpenConns = n
	}
}

// WithMaxConns sets both MaxIdleConns and MaxOpenConns to the same value.
func WithMaxConns(n int) Option {
	return func(o *option) {
		o.MaxIdleConns = n
		o.MaxOpenConns = n
	}
}
