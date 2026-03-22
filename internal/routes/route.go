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

	resource.Get("", handler.FindAllPaginated)
	resource.Get("/:id", handler.FindDetailById)
	resource.Post("", handler.Create)
}

func BillingRoutes(app *fiber.App) {
	service := service.NewBillingService()
	handler := handler.NewBillingHandler(service)

	app.Post("/billing-types", handler.CreateBillingType)
	app.Patch("/billing-types/:billing_type_id", handler.UpdateBillingType)
	app.Get("/billing-types", handler.GetAllBillingType)
	app.Get("/billing-types/:billing_type_id", handler.GetBillingType)

	app.Post("/billings", handler.CreateBilling)
	app.Get("/billings", handler.GetAllBilling)
	app.Get("/billings/:billing_id", handler.GetBilling)
}
