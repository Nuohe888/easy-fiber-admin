package vo

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

const (
	ERROR   = 400
	SUCCESS = 0
)

func ResultErr(err error, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(200).JSON(Response{
			Code:    ERROR,
			Data:    nil,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(Response{
		Code:    SUCCESS,
		Data:    nil,
		Message: "",
	})
}

func ResultOK(data any, c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{
		Code:    SUCCESS,
		Data:    data,
		Message: "ok",
	})
}

func BodyParser(body any, c *fiber.Ctx) error {
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(200).JSON(Response{
			Code:    ERROR,
			Data:    nil,
			Message: "参数错误",
		})
	}
	return nil
}

type List struct {
	Items any   `json:"items"`
	Total int64 `json:"total"`
}
