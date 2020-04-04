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
	database.Connect()

	app := fiber.New(&fiber.Settings{
		TemplateFolder:    "./views",
		TemplateExtension: ".html",
	})

	app.Use(recover.New(recover.Config{
		Handler: routes.Error,
	}))
	app.Use(logger.New())
	app.Use(helmet.New())

	app.Get("/", routes.Index)
	app.Get("/json", routes.JSON)
	app.Get("/panic", routes.Panic)

	app.Static("/", "./public")

	app.Use(routes.NotFound)

	app.Listen(3000)
}
