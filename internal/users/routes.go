package users

import (
	"regiszanandrea/fiber-fx/internal/users/handlers"

	"github.com/gofiber/fiber/v2"
)

func Boot(app *fiber.App) {
	registerRoutes(app)
}

func registerRoutes(app *fiber.App) {
	group := app.Group("users")

	group.Get("/", handlers.List)
}
