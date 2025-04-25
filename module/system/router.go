package system

import (
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internel/controller"
)

func Router(r fiber.Router) {
	r.Get("ping", controller.ApiCtl.Ping)
}
