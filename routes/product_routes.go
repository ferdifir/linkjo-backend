package routes

import (
	"linkjo/app/controllers"
	"linkjo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	products := app.Group("/products", middlewares.JWTMiddleware(), middlewares.ExtractTenantID())

	products.Get("/", controllers.GetProducts)
	products.Get("/:id", controllers.GetProductByID)
	products.Post("/", controllers.CreateProduct)

	categories := app.Group("/categories")
	categories.Get("/", controllers.GetCategories)
}
