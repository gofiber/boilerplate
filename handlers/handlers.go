package handlers

import (
	"boilerplate/database"
	"boilerplate/models"

	"github.com/gofiber/fiber/v2"
)

// UserGet returns a user
func UserList(c *fiber.Ctx) error {
	users := database.Get()
	return c.JSON(fiber.Map{
		"success": true,
		"user":    users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		Name: c.FormValue("user"),
	}
	database.Insert(user)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
