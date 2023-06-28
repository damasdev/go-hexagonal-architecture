package server

import (
	"context"
	"os"
	"time"

	"github.com/damasdev/fiber/internal/interfaces/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server interface {
	Run() error
	Stop(context.Context) error

	RegisterMiddleware()
	RegisterHandler()
}

type fiberServer struct {
	app *fiber.App
}

func New() Server {
	return &fiberServer{
		app: fiber.New(fiber.Config{
			IdleTimeout: time.Second * 5,
		}),
	}
}

func (f *fiberServer) Run() error {
	return f.app.Listen(":" + os.Getenv("APP_PORT"))
}

func (f *fiberServer) Stop(ctx context.Context) error {
	return f.app.ShutdownWithContext(ctx)
}

func (f *fiberServer) RegisterMiddleware() {
	// Register Middleware
	f.app.Use(recover.New())
}

func (f *fiberServer) RegisterHandler() {
	routes.API(f.app)
}
