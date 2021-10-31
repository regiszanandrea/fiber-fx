package system

import (
	"regiszanandrea/fiber-fx/internal/system/handlers"

	"github.com/gofiber/fiber/v2"
)

func Boot(app *fiber.App) {
	registerRoutes(app)
}

func registerRoutes(app *fiber.App) {
	group := app.Group("system")

	group.Get("/health", handlers.Health)
}
