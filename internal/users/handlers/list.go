package handlers

import "github.com/gofiber/fiber/v2"

func List(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": "john",
		"age":  20,
	})
}
