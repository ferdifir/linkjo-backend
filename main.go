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
	// config.DB.AutoMigrate(&models.Product{})
	// config.DB.AutoMigrate(&models.Categories{})
	// config.DB.AutoMigrate(&models.Order{})
	// config.DB.AutoMigrate(&models.OrderDe{})

	app := fiber.New()
	routes.AuthRoutes(app)
	routes.SetupProductRoutes(app)

	app.Listen(":8080")
}
