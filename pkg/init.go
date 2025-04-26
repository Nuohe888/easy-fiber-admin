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

	server.Init(cfg.Server.Port)

	sql.Init(&cfg.Sql)
}
