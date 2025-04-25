package system

import (
	"go-server/module/system/internel/controller"
	"go-server/module/system/internel/service"
)

func init() {
	service.Init()
	controller.Init()
}
