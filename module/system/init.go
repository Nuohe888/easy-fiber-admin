package system

import (
	"go-server/module/system/internal/controller"
	"go-server/module/system/internal/pkg/casbin"
	"go-server/module/system/internal/service"
	"go-server/pkg/sql"
)

func Init() {
	casbin.Init(sql.Get())
	//给admin添加权限
	casbin.Get().AddPolicy("admin", "*", "*")

	service.Init()
	controller.Init()
}
