package system

import (
	"easy-fiber-admin/module/system/internal/controller"
	middleware2 "easy-fiber-admin/pkg/common/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	r.Post("login", controller.UserCtl.Login)
	r.Post("refresh", controller.UserCtl.Refresh)

	//r.Post("update", controller.ApiCtl.UpdateFile)
	//r.Post("delFile", controller.ApiCtl.DelFile)

	auth := r.Group("auth")
	auth.Use(middleware2.JWT()).
		Use(middleware2.Casbin())

	auth.Post("/file/upload/img", controller.ApiCtl.FileUploadImg)
	auth.Post("logout", controller.UserCtl.Logout)
	auth.Get("userInfo", controller.UserCtl.Info)
	auth.Post("editPassword", controller.UserCtl.EditPassword)

	auth.Put("user/:id", controller.UserCtl.Put)
	auth.Post("user", controller.UserCtl.Add)
	auth.Delete("user/:id", controller.UserCtl.Del)
	auth.Get("user", controller.UserCtl.Get)
	auth.Get("user/list", controller.UserCtl.List)
	auth.Get("user/status", controller.UserCtl.GetStatus)

	auth.Put("role/:id", controller.RoleCtl.Put)
	auth.Post("role", controller.RoleCtl.Add)
	auth.Delete("role/:id", controller.RoleCtl.Del)
	auth.Get("role", controller.RoleCtl.Get)
	auth.Get("role/list", controller.RoleCtl.List)
	auth.Get("role/list/all", controller.RoleCtl.ListAll)
	auth.Get("role/status", controller.RoleCtl.GetStatus)

	auth.Put("user_center/:id", controller.UserCenterCtl.Put)
	auth.Post("user_center", controller.UserCenterCtl.Add)
	auth.Delete("user_center/:id", controller.UserCenterCtl.Del)
	auth.Get("user_center", controller.UserCenterCtl.Get)
	auth.Get("user_center/list", controller.UserCenterCtl.List)
	auth.Get("user_center/status", controller.UserCenterCtl.GetStatus)
}
