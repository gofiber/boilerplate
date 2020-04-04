package routes

import "github.com/gofiber/fiber"

// Error ...
func Error(c *fiber.Ctx, err error) { 
	c.Status(500)
	c.Render("500", fiber.Map{
		"Title":   "500 Bad Request - Fiber",
		"Message": err.Error(),
	})
}
