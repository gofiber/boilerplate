package routes

import "github.com/gofiber/fiber"

// NotFound route
func NotFound(c *fiber.Ctx) {
	c.Status(404)
	c.Render("404", fiber.Map{
		"Title":   "404 Not Found - Fiber",
		"Message": "Sorry, but the page you were looking for could not be found.",
	})
}
