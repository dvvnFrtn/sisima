package app

import (
	"fmt"

	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/logger"
	route "github.com/dvvnFrtn/sisima/internal/routes"
	"github.com/gofiber/fiber/v3"
)

type Config struct {
	EnableLogger bool
}

func New(cfg Config) *fiber.App {
	db, err := config.ConnectDatabase()
	if err != nil {
		panic("can't connect to database")
	}
	fmt.Println("database successfuly connected")

	app := fiber.New()

	if cfg.EnableLogger {
		app.Use(logger.HTTPLogger())
	}

	route.RegisterRoutes(app, db)

	return app
}
