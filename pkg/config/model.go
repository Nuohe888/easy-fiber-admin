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
	Redis  RedisConfig   `toml:"redis"`
	Sentry SentryConfig  `toml:"sentry"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type SentryConfig struct {
	Dsn string `toml:"dsn"`
}

var cfg Config

func Get() *Config {
	return &cfg
}
