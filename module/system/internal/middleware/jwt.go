package middleware

import (
	"context"
	"easy-fiber-admin/module/system/internal/utils"
	"easy-fiber-admin/module/system/internal/vo"
	"easy-fiber-admin/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := utils.GetUserToken(c)
		if err != nil {
			return c.Status(200).JSON(fiber.Map{
				"code": 401,
				"msg":  err.Error(),
			})
		}
		claims, err := jwt.VerifyToken[*vo.UserInfoJwtClaims](token)
		if err != nil {
			return c.Status(200).JSON(fiber.Map{
				"code": 401,
				"msg":  err.Error(),
			})
		}
		c.SetUserContext(context.WithValue(c.UserContext(), "user", claims))
		return c.Next()
	}
}
