package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App
var port int

func Init(_port int) {
	app = fiber.New()
	port = _port
}

func Get() *fiber.App {
	return app
}

func Start() {
	app.Listen(fmt.Sprintf(":%d", port))
}
