package routes

import (
	"regiszanandrea/fiber-fx/internal/system"
	"regiszanandrea/fiber-fx/internal/users"

	"github.com/gofiber/fiber/v2"
)

func Boot() interface{} {
	return func(app *fiber.App) {
		system.Boot(app)
		users.Boot(app)
	}
}
