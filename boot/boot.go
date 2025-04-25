package boot

import (
	"go-server/module/system"
	"go-server/pkg/server"
)

func Boot() {
	//初始化
	initBoot()

	system.Router(server.Get().Group("/api/admin"))

	//运行Server
	server.Start()
}
