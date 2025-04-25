package pkg

import (
	"go-server/pkg/config"
	"go-server/pkg/server"
	"go-server/pkg/sql"
)

func Init() {
	config.Init()
	cfg := config.Get()

	server.Init(cfg.Server.Port)

	sql.Init(cfg.Sql.User, cfg.Sql.Pass, cfg.Sql.Host, cfg.Sql.DbName,
		cfg.Sql.Port, cfg.Sql.MaxIdleConns, cfg.Sql.MaxOpenConns)
}
