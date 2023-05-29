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
)

func main() {
	containerConf := container.NewContainer()
	containerConf.Init()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

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
