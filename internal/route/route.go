package route

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("hello!")
	})
}
