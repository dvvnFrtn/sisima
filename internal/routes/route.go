package route

import (
	handler "github.com/dvvnFrtn/sisima/internal/handlers"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
)

func IndexRoutes(app *fiber.App) {
	app.Get("/", handler.IndexHandler)
}

func StudentRoutes(app *fiber.App) {
	service := service.NewStudentService()
	handler := handler.NewStudentHandler(service)

	resource := app.Group("/student")

	resource.Get("/:id", handler.FindDetailById)
	resource.Post("", handler.Create)
}
