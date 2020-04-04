package routes

import "github.com/gofiber/fiber"

// JSON route
func JSON(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"success": true,
		"payload": "Some data",
	})
}
