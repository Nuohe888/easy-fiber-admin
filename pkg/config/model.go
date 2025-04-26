package config

import (
	"go-server/pkg/logger"
	"go-server/pkg/server"
	"go-server/pkg/sql"
)

type Config struct {
	Server server.Config `toml:"server"`
	Sql    sql.Config    `toml:"sql"`
	Log    logger.Config `toml:"log"`
}

var cfg Config

func Get() *Config {
	return &cfg
}
