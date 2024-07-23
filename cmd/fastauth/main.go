package main

import (
	"fastauth/config"
	"fastauth/handlers"
	"fastauth/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Implementing server
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("api")
	v1 := api.Group("v1")

	// Loading config
	cfg := config.GetConfig()

	// Getting database
	db := storage.InitDb(cfg)

	v1.Post("token", handlers.CreateTokenHandler(cfg.SecretKey, db))
	v1.Post("register", handlers.RegisterHandler(db))

	// Running server
	app.Listen(cfg.Port)
}