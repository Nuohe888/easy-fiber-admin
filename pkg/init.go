package pkg

import (
	"go-server/pkg/config"
	"go-server/pkg/logger"
	"go-server/pkg/server"
	"go-server/pkg/sql"
)

func Init() {
	config.Init()
	cfg := config.Get()

	logger.Init(&cfg.Log)

	sql.Init(&cfg.Sql)

	server.Init(cfg.Server.Port)
}
