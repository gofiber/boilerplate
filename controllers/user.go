package controllers

import (
	"boilerplate/models"

	"github.com/gofiber/fiber"
)

// UserGet returns a user
func UserGet(c *fiber.Ctx) {
	user := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) {
	user := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// UserUpdate updates a user
func UserUpdate(c *fiber.Ctx) {
	user := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// UserDelete deletes a user
func UserDelete(c *fiber.Ctx) {
	user := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}
