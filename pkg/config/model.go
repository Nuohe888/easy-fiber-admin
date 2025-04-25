package config

import (
	"go-server/pkg/server"
	"go-server/pkg/sql"
)

type Config struct {
	Server server.Cfg `toml:"server"`
	Sql    sql.Cfg    `toml:"sql"`
}

var cfg Config

func Get() *Config {
	return &cfg
}
