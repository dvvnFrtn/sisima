package main

import (
	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/logger"
	"github.com/dvvnFrtn/sisima/internal/models"
	"github.com/dvvnFrtn/sisima/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDatabase()
	model.Migrate()

	if config.IsDevelopment() {
		logger.InitSQLite()
	}

	app := fiber.New()

	app.Use(logger.HTTPLogger())

	route.IndexRoutes(app)
	route.StudentRoutes(app)

	app.Listen(":8888")
}
