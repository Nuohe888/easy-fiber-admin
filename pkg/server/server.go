package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var app *fiber.App
var port int

func Init(_port int) {
	app = newFiber()
	port = _port
}

func Get() *fiber.App {
	return app
}

func Start() {
	app.Listen(fmt.Sprintf(":%d", port))
}

func Stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭服务器中 ...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("服务器退出失败:", err)
	}
	log.Println("服务器已关闭...")
}
