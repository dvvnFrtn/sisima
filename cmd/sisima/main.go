package main

import (
	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/logger"
	model "github.com/dvvnFrtn/sisima/internal/models"
	route "github.com/dvvnFrtn/sisima/internal/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	config.ConnectDatabase()
	model.Migrate()

	if config.IsDevelopment() {
		logger.InitSQLite()
	}

	app := fiber.New()

	app.Use(logger.HTTPLogger())

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	route.IndexRoutes(app)
	route.StudentRoutes(app)
	route.BillingRoutes(app)

	app.Listen(":8888")
}
