package system

import (
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internal/controller"
	"go-server/module/system/internal/middleware"
)

func Router(r fiber.Router) {
	r.Get("ping", controller.ApiCtl.Ping)

	r.Post("/login", controller.UserCtl.Login)
	auth := r.Group("/auth")
	auth.Use(middleware.UserJwt())
	auth.Get("/codes", controller.UserCtl.Codes)
	auth.Post("/logout", controller.UserCtl.Logout)
	auth.Get("/user/info", controller.UserCtl.Info)
}
