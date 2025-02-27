package main

import (
	// "linkjo/app/models"
	"linkjo/config"
	"linkjo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	// config.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	routes.AuthRoutes(app)
	routes.SetupProductRoutes(app)

	app.Listen(":8080")
}
