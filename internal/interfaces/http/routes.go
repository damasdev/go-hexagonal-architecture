package routes

import (
	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {
	// repository

	// service

	app.Get("/", func(c *fiber.Ctx) error {
		defer func() {
			logger.Logger.Debug("hello world", log.WithData(map[string]any{"hello": "world"}))
		}()
		return c.SendString("Hello, World!")
	})
}
