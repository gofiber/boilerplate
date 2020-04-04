package routes

import "github.com/gofiber/fiber"

// Index route
func Index(c *fiber.Ctx) {
	c.Render("index", fiber.Map{
		"Title":   "Fiber",
		"Message": "Hello, World!",
	})
}
