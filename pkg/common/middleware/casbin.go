package middleware

import (
	"easy-fiber-admin/pkg/casbin"
	"easy-fiber-admin/pkg/common/utils"
	"easy-fiber-admin/pkg/common/vo"
	"github.com/gofiber/fiber/v2"
)

func Casbin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		info := utils.GetUserInfo(c)
		obj := c.Path()
		act := c.Method()
		enforcer := casbin.GetAdmin()
		ok, err := enforcer.Enforce(info.RoleCode, obj, act)
		if err != nil || !ok {
			return c.Status(200).JSON(vo.Response{
				Code:    403,
				Data:    nil,
				Message: "权限不足",
			})
		}
		return c.Next()
	}
}
