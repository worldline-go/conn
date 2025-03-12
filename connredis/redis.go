package connredis

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/worldline-go/conn/tlscfg"
)

type Config struct {
	ClientName string   `cfg:"client_name"`
	Address    []string `cfg:"address"`
	UserName   string   `cfg:"username"`
	Password   string   `cfg:"password"`

	TLS tlscfg.TLSConfig `cfg:"tls"`
}

func New(cfg Config) (redis.UniversalClient, error) {
	tlsConfig, err := cfg.TLS.Generate()
	if err != nil {
		return nil, err
	}

	if len(cfg.Address) == 0 {
		return nil, errors.New("no address provided")
	}

	if len(cfg.Address) > 1 {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      cfg.Address,
			Username:   cfg.UserName,
			Password:   cfg.Password,
			ClientName: cfg.ClientName,
			TLSConfig:  tlsConfig,
		}), nil
	} else {
		return redis.NewClient(&redis.Options{
			Addr:       cfg.Address[0],
			Username:   cfg.UserName,
			Password:   cfg.Password,
			ClientName: cfg.ClientName,
			TLSConfig:  tlsConfig,
		}), nil
	}
}
