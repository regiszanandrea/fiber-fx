package main

import (
	"context"
	"log"
	"regiszanandrea/fiber-fx/internal/routes"
	"regiszanandrea/fiber-fx/internal/server"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle) *fiber.App {
	server := fiber.New()

	return server
}

func Register(server *fiber.App) {
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func main() {

	app := fx.New(
		fx.Provide(NewServer),
		fx.Invoke(Register),
		fx.Invoke(routes.Boot()),
		fx.Invoke(server.Boot()),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
