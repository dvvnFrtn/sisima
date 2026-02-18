package main

import (
	"os"

	"github.com/dvvnFrtn/sisima/internal/app"
)

func main() {
	app := app.New(app.Config{
		EnableLogger: true,
	})

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
