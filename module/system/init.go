package system

import (
	"easy-fiber-admin/module/system/internal/controller"
	"easy-fiber-admin/module/system/internal/pkg/casbin"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/pkg/sql"
)

func Init() {
	casbin.Init(sql.Get())
	//给admin添加权限
	casbin.Get().AddPolicy("admin", "*", "*")

	service.Init()
	controller.Init()
}
