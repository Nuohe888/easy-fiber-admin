package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internal/service"
	"go-server/module/system/internal/vo"
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

func (i *apiCtl) Dict(c *fiber.Ctx) error {
	return vo.ResultOK(c.JSON([]any{
		map[string]interface{}{
			"label": "禁用",
			"value": 0,
			"color": "warning",
		},
		map[string]interface{}{
			"label": "启用",
			"value": 1,
			"color": "purple",
		},
	}), c)

}
