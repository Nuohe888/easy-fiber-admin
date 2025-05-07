package boot

import (
	"easy-fiber-admin/model"
	"easy-fiber-admin/module"
	"easy-fiber-admin/module/system"
	"easy-fiber-admin/pkg/server"
)

func Boot() {
	//初始化
	initBoot()

	//模块初始化
	module.Init()

	//初始化数据库
	model.Init()

	//注册后台路由
	system.Router(server.Get().Group("/api/admin"))

	//运行Server
	go server.Start()
	server.Stop()
}
