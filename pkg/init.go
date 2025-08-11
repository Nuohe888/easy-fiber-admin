package pkg

import (
	"easy-fiber-admin/pkg/casbin"
	"easy-fiber-admin/pkg/config"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/redis"
	"easy-fiber-admin/pkg/server"
	"easy-fiber-admin/pkg/sql"
	"easy-fiber-admin/pkg/sqlite"
)

func Init() {
	config.Init()
	cfg := config.Get()

	logger.Init(&cfg.Log)

	sql.Init(&cfg.Sql)
	sqlite.Init(&cfg.Sqlite)
	redis.Init(&cfg.Redis)

	casbin.Init(sql.Get())

	casbin.GetAdmin().AddPolicy("admin", "*", "*")

	server.Init(cfg.Server.Port)
}
