package server

import (
	"fmt"

	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run()
}

type server struct{}

func New() Server {
	return &server{}
}

func (srv *server) Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(fmt.Sprintf(":%d", 3000)); err != nil {
		logger.Logger.Fatal("failed to start server", err)
	}
}
