package routes

import (
	"linkjo/app/controllers"
	"linkjo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)

	users := app.Group("/users", middlewares.JWTMiddleware(), middlewares.ExtractTenantID())
	users.Get("/", controllers.GetUsers)
	users.Post("/banner", middlewares.UploadImagesMiddleware, controllers.UpdateBanner)
	users.Patch("/status", controllers.UpdateStatusTenant)
}
