package server

import (
	"context"
	"fmt"
	"os"

	"github.com/damasdev/fiber/internal/interfaces/http/routes"
	"github.com/damasdev/fiber/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server interface {
	Run() error
	Stop(context.Context) error

	RegisterMiddleware()
	RegisterHook()
	RegisterHandler()
}

type fiberServer struct {
	app *fiber.App
}

func New() Server {
	return &fiberServer{
		app: fiber.New(config.FiberConfig()),
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
	// Register Router
	routes.API(f.app)
}

func (f *fiberServer) RegisterHook() {
	// Register Hook On Shutdown
	f.app.Hooks().OnShutdown(func() error {
		fmt.Println("shutdown..")
		return nil
	})
}
