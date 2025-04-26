package module

import (
	"go-server/module/system/internel/controller"
	"go-server/module/system/internel/service"
)

func Init() {
	service.Init()
	controller.Init()
}
