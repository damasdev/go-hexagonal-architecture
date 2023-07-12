package routes

import (
	userService "go-hexagonal-architecture/internal/core/service/user"
	userRepository "go-hexagonal-architecture/internal/infrastructure/repository/mongodb/user"
	userHandler "go-hexagonal-architecture/internal/interfaces/http/handler/user"
	"go-hexagonal-architecture/pkg/config"

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
