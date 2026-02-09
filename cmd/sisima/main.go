package main

import (
	"log"
	"os"

	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/logger"
	"github.com/dvvnFrtn/sisima/internal/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	log.Println("ENVIRONMENT raw =", os.Getenv("ENVIRONMENT"))
	log.Println("IsDevelopment =", config.IsDevelopment())

	if config.IsDevelopment() {
		logger.InitSQLite()
	}

	app := fiber.New()
	app.Use(logger.HTTPLogger())

	route.Register(app)
	app.Listen(":8888")
}
