package routes

import (
	"linkjo/app/controllers"
	"linkjo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App) {
	orders := app.Group("/orders", middlewares.JWTMiddleware(), middlewares.ExtractTenantID())
	orders.Post("/", controllers.CreateOrder)
	orders.Get("/", controllers.GetOrders)
	orders.Put("/:id", controllers.UpdatePaymentStatus)

	statistics := app.Group("/statistics", middlewares.JWTMiddleware(), middlewares.ExtractTenantID())
	statistics.Get("/", controllers.GetStatistics)
}
