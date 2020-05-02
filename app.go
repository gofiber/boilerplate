package main

import (
	"boilerplate/controllers"
	"boilerplate/database"
	"boilerplate/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
)

func main() {
	// Connected with database
	database.Connect()
	// Create fiber app
	app := fiber.New()
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	// Test recover middleware
	app.Get("/panic", func(c *fiber.Ctx) {
		panic("Hi I'm an error!")
	})
	// Register user routes
	routes.User(app)
	// Setup static files
	app.Static("/", "./public")
	// Handle not founds
	app.Use(controllers.NotFound)
	// Listen on port 3000
	app.Listen(3000)
}
