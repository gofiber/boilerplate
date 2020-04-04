package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
)

func main() {
	// Connect with database
	database.Connect()

	// App settings -> https://fiber.wiki/application#settings
	app := fiber.New(&fiber.Settings{
		TemplateFolder:    "./views",
		TemplateExtension: ".html",
	})

	// https://fiber.wiki/middleware#recover
	app.Use(recover.New(recover.Config{
		Handler: routes.Error,
	}))

	// https://fiber.wiki/middleware#logger
	app.Use(logger.New())

	// https://fiber.wiki/middleware#helmet
	app.Use(helmet.New())

	// Routes -> https://fiber.wiki/application#http-methods
	app.Get("/", routes.Index)
	app.Get("/json", routes.JSON)

	// Static -> https://fiber.wiki/application#static
	app.Static("/", "./public")

	// Not found -> https://github.com/gofiber/fiber#custom-404-response
	app.Use(routes.NotFound)

	// Listen -> https://fiber.wiki/application#listen
	app.Listen(3000)
}
