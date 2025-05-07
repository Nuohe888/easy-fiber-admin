package system

import (
	"easy-fiber-admin/module/system/internal/controller"
	"easy-fiber-admin/module/system/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	r.Get("ping", controller.ApiCtl.Ping)

	r.Post("/auth/login", controller.UserCtl.Login)
	r.Post("/auth/refresh", controller.UserCtl.Refresh)

	auth := r.Group("")
	auth.Use(middleware.JWT()).
		Use(middleware.Casbin())
	auth.Get("/auth/codes", controller.UserCtl.Codes)
	auth.Post("/auth/logout", controller.UserCtl.Logout)
	auth.Get("/user/info", controller.UserCtl.Info)

	auth.Put("user/:id", controller.UserCtl.Put)
	auth.Post("user", controller.UserCtl.Add)
	auth.Delete("user/:id", controller.UserCtl.Del)
	auth.Get("user", controller.UserCtl.Get)
	auth.Get("user/list", controller.UserCtl.List)

	auth.Put("role/:id", controller.RoleCtl.Put)
	auth.Post("role", controller.RoleCtl.Add)
	auth.Delete("role/:id", controller.RoleCtl.Del)
	auth.Get("role", controller.RoleCtl.Get)
	auth.Get("role/list", controller.RoleCtl.List)
	auth.Get("role/list/all", controller.RoleCtl.ListAll)

	auth.Get("dict", controller.DictTypeCtl.Dict)

	auth.Post("dictType", controller.DictTypeCtl.Add)
	auth.Delete("dictType/:id", controller.DictTypeCtl.Del)
	auth.Put("dictType/:id", controller.DictTypeCtl.Put)
	auth.Get("dictType", controller.DictTypeCtl.Get)
	auth.Get("dictType/list", controller.DictTypeCtl.List)

	auth.Post("dictData", controller.DictDataCtl.Add)
	auth.Delete("dictData/:id", controller.DictDataCtl.Del)
	auth.Put("dictData/:id", controller.DictDataCtl.Put)
	auth.Get("dictData", controller.DictDataCtl.Get)
	auth.Get("dictData/list", controller.DictDataCtl.List)
}
