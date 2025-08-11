package config

import (
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/redis"
	"easy-fiber-admin/pkg/server"
	"easy-fiber-admin/pkg/sql"
	"easy-fiber-admin/pkg/sqlite"
)

type Config struct {
	Server server.Config `toml:"server"`
	Sql    sql.Config    `toml:"sql"`
	Sqlite sqlite.Config `toml:"sqlite"`
	Redis  redis.Config  `toml:"redis"`
	Log    logger.Config `toml:"log"`
}

var cfg Config

func Get() *Config {
	return &cfg
}
