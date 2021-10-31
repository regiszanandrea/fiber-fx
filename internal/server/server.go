package server

import "github.com/gofiber/fiber/v2"

func Boot() func(app *fiber.App) error {
	return func(app *fiber.App) error {
		if err := app.Listen(":8080"); err != nil {
			return err
		}
		return nil
	}
}
