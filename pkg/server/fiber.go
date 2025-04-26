package server

import (
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func newFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
		JSONDecoder:       gojson.Unmarshal,
		JSONEncoder:       gojson.Marshal,
	})
	app.Use(recover2.New(recover2.Config{
		EnableStackTrace:  true,
		StackTraceHandler: recover2.ConfigDefault.StackTraceHandler,
	}))
	app.Use(cors.New())
	app.Get("/metrics", monitor.New(monitor.ConfigDefault))
	app.Use(logger.New(logger.Config{
		Format:        "${time} | ${green}${status}${reset} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:    "15:04:05",
		Output:        os.Stdout,
		DisableColors: true,
	}))
	return app
}
