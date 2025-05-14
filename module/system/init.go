package system

import (
	"easy-fiber-admin/module/system/internal/controller"
	"easy-fiber-admin/module/system/internal/service"
)

func Init() {
	service.Init()
	controller.Init()
}
