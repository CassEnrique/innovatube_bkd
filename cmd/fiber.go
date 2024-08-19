package main

import (
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newHTTP() *fiber.App {
	f := fiber.New()

	f.Use(logger.New())
	f.Use(recover.New())

	corsConfig := cors.Config{
		AllowOrigins: os.Getenv("ALLOWED_ORIGINS"),
		AllowMethods: os.Getenv("ALLOWED_METHODS"),
	}

	f.Use(cors.New(corsConfig))

	return f
}
