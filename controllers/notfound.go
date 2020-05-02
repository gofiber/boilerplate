package controllers

import "github.com/gofiber/fiber"

// NotFound returns 404 "Not Found"
func NotFound(c *fiber.Ctx) {
	c.SendStatus(fiber.StatusNotFound)
}
