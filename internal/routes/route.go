package route

import (
	handler "github.com/dvvnFrtn/sisima/internal/handlers"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	// index resource
	app.Get("/", handler.IndexHandler)

	// student resources
	service := service.NewStudentService()
	handler := handler.NewStudentHandler(service)

	resource := app.Group("/student")
	resource.Get("", handler.FindAllPaginated)
	resource.Get("/:id", handler.FindDetailById)
	resource.Post("", handler.Create)
	resource.Delete("/:id", handler.Delete)
}
