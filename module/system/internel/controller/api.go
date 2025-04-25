package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internel/service"
)

type apiCtl struct {
	srv *service.ApiSrv
}

var ApiCtl *apiCtl

func InitApiCtl() {
	ApiCtl = &apiCtl{
		srv: service.GetApiSrv(),
	}
}

func (i *apiCtl) Ping(c *fiber.Ctx) error {
	err := i.srv.Ping()
	if err != nil {
		return c.JSON(fiber.Map{"code": "400", "msg": err.Error()})
	}
	return c.JSON(fiber.Map{"code": "200"})
}
