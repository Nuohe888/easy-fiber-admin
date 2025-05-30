package boot

import (
	"easy-fiber-admin/pkg"
	"easy-fiber-admin/plugin"
)

func initBoot() {
	//包初始化
	pkg.Init()

	//插件初始化
	plugin.Init()

	//放在这里是因为你可以从配置文件加载和数据库加载任选
	plugin.InitStorage("local")
}
