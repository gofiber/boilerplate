package routes

import "github.com/gofiber/fiber"

// NotFound sends a 404 "Not Found"
func NotFound(c *fiber.Ctx) {
	c.SendStatus(fiber.StatusNotFound)
}
