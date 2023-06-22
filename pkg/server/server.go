package server

import (
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run()
}

type server struct {
	logger logger.Logger
}

func New(logger logger.Logger) Server {
	return &server{
		logger: logger,
	}
}

func (srv *server) Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3000"); err != nil {
		srv.logger.Fatal("failed to start server", err)
	}
}
