package server

import (
	"os"
	"time"

	"github.com/damasdev/fiber/internal/interfaces/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server interface {
	Run() error
	Stop() error

	RegisterMiddleware()
	RegisterHandler()
}

type fiberServer struct {
	*fiber.App
}

func New() Server {
	return &fiberServer{
		fiber.New(),
	}
}

func (f *fiberServer) Run() error {
	return f.Listen(":" + os.Getenv("APP_PORT"))
}

func (f *fiberServer) Stop() error {
	return f.ShutdownWithTimeout(5 * time.Second)
}

func (f *fiberServer) RegisterMiddleware() {
	f.Use(recover.New())
}

func (f *fiberServer) RegisterHandler() {
	routes.API(f)
}
