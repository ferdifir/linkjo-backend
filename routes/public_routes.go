package routes

import (
	"linkjo/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	public := app.Group("/public")
	public.Get("/products", controllers.GetProductsByTenantID)
	public.Post("/order", controllers.CreateOrderByTenantID)
}
