package pkg

import (
	"easy-fiber-admin/pkg/config"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/server"
	"easy-fiber-admin/pkg/sql"
)

func Init() {
	config.Init()
	cfg := config.Get()

	logger.Init(&cfg.Log)

	sql.Init(&cfg.Sql)

	server.Init(cfg.Server.Port)
}
