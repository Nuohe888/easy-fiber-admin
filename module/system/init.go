package system

import (
	"go-server/module/system/internal/controller"
	"go-server/module/system/internal/service"
)

func Init() {
	service.Init()
	controller.Init()
}
