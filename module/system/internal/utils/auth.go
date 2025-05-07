package utils

import (
	"easy-fiber-admin/module/system/internal/vo"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetUserInfo(c *fiber.Ctx) *vo.UserInfoJwtClaims {
	value := c.UserContext().Value("user")
	return value.(*vo.UserInfoJwtClaims)
}

func GetUserToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	if len(authHeader) == 0 {
		return "", errors.New("没有传入Token")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", errors.New("Token格式不正确，必须以Bearer开头")
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)

	if len(token) == 0 {
		return "", errors.New("token为空")
	}

	return token, nil
}
