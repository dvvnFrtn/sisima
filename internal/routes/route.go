package route

import (
	handler "github.com/dvvnFrtn/sisima/internal/handlers"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	// index resource
	app.Get("/", handler.IndexHandler)

	// student resources
	service := service.NewStudentService(db)
	handler := handler.NewStudentHandler(service)

	resource := app.Group("/student")
	resource.Get("", handler.FindAllPaginated)
	resource.Get("/:id", handler.FindDetailById)
	resource.Post("", handler.Create)
	resource.Delete("/:id", handler.Delete)
}
