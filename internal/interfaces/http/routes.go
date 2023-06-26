package routes

import (
	userHandler "github.com/damasdev/fiber/internal/interfaces/http/user"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {

	// Repository

	// Server

	// Handler
	userHandler := userHandler.New()

	// Routes
	api := app.Group("api")

	// User Route
	api.Get("/users", userHandler.FindAll)
}
