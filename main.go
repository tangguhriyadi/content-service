package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tangguhriyadi/content-service/docs"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/container"
	route "github.com/tangguhriyadi/content-service/internals/server/http"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	containerConf := container.InitContainer()

	baseUrl := "18.140.2.142:8000"

	docs.SwaggerInfo.Title = "Content Service Dapur Santet"
	docs.SwaggerInfo.Description = "test"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = baseUrl
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	app := fiber.New()
	app.Use(cors.New())
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
