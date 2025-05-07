package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internal/utils"
	"go-server/module/system/internal/vo"
	"go-server/pkg/jwt"
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
