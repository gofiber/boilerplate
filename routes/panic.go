package routes

import "github.com/gofiber/fiber"

// Panic test
func Panic(c *fiber.Ctx) {
	// This should trigger the recover function
	panic("Hi, I'm an error!")
}
