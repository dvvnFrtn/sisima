package main

import (
	"os"

	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/logger"
	model "github.com/dvvnFrtn/sisima/internal/models"
	route "github.com/dvvnFrtn/sisima/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.ConnectDatabase()
	model.Migrate()

	if config.IsDevelopment() {
		logger.InitSQLite()
	}

	app := fiber.New()

	app.Use(logger.HTTPLogger())

	route.RegisterRoutes(app)

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
