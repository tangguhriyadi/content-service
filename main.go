package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	route "github.com/tangguhriyadi/content-service/internals/server/http"
)

func main() {
	containerConf := container.InitContainer()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	route.HttpRouteInit(app, containerConf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":8082"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here

}
