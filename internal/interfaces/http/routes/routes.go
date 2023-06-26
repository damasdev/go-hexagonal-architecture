package routes

import (
	userHandler "github.com/damasdev/fiber/internal/interfaces/http/v1/user"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {

	// Repository

	// Server

	// Handler
	userHandler := userHandler.New()

	// Routes
	v1 := app.Group("api/v1")

	// User Route
	v1.Get("/users", userHandler.FindAll)
}
