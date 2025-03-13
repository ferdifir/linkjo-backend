package main

import (
	// "linkjo/app/models"
	"linkjo/config"
	"linkjo/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	config.ConnectDB()

	// models.Migration()

	app := fiber.New()
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	routes.AuthRoutes(app)
	routes.SetupProductRoutes(app)
	routes.SetupOrderRoutes(app)
	routes.PublicRoutes(app)

	app.Listen(":8080")
}
