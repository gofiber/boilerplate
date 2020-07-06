package handlers

import (
	"boilerplate/database"
	"boilerplate/models"

	"github.com/gofiber/fiber"
)

// UserGet returns a user
func UserList(c *fiber.Ctx) {
	users := database.Get()
	if err := c.JSON(fiber.Map{
		"success": true,
		"user":    users,
	}); err != nil {
		c.Next(err)
	}
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) {
	user := &models.User{
		Name: c.FormValue("user"),
	}
	database.Insert(user)
	if err := c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	}); err != nil {
		c.Next(err)
	}
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) {
	if err := c.Status(404).SendFile("./static/private/404.html"); err != nil {
		c.Next(err)
	}
}
