package routes

import (
	"boilerplate/controllers"

	"github.com/gofiber/fiber"
)

// User registers all user routes
func User(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/:id", controllers.UserGet)
	user.Post("/", controllers.UserCreate)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDelete)
}
