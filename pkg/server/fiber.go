package server

import (
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"os"

	"github.com/swaggo/fiber-swagger"
	// swagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
	_ "easy-fiber-admin/docs" // Necessary for swag init to link generated docs

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
)

func newFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
		JSONDecoder:       gojson.Unmarshal,
		JSONEncoder:       gojson.Marshal,
	})

	// Add Sentry middleware
	// Make sure to initialize Sentry before this!
	if sentry.CurrentHub().Client() != nil { // Check if Sentry was initialized
		app.Use(sentryfiber.New(sentryfiber.Options{
			Repanic:         true, // Repanic so Fiber's default error handler can also catch it
			WaitForDelivery: false, // False for not blocking response, true if critical
		}))
	}

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

	// Swagger endpoint
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Static("/upload", "./upload")
	return app
}
