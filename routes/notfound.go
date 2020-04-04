package routes

import "github.com/gofiber/fiber"

// NotFound route
func NotFound(c *fiber.Ctx) {
	c.Status(404)
	c.Render("404", fiber.Map{
		"Title":   "404 Not Found - Fiber",
		"Message": "The page you requested was not found!",
	})
}
