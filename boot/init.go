package boot

import (
	"easy-fiber-admin/pkg"
	"easy-fiber-admin/pkg/config"
	"easy-fiber-admin/plugin"
)

func initBoot() {
	//包初始化
	pkg.Init()

	//插件初始化
	plugin.Init()

	//配置文件固定 数据库动态 代码动态加载任选
	//这里选择的是配置文件
	//如果你想做到图片 视频 文件不同上传位置就可以考虑使用的时候重新Init
	err := plugin.InitStorage(config.Get().Server.Storage)
	if err != nil {
		panic("初始化存储失败: " + err.Error())
	}

	err = plugin.GetStorage().Init("", "", "", "./upload/file", false)
	if err != nil {
		panic("初始化存储失败: " + err.Error())
	}

	//密码这里修改会影响初始化后端用户密码
	//如果要改变这里后台用户Login方法和修改密码都需要改变
	//二次封住可以定义常量 AdminCryptoType UserCryptoType
	err = plugin.InitCrypto(plugin.CryptoTypeSHA256)
	if err != nil {
		panic("初始化密码加密失败: " + err.Error())
	}
}
