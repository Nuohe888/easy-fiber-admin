package config

import (
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/server"
	"easy-fiber-admin/pkg/sql"
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
