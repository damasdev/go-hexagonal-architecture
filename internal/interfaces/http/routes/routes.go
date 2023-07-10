package routes

import (
	userService "github.com/damasdev/fiber/internal/core/service/user"
	userRepository "github.com/damasdev/fiber/internal/infrastructure/repository/mongodb/user"
	userHandler "github.com/damasdev/fiber/internal/interfaces/http/handler/user"
	"github.com/damasdev/fiber/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {

	// Instance
	db := config.MongoDB

	// Repository
	userRepository := userRepository.New(db)

	// Service
	userService := userService.New(userRepository)

	// Handler
	userHandler := userHandler.New(userService)

	// Routes
	api := app.Group("api")

	// User Route
	api.Get("/users", userHandler.List)
}
