package routes

import (
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {
	// repository

	// service

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
